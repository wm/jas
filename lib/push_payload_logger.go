package jas

import (
	"bytes"
	"log"
	"text/template"

	"github.com/google/go-github/github"
)

// PushPayloadLogger prints some basic payload information for push events.
func PushPayloadLogger(e []string, p *github.WebHookPayload) {
	for _, event := range e {
		if event == "push" {
			prettyPrintPayload(p)
		}
	}
}

func prettyPrintPayload(payload *github.WebHookPayload) {
	var prettyBuffer bytes.Buffer
	meta, err := template.New("Meta").Parse(`
  Pusher:       {{.Pusher}}
  Repo:         {{.Repo.Name}}
  Compare:      {{.Compare}}
  Commits:
    Before:     {{.Before}}
    After:      {{.After}}
  `)

	if err != nil {
		log.Println("Error: failed to parse payload template")
		return
	}
	meta.Execute(&prettyBuffer, payload)
	log.Println("Push[Payload]:", prettyBuffer.String())
}
