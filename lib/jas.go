package jas

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/google/go-github/github"
)

// Handler handler is an interface that objects can implement to be registered
// to get called on incoming webhook payloads.
//
// Each handler will have its HandlePayload method called in the order
// registered.
type Handler interface {
	HandlePayload(e []string, p *github.WebHookPayload)
}

// HandlerFunc is an adapter to allow the use of ordinary functions as Jas
// handlers. If f is a function with the appropriate signature, HandlerFunc(f)
// is a Handler object who's HandlePayload method calls f.
type HandlerFunc func(e []string, p *github.WebHookPayload)

func (h HandlerFunc) HandlePayload(e []string, p *github.WebHookPayload) {
	h(e, p)
}

// Jas is a stack of Handlers that can be invoked as a Handler. Jas
// handlers are evaluated in the order that they are added to the stack using
// the Register and RegisterHandler methods.
type Jas struct {
	handlers []Handler
}

// New returns a new Jas instance with no handlers preconfigured.
func NewJas() *Jas {
	return &Jas{}
}

// RegisteHandlerr adds a Handler onto the handlers stack. Handlers are
// invoked in the order they are added.
func (j *Jas) RegisterHandler(handler Handler) {
	j.handlers = append(j.handlers, handler)
}

// RegisterHandlerFunc adds a HandlerFunc onto the handlers stack. Handlers are
// invoked in the order they are added to a Negroni.
func (j *Jas) RegisterHandlerFunc(handler HandlerFunc) {
	j.RegisterHandler(HandlerFunc(handler))
}

func (j *Jas) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	e := r.Header["X-Github-Event"]
	if e == nil {
		log.Println("Body had no X-Github-Event")
	} else {

		payload, err := githubWebhookDecoder(&r.Body)
		if err != nil {
			log.Println(err)
		}

		for _, handler := range j.handlers {
			handler.HandlePayload(e, payload)
		}
	}
}

// githubWebhookDecoder decodes the payload into a github.WebHookPayload
func githubWebhookDecoder(body *io.ReadCloser) (*github.WebHookPayload, error) {
	var payload github.WebHookPayload
	decoder := json.NewDecoder(*body)
	err := decoder.Decode(&payload)

	if err != nil {
		return nil, err
	}

	return &payload, nil
}
