package models

import (
	"fmt"
	"strings"
)

type Deck struct {
	// ID uniquely identifies the deck from any other deck.
	ID string

	// Shuffled is whether the deck was shuffled or not when it was created.
	Shuffled bool

	// Cards that can still be drawn.
	// The next card to be drawn is at position 0, if there is one,
	// then the one at position 1, if there is one, and so on.
	Cards []Card
}

type Card struct {
	// Code is the unique identifier of the type of card,
	// e.g. "AS" is Ace of Spades, "10H" is 10 of Hearts, etc.
	Code string

	// Value is the face value of the card e.g. "ACE", "10", "QUEEN".
	Value string

	// Suit can be "CLUBS", "HEARTS", "SPADES" or "DIAMONDS".
	Suit string
}

var FullDeck52 = []Card{
	{Code: "AS", Value: "ACE", Suit: "SPADES"},
	{Code: "2S", Value: "2", Suit: "SPADES"},
	{Code: "3S", Value: "3", Suit: "SPADES"},
	{Code: "4S", Value: "4", Suit: "SPADES"},
	{Code: "5S", Value: "5", Suit: "SPADES"},
	{Code: "6S", Value: "6", Suit: "SPADES"},
	{Code: "7S", Value: "7", Suit: "SPADES"},
	{Code: "8S", Value: "8", Suit: "SPADES"},
	{Code: "9S", Value: "9", Suit: "SPADES"},
	{Code: "10S", Value: "10", Suit: "SPADES"},
	{Code: "JS", Value: "JACK", Suit: "SPADES"},
	{Code: "QS", Value: "QUEEN", Suit: "SPADES"},
	{Code: "KS", Value: "KING", Suit: "SPADES"},

	{Code: "AD", Value: "ACE", Suit: "DIAMONDS"},
	{Code: "2D", Value: "2", Suit: "DIAMONDS"},
	{Code: "3D", Value: "3", Suit: "DIAMONDS"},
	{Code: "4D", Value: "4", Suit: "DIAMONDS"},
	{Code: "5D", Value: "5", Suit: "DIAMONDS"},
	{Code: "6D", Value: "6", Suit: "DIAMONDS"},
	{Code: "7D", Value: "7", Suit: "DIAMONDS"},
	{Code: "8D", Value: "8", Suit: "DIAMONDS"},
	{Code: "9D", Value: "9", Suit: "DIAMONDS"},
	{Code: "10D", Value: "10", Suit: "DIAMONDS"},
	{Code: "JD", Value: "JACK", Suit: "DIAMONDS"},
	{Code: "QD", Value: "QUEEN", Suit: "DIAMONDS"},
	{Code: "KD", Value: "KING", Suit: "DIAMONDS"},

	{Code: "AC", Value: "ACE", Suit: "CLUBS"},
	{Code: "2C", Value: "2", Suit: "CLUBS"},
	{Code: "3C", Value: "3", Suit: "CLUBS"},
	{Code: "4C", Value: "4", Suit: "CLUBS"},
	{Code: "5C", Value: "5", Suit: "CLUBS"},
	{Code: "6C", Value: "6", Suit: "CLUBS"},
	{Code: "7C", Value: "7", Suit: "CLUBS"},
	{Code: "8C", Value: "8", Suit: "CLUBS"},
	{Code: "9C", Value: "9", Suit: "CLUBS"},
	{Code: "10C", Value: "10", Suit: "CLUBS"},
	{Code: "JC", Value: "JACK", Suit: "CLUBS"},
	{Code: "QC", Value: "QUEEN", Suit: "CLUBS"},
	{Code: "KC", Value: "KING", Suit: "CLUBS"},

	{Code: "AH", Value: "ACE", Suit: "HEARTS"},
	{Code: "2H", Value: "2", Suit: "HEARTS"},
	{Code: "3H", Value: "3", Suit: "HEARTS"},
	{Code: "4H", Value: "4", Suit: "HEARTS"},
	{Code: "5H", Value: "5", Suit: "HEARTS"},
	{Code: "6H", Value: "6", Suit: "HEARTS"},
	{Code: "7H", Value: "7", Suit: "HEARTS"},
	{Code: "8H", Value: "8", Suit: "HEARTS"},
	{Code: "9H", Value: "9", Suit: "HEARTS"},
	{Code: "10H", Value: "10", Suit: "HEARTS"},
	{Code: "JH", Value: "JACK", Suit: "HEARTS"},
	{Code: "QH", Value: "QUEEN", Suit: "HEARTS"},
	{Code: "KH", Value: "KING", Suit: "HEARTS"},
}

// Converts card-codes string to card objects,
// e.g. "AS,KD,AC" becomes those 3 cards.
// Errors if a card code is invalid.
func CardsFromCodes(cardCodes string) ([]Card, error) {
	// TODO we could have this mapping calculated once in the package instead of every call; use init() func?
	cardFromCode := make(map[string]Card, len(FullDeck52))
	for _, card := range FullDeck52 {
		cardFromCode[card.Code] = card
	}

	codes := strings.Split(cardCodes, ",")
	cards := make([]Card, len(codes))
	for i, code := range codes {
		card, ok := cardFromCode[code]
		if !ok {
			return nil, fmt.Errorf("card invalid with code '%s'", code)
		}
		cards[i] = card
	}

	return cards, nil
}
