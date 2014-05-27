# Jas

[![GoDoc](https://godoc.org/github.com/wm/jas/lib?status.png)](https://godoc.org/github.com/wm/jas/lib)

Captain 'Jas' Hook: A Github Hook handler.

Jas can be included in a net/http server to listen on a path for github webhooks.

## Basic Setup

Install the Jas package (**go 1.1** and greater is required):
```
go get github.com/wm/jas
```

```go
package main

import (
	"net/http"

	"github.com/wm/jas/lib"
)

func main() {
	j := jas.NewJas()
	j.RegisterHandlerFunc(jas.PushPayloadLogger)
	j.RegisterHandler(jas.NewFileChangeEmailer(jas.FileChangeEmailerOptions{
		Emails: &[]string{"will@example.com", "eliot@example.com"},
		Files:  &[]string{"db/structure.sql", "db/schema.rb"},
	}))

	mux := http.NewServeMux()
	mux.Handle("/", j)
	http.ListenAndServe(":1234", mux)
}

```

Then run your server:
```
go run main.go
```

You will now have a web server running on port 1234 that can handle incomming
Github webhooks. Next you need to add a webhook to a github repo. See the
[Github Webhook Guide](https://developer.github.com/webhooks/) for help.

See [main.go](main.go) for an example server.

## Handlers

TODO
