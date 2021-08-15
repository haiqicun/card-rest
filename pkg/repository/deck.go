package repository

import (
	"github.com/haiqicun/card-rest/pkg/models"
	"github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
)

//CreateDeck is to create a new deck entry in the table
func (db Database) CreateDeck(deck *models.Deck) error {
	deck.DeckId = uuid.NewV4().String()
	deck.Remaining = int(len(deck.CardCodes))
	statement := `INSERT INTO decks (deck_id, shuffled, remaining, card_codes) VALUES ($1, $2, $3, $4)`
	_, err := db.Conn.Exec(statement, deck.DeckId, deck.Shuffled, deck.Remaining, pq.Array(deck.CardCodes))
	if err != nil {
		return err
	}

	return nil
}

//GetDeckById is to get a deck entry from the table based on "deck_id"
func (db Database) GetDeckById(deckId string) (models.Deck, error) {
	deck := models.Deck{}
	query := `SELECT * FROM decks WHERE deck_id = $1`

	row := db.Conn.QueryRow(query, deckId)
	err := row.Scan(&deck.DeckId, &deck.Shuffled, &deck.Remaining, pq.Array(&deck.CardCodes))

	if err != nil {
		return deck, err
	}
	return deck, nil
}

//UpdateDeck is to update a deck based on the "deck_id" and the updating contents
func (db Database) UpdateDeck(deck *models.Deck) error {
	deck.Remaining = int(len(deck.CardCodes))
	statement := `UPDATE decks SET remaining=$1, card_codes=$2 WHERE deck_id=$3`
	_, err := db.Conn.Exec(statement, deck.Remaining, pq.Array(deck.CardCodes), deck.DeckId)

	if err != nil {
		return err
	}
	return nil
}
