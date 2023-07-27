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

		input := `Cat cat cat. cat's Cats cats cats. Catnip tomcat`

		want := `Gopher gopher gopher. gopher's Gophers gophers gophers. Catnip tomcat`

		got := facts.GetGopherFacts(FakeFactsGetter{input})

		if got != want {
			t.Errorf("Got: %s\nWant: %s", got, want)
		}
	})
}
