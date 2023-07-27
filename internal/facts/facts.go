package facts

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"
	"strings"
)

type FactsGetter interface {
	GetCatFacts() string
}

type CatFactResponse struct {
	Data []string
}

var REGEX_CAT_WORD_INSENSITIVE = regexp.MustCompile(`(?i)\bcats?\b`)

func GetGopherFacts(factsGetter FactsGetter) string {

	catFact := factsGetter.GetCatFacts()

	gopherFact := REGEX_CAT_WORD_INSENSITIVE.ReplaceAllStringFunc(catFact, func(match string) string {
		if strings.HasPrefix(match, "cat") {
			return handlePlural("gopher", isPlural(match))
		} else {
			return handlePlural("Gopher", isPlural(match))
		}
	})

	return gopherFact
}

type HttpFactsGetter struct{}

func (h HttpFactsGetter) GetCatFacts() string {
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

func isPlural(word string) bool {
	return strings.HasSuffix(word, "s")
}

func handlePlural(word string, isPlural bool) string {
	if isPlural {
		return word + "s"
	} else {
		return word
	}
}
