package store

import (
	"fmt"

	"github.com/frezzle/cards-api/pkg/models"
)

// Note: no business logic should exist here (i.e. the rules of how cards/decks change),
// only how to store them.

type MemoryStore struct {
	decks map[string]models.Deck
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		decks: make(map[string]models.Deck),
	}
}

func (ms *MemoryStore) CreateDeck(deck models.Deck) error {
	ms.decks[deck.ID] = deck
	return nil
}

func (ms *MemoryStore) GetDeck(id string) (models.Deck, error) {
	deck, ok := ms.decks[id]
	if !ok {
		return models.Deck{}, fmt.Errorf("deck does not exist with id '%s'", id)
	}
	return deck, nil
}

func (ms *MemoryStore) UpdateDeck(deck models.Deck) error {
	_, ok := ms.decks[deck.ID]
	if !ok {
		return fmt.Errorf("deck does not exist with id '%s'", deck.ID)
	}
	ms.decks[deck.ID] = deck
	return nil
}
