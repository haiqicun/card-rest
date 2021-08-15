package utils

import (
	"math/rand"
	"strings"
	"time"

	"github.com/haiqicun/card-rest/pkg/models"
)

//Constants of the deck
const (
	SPADES   = "S"
	DIAMONDS = "D"
	CLUBS    = "C"
	HEARTS   = "H"

	ACE   = "A"
	TWO   = "2"
	THREE = "3"
	FOUR  = "4"
	FIVE  = "5"
	SIX   = "6"
	SEVEN = "7"
	EIGHT = "8"
	NINE  = "9"
	TEN   = "10"
	JACK  = "J"
	QUEEN = "Q"
	KING  = "K"
)

//Suit and Cards default values
var (
	SuitDefault  = []string{SPADES, DIAMONDS, CLUBS, HEARTS}
	CardsDefault = []string{ACE, TWO, THREE, FOUR, FIVE, SIX, SEVEN, EIGHT, NINE, TEN, JACK, QUEEN, KING}
)

//SuitMap from codes to values
var SuitMap = map[string]string{
	SPADES:   "SPADES",
	CLUBS:    "CLUBS",
	DIAMONDS: "DIAMONDS",
	HEARTS:   "HEARTS",
}

//CardsMap from codes to values
var CardsMap = map[string]string{
	ACE:   "ACE",
	TWO:   TWO,
	THREE: THREE,
	FOUR:  FOUR,
	FIVE:  FIVE,
	SIX:   SIX,
	SEVEN: SEVEN,
	EIGHT: EIGHT,
	NINE:  NINE,
	TEN:   TEN,
	JACK:  "JACK",
	QUEEN: "QUEEN",
	KING:  "KING",
}

//GetDefaultCodes is to generate the default 52 codes.
func GetDefaultCodes() (codes []string) {
	for _, s := range SuitDefault {
		for _, c := range CardsDefault {
			code := c + s
			codes = append(codes, code)
		}
	}

	return codes
}

//ShuffleDeck is to shuffle the deck randomly.
func ShuffleDeck(codes []string) {
	rand.Seed(time.Now().UnixNano())
	for i := len(codes) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		codes[i], codes[j] = codes[j], codes[i]
	}
}

//GenerateCardsWithCodes is to generate a list of Card objects based on the codes.
func GenerateCardsWithCodes(codes []string) (models.Cards, error) {
	if len(codes) == 0 {
		return models.Cards{}, nil
	}

	cards := make(models.Cards, 0, len(codes))

	for _, code := range codes {
		var (
			c string
			s string
		)
		if strings.Contains(code, TEN) {
			c = code[0:2]
			s = code[2:3]
		} else {
			c = code[0:1]
			s = code[1:2]
		}
		card := &models.Card{
			Value: CardsMap[c],
			Suit:  SuitMap[s],
			Code:  code,
		}
		cards = append(cards, card)
	}

	return cards, nil
}

//CheckCardCodes is to check the codes are good codes or not.
//Good codes are the code contained in 52 default codes.
func CheckCardCodes(codes []string) bool {
	validCodes := GetDefaultCodes()
	tempCodes := []string{}
	for _, code := range codes {
		if !find(validCodes, code) || find(tempCodes, code) {
			return false
		}
		tempCodes = append(tempCodes, code)
	}
	return true
}

func find(items []string, s string) bool {
	for _, i := range items {
		if i == s {
			return true
		}
	}

	return false
}
