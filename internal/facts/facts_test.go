package facts_test

import (
	"sus-gopher-facts/internal/facts"
	"testing"
)

type FakeFactsGetter struct {
	GetCatFactsResponse string
}

func (f FakeFactsGetter) GetCatFacts() string {
	return f.GetCatFactsResponse
}

func TestGetGopherFacts(t *testing.T) {

	t.Run("should replace all the word cats maintaining casing", func(t *testing.T) {

		input := `Ends with a cat.
		Cat starts a paragraph. Cat starts the phrase. Ends with cats.
		Cats starts a paragraph. Catnip shouldn't change.`

		want := `Ends with a gopher.
		Gopher starts a paragraph. Gopher starts the phrase. Ends with gophers.
		Gophers starts a paragraph. Catnip shouldn't change.`

		got := facts.GetGopherFacts(FakeFactsGetter{input})

		if got != want {
			t.Errorf("Got: %s\nWant: %s", got, want)
		}
	})
}
