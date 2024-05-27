package decks

import (
	"fmt"

	"github.com/frezzle/cards-api/pkg/models"
	"github.com/frezzle/cards-api/pkg/store"
	"github.com/google/uuid"
)

type Service struct {
	store store.Store
}

func New(store store.Store) *Service {
	return &Service{
		store: store,
	}
}

func (s *Service) CreateDeck(shuffled bool, cardCodes string) (models.Deck, error) {
	cards := models.FullDeck52
	var err error
	if cardCodes != "" {
		cards, err = models.CardsFromCodes(cardCodes)
		if err != nil {
			return models.Deck{}, fmt.Errorf("failed parsing cards: %w", err)
		}
	}

	if shuffled {
		cards = shuffle(cards)
	}

	deck := models.Deck{
		ID:       uuid.NewString(),
		Shuffled: shuffled,
		Cards:    cards,
	}
	err = s.store.CreateDeck(deck)
	if err != nil {
		return models.Deck{}, fmt.Errorf("failed creating deck: %w", err)
	}

	return deck, nil
}

func (s *Service) OpenDeck(id string) (models.Deck, error) {
	deck, err := s.store.GetDeck(id)
	if err != nil {
		return models.Deck{}, fmt.Errorf("failed getting deck: %w", err)
	}
	return deck, nil
}

// DrawCards draws the number of cards from the deck.
// It errors if there's not enough cards remaining or if you try to draw a non-positive number of cards.
func (s *Service) DrawCards(deckId string, count int) ([]models.Card, error) {
	if count < 1 {
		return nil, fmt.Errorf("not allowed to draw %d cards", count)
	}

	deck, err := s.store.GetDeck(deckId)
	if err != nil {
		return nil, fmt.Errorf("failed getting deck: %w", err)
	}

	if count > len(deck.Cards) {
		return nil, fmt.Errorf("deck only has %d cards, cannot draw %d", len(deck.Cards), count)
	}

	cardsDrawn := deck.Cards[:count]
	deck.Cards = deck.Cards[count:]
	err = s.store.UpdateDeck(deck)
	if err != nil {
		return nil, fmt.Errorf("failed updating deck: %w", err)
	}

	return cardsDrawn, nil
}
