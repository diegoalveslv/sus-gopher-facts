package main

import (
	"fmt"
	"log"
	"net/http"
	"sus-gopher-facts/internal/facts"
)

var (
	serverPort      = "8080"
	httpFactsGetter = facts.HttpFactsGetter{}
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	catFactEndpoint := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodGet {
			gopherFacts := facts.GetGopherFacts(httpFactsGetter)

			fmt.Fprint(w, gopherFacts)
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	log.Printf("Starting server at %s", serverPort)
	log.Panic(http.ListenAndServe(":"+serverPort, catFactEndpoint))
}
