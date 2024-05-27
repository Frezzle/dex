package store

import (
	"testing"

	"github.com/frezzle/cards-api/pkg/models"
	"github.com/stretchr/testify/assert"
)

func TestMemoryStore(t *testing.T) {
	ms := NewMemoryStore()

	t.Run("error when getting non-existent deck", func(t *testing.T) {
		_, err := ms.GetDeck("d1")
		assert.NotNil(t, err)
	})

	t.Run("error updating non-existent deck", func(t *testing.T) {
		err := ms.UpdateDeck(models.Deck{
			ID:       "d1",
			Shuffled: true,
			Cards:    []models.Card{},
		})
		assert.NotNil(t, err)
	})

	t.Run("create deck, retrieve it, update it, retrieve it again", func(t *testing.T) {
		err := ms.CreateDeck(models.Deck{
			ID:       "d1",
			Shuffled: true,
			Cards: []models.Card{
				{
					Code:  "AS",
					Value: "ACE",
					Suit:  "SPADES",
				},
			},
		})
		assert.Nil(t, err)

		deck, err := ms.GetDeck("d1")
		assert.Nil(t, err)
		assert.Equal(t, models.Deck{
			ID:       "d1",
			Shuffled: true,
			Cards: []models.Card{
				{
					Code:  "AS",
					Value: "ACE",
					Suit:  "SPADES",
				},
			},
		}, deck)

		err = ms.UpdateDeck(models.Deck{
			ID:       "d1",
			Shuffled: false,
			Cards:    []models.Card{},
		})
		assert.Nil(t, err)

		deck, err = ms.GetDeck("d1")
		assert.Nil(t, err)
		assert.Equal(t, models.Deck{
			ID:       "d1",
			Shuffled: false,
			Cards:    []models.Card{},
		}, deck)
	})
}
