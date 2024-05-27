package api

import (
	"net/http"

	"github.com/frezzle/cards-api/pkg/decks"
)

type Server struct {
	serveMux     *http.ServeMux
	decksService *decks.Service
}

func New(decksService *decks.Service) *Server {
	s := &Server{
		serveMux:     http.NewServeMux(),
		decksService: decksService,
	}

	s.serveMux.HandleFunc("POST /decks", s.createDeck)
	s.serveMux.HandleFunc("GET /decks/{id}", s.openDeck)
	s.serveMux.HandleFunc("POST /decks/{id}/draw", s.drawCards)

	return s
}

func (s *Server) Serve(address string) error {
	return http.ListenAndServe(address, s.serveMux)
}
