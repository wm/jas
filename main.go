package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/wm/jas/lib"
)

func main() {
	port := flag.String("port", "2468", "port to listen on")
	logFile := flag.String("log", "./log/server.log", "log file")
	flag.Parse()

	setLog(logFile)

	j := jas.NewJas()
	j.RegisterHandlerFunc(jas.PushPayloadLogger)
	j.RegisterHandler(jas.NewFileChangeEmailer(jas.FileChangeEmailerOptions{
		Emails: &[]string{"will@example.com", "eliot@example.com"},
		Files:  &[]string{"db/structure.sql", "db/schema.rb"},
	}))

	mux := http.NewServeMux()
	mux.Handle("/", j)
	http.ListenAndServe(":"+*port, mux)
}

func setLog(logFile *string) {
	log_handler, err := os.OpenFile(*logFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		panic("cannot write to log")
	}
	log.SetOutput(log_handler)
	log.SetFlags(5)
}
