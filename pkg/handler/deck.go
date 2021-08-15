package handler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/haiqicun/card-rest/pkg/models"
	"github.com/haiqicun/card-rest/pkg/repository"
	deck_util "github.com/haiqicun/card-rest/pkg/utils"
)

// The routes map the endpoints
func deck(router chi.Router) {
	router.Post("/", createDeck)
	router.Route("/{deckID}", func(router chi.Router) {
		router.Get("/", openDeck)
		router.Put("/draw/{counts}", drawDeck)
	})
}

// The function to create a deck entry in database
// and return "deck_id", "shuffled" and "remaining" as response.
func createDeck(w http.ResponseWriter, r *http.Request) {
	deck := &models.Deck{}
	deck.Shuffled = false
	if err := render.Bind(r, deck); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}

	if len(deck.Codes) == 0 {
		deck.CardCodes = deck_util.GetDefaultCodes()
	} else {
		deck.CardCodes = strings.Split(deck.Codes, ",")
	}
	if !deck_util.CheckCardCodes(deck.CardCodes) {
		render.Render(w, r, InvalidParameterErrorRenderer("Invalid card codes!"))
		return
	}
	if deck.Shuffled {
		deck_util.ShuffleDeck(deck.CardCodes)
	}

	if err := dbInstance.CreateDeck(deck); err != nil {
		render.Render(w, r, ErrorRenderer(err))
		return
	}
	result := &models.Deck{}
	result.DeckId = deck.DeckId
	result.Shuffled = deck.Shuffled
	result.Remaining = deck.Remaining
	if err := render.Render(w, r, result); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}

// To get a deck entry from database based on the "deckID"
func openDeck(w http.ResponseWriter, r *http.Request) {
	deckID := chi.URLParam(r, "deckID")
	deck, err := dbInstance.GetDeckById(deckID)
	if err != nil {
		if err == repository.ErrNoMatch {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ErrorRenderer(err))
		}
		return
	}
	cards, err := deck_util.GenerateCardsWithCodes(deck.CardCodes)
	if err != nil {
		render.Render(w, r, ErrorRenderer(err))
	}
	deck.Cards = cards
	if err := render.Render(w, r, &deck); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}

//To draw "counts" cards of a deck based on the "deckID" and update the database
func drawDeck(w http.ResponseWriter, r *http.Request) {
	deckID := chi.URLParam(r, "deckID")
	counts := chi.URLParam(r, "counts")
	if deckID == "" || counts == "" {
		render.Render(w, r, InvalidParameterErrorRenderer("deckId and counts are required"))
		return
	}
	countNum, err := strconv.ParseInt(counts, 10, 64)
	if err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}

	deck, err := dbInstance.GetDeckById(deckID)
	if err != nil {
		if err == repository.ErrNoMatch {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ErrorRenderer(err))
		}
		return
	}

	if deck.Remaining == 0 || countNum > int64(deck.Remaining) {
		render.Render(w, r, InvalidParameterErrorRenderer("counts is larger than remaining"))
		return
	}

	drawnCodes := deck.CardCodes[:countNum]
	left := deck.CardCodes[countNum:]
	deck.CardCodes = left

	err1 := dbInstance.UpdateDeck(&deck)

	if err1 != nil {
		if err1 == repository.ErrNoMatch {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ServerErrorRenderer(err1))
		}
		return
	}

	cards, err := deck_util.GenerateCardsWithCodes(drawnCodes)
	if err != nil {
		render.Render(w, r, ErrorRenderer(err))
		return
	}
	result := &models.Deck{}
	result.DeckId = deck.DeckId
	result.Remaining = deck.Remaining
	result.DrawnCards = cards
	if err := render.Render(w, r, result); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}
