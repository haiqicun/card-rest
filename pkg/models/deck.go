package models

import (
	"net/http"
)

type (
	// Deck is the object contains deck attributes.
	Deck struct {
		DeckId     string   `json:"deck_id" db:"deck_id"`
		Shuffled   bool     `json:"shuffled" db:"shuffled"`
		Remaining  int      `json:"remaining" db:"remaining"`
		Codes      string   `json:"codes,omitempty"`
		Cards      Cards    `json:"cards,omitempty"`
		CardCodes  []string `json:"-" db:"card_codes"`
		DrawnCards Cards    `json:"drawncards,omitempty"`
	}

	// Card isthe object contains card attributes.
	Card struct {
		Value string `json:"value"`
		Suit  string `json:"suit"`
		Code  string `json:"code"`
	}

	// Cards is a list of Card objects.
	Cards []*Card
)

// Bind is the function to bind the request parameters to the Deck object.
func (*Deck) Bind(r *http.Request) error {
	return nil
}

// Render is the function to render the Deck object as response.
func (*Deck) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
