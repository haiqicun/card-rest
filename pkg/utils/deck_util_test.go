package utils

import (
	"reflect"
	"testing"

	"github.com/haiqicun/card-rest/pkg/models"
)

func TestGetDefaultCodes(t *testing.T) {
	got := GetDefaultCodes()
	want := []string{"AS", "2S", "3S", "4S", "5S", "6S", "7S", "8S", "9S", "10S", "JS", "QS", "KS",
		"AD", "2D", "3D", "4D", "5D", "6D", "7D", "8D", "9D", "10D", "JD", "QD", "KD",
		"AC", "2C", "3C", "4C", "5C", "6C", "7C", "8C", "9C", "10C", "JC", "QC", "KC",
		"AH", "2H", "3H", "4H", "5H", "6H", "7H", "8H", "9H", "10H", "JH", "QH", "KH"}

	result := reflect.DeepEqual(got, want)
	if result == false {
		t.Errorf("TestGetDefaultCodes failed")
	}
}

func TestCheckCardCodes(t *testing.T) {
	t.Run("TestCheckCardCodesWithValidCodes", func(t *testing.T) {
		cardCodes := []string{"AS", "2D", "10H", "KC"}
		got := CheckCardCodes(cardCodes)

		if got == false {
			t.Errorf("TestGetDefaultCodes failed")
		}
	})
	t.Run("TestCheckCardCodesWithInvalidCodes", func(t *testing.T) {
		cardCodes := []string{"AS", "14D", "10H", "KC"}
		got := CheckCardCodes(cardCodes)

		if got == true {
			t.Errorf("TestGetDefaultCodes failed")
		}
	})
}

func TestGenerateCardsWithCodes(t *testing.T) {
	t.Run("TestGenerateCardsWithCodes", func(t *testing.T) {
		cardCodes := []string{"AS", "2D", "10H", "KC"}
		got, err := GenerateCardsWithCodes(cardCodes)
		want := models.Cards{
			{Value: "ACE", Suit: "SPADES", Code: "AS"},
			{Value: "2", Suit: "DIAMONDS", Code: "2D"},
			{Value: "10", Suit: "HEARTS", Code: "10H"},
			{Value: "KING", Suit: "CLUBS", Code: "KC"},
		}
		if err != nil {
			t.Errorf("TestGenerateCardsWithCodes with err")
		}
		result := reflect.DeepEqual(got, want)

		if result == false {
			t.Errorf("TestGenerateCardsWithCodes failed")
		}
	})

	t.Run("TestGenerateCardsWithCodesWithEmpty", func(t *testing.T) {
		cardCodes := []string{}
		got, err := GenerateCardsWithCodes(cardCodes)
		want := models.Cards{}
		if err != nil {
			t.Errorf("Calling GenerateCardsWithCodes with err")
		}
		result := reflect.DeepEqual(got, want)

		if result == false {
			t.Errorf("TestGenerateCardsWithCodesWithEmpty failed")
		}
	})
}
