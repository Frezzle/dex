package api

import (
	"fmt"
	"net/http"
	"strconv"
)

func (s *Server) createDeck(w http.ResponseWriter, r *http.Request) {
	shuffled := r.URL.Query().Has("shuffled")
	cards := r.URL.Query().Get("cards")

	deck, err := s.decksService.CreateDeck(shuffled, cards)
	if err != nil {
		http.Error(w, fmt.Errorf("failed creating deck: %w", err).Error(), http.StatusInternalServerError)
		return
	}

	JSON(w, struct {
		DeckID    string `json:"deck_id"`
		Shuffled  bool   `json:"shuffled"`
		Remaining int    `json:"remaining"`
	}{
		DeckID:    deck.ID,
		Shuffled:  deck.Shuffled,
		Remaining: len(deck.Cards),
	})
}

func (s *Server) openDeck(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	deck, err := s.decksService.OpenDeck(id)
	if err != nil {
		http.Error(w, fmt.Errorf("failed getting deck: %w", err).Error(), http.StatusBadRequest)
		return
	}

	type card struct {
		Code  string `json:"code"`
		Value string `json:"value"`
		Suit  string `json:"suit"`
	}
	response := struct {
		DeckID    string `json:"deck_id"`
		Shuffled  bool   `json:"shuffled"`
		Remaining int    `json:"remaining"`
		Cards     []card `json:"cards"`
	}{
		DeckID:    deck.ID,
		Shuffled:  deck.Shuffled,
		Remaining: len(deck.Cards),
		Cards:     make([]card, len(deck.Cards)),
	}
	for i, c := range deck.Cards {
		response.Cards[i] = card{
			Code:  c.Code,
			Value: c.Value,
			Suit:  c.Suit,
		}
	}
	JSON(w, response)
}

func (s *Server) drawCards(w http.ResponseWriter, r *http.Request) {
	deckId := r.PathValue("id")
	count := 1
	var err error
	if r.URL.Query().Has("count") {
		count, err = strconv.Atoi(r.URL.Query().Get("count"))
		if err != nil {
			http.Error(w, "count param is not a valid int", http.StatusBadRequest)
			return
		}
	}

	cards, err := s.decksService.DrawCards(deckId, count)
	if err != nil {
		http.Error(w, fmt.Errorf("failed drawing cards: %w", err).Error(), http.StatusBadRequest)
		return
	}

	type card struct {
		Code  string `json:"code"`
		Value string `json:"value"`
		Suit  string `json:"suit"`
	}
	response := struct {
		Cards []card `json:"cards"`
	}{
		Cards: make([]card, len(cards)),
	}
	for i, c := range cards {
		response.Cards[i] = card{
			Code:  c.Code,
			Value: c.Value,
			Suit:  c.Suit,
		}
	}
	JSON(w, response)
}
