package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type CatFactResponse struct {
	Data []string 
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	catFactEndpoint := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodGet {
			catFact := getCatFact()
			catFact = strings.Replace(catFact, "cat", "gopher", -1)
			catFact = strings.Replace(catFact, "cats", "gophers", -1)
			catFact = strings.Replace(catFact, "Cat", "Gopher", -1)
			catFact = strings.Replace(catFact, "Cats", "Gophers", -1)

			fmt.Fprint(w, catFact)
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	log.Print("Starting server")
	log.Panic(http.ListenAndServe(":8080", catFactEndpoint))
}

func getCatFact() string {
	url := "https://meowfacts.herokuapp.com"

	response, err := http.Get(url)
	if err != nil {
		log.Println("Error occurred calling API:", err)
		return ""
	}
	defer response.Body.Close()

	var catFactResponse CatFactResponse
	err = json.NewDecoder(response.Body).Decode(&catFactResponse)
	if err != nil {
		log.Println("Error reading response:", err)
		return ""
	}

	return catFactResponse.Data[0]
}
