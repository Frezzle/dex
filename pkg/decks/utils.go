package decks

import (
	"math/rand"

	"github.com/frezzle/cards-api/pkg/models"
)

// Shuffle copies and shuffles a collection of cards.
// The cards source is not altered.
func shuffle(cards []models.Card) []models.Card {
	// Source: https://stackoverflow.com/a/12264918
	dest := make([]models.Card, len(cards))
	perm := rand.Perm(len(cards))
	for i, v := range perm {
		dest[v] = cards[i]
	}
	return dest
}
