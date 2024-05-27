package main

import (
	"fmt"
	"log"

	"github.com/frezzle/cards-api/pkg/api"
	"github.com/frezzle/cards-api/pkg/decks"
	"github.com/frezzle/cards-api/pkg/store"
)

func main() {
	store := store.NewMemoryStore()
	decksService := decks.New(store)
	server := api.New(decksService)
	addr := "localhost:3000"
	log.Printf("Listening on %s...\n", addr)
	if err := server.Serve(addr); err != nil {
		log.Fatalln(fmt.Errorf("server error: %w", err))
	}
}
