// Captain 'Jas' Hook: A Github Hook handler.
//
// Jas can be included in a net/http server to listen on a path for github webhooks.
//
// For a full guide visit http://github.com/wm/jas
//
//      package main
//
//      import (
//      	"net/http"
//
//      	"github.com/wm/jas/lib"
//      )
//
//      func main() {
//      	j := jas.NewJas()
//      	j.RegisterHandlerFunc(jas.PushPayloadLogger)
//      	j.RegisterHandler(jas.NewFileChangeEmailer(jas.Options{
//      		Emails: &[]string{"will@example.com", "eliot@example.com"},
//      		Files:  &[]string{"db/structure.sql", "db/schema.rb"},
//      	}))
//
//      	mux := http.NewServeMux()
//      	mux.Handle("/", j)
//      	http.ListenAndServe(":1234", mux)
// }
package jas
