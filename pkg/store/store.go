package store

import "github.com/frezzle/cards-api/pkg/models"

type Store interface {
	CreateDeck(deck models.Deck) error
	GetDeck(id string) (models.Deck, error)
	UpdateDeck(deck models.Deck) error
}
