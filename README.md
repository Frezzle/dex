# dex, a general-purpose cards API

An API for creating decks of cards, getting their current state, and drawing cards from them.

The cards are currently specific to the standard 52-card French-suited deck.

## Running the app

Install Go, then `make run` from the root project directory.

## Endpoints

Create an unshuffled standard 52-card deck:
```console
curl -X POST http://localhost:3000/decks
```
```json
{
  "deck_id": "aa7f842d-50cf-4269-80d7-fb20553593c3",
  "shuffled": false,
  "remaining": 52
}
```

Create a shuffled deck:
```console
curl -X POST http://localhost:3000/decks?shuffled
```
```json
{
  "deck_id": "bb7f842d-50cf-4269-80d7-fb20553593c3",
  "shuffled": true,
  "remaining": 52
}
```

Create a deck with your choice of cards (see [here](./pkg/models/models.go) for valid card codes):
```console
curl -X POST http://localhost:3000/decks?cards=AH,2S,KD
```
```json
{
  "deck_id": "cc7f842d-50cf-4269-80d7-fb20553593c3",
  "shuffled": false,
  "remaining": 52
}
```

Retrieve a deck (see [here](./pkg/models/models.go) for expected card values):
```console
curl -X GET http://localhost:3000/decks/{id}
```
```json
{
  "deck_id": "75b3479e-403f-45f1-bba4-ebbe090818b4",
  "shuffled": true,
  "remaining": 4,`
  "cards": [
    {
      "code": "AS",
      "value": "ACE",
      "suit": "SPADES"
    },
    ...
  ]
}
```

Draw one card from a deck (cards are drawn from the start of the `cards` array shown above):
```console
curl -X POST http://localhost:3000/decks/{id}/draw
```
```json
{
  "cards": [
    {
      "code": "AS",
      "value": "ACE",
      "suit": "SPADES"
    }
  ]
}
```

Draw multiple cards from a deck e.g. 2:
```console
curl -X POST http://localhost:3000/decks/{id}/draw?count=2
```
```json
{
  "cards": [
    {
      "code": "AS",
      "value": "ACE",
      "suit": "SPADES"
    },
    {
      "code": "QH",
      "value": "QUEEN",
      "suit": "HEARTS"
    },
  ]
}
```

## Architecture

### Store

The `store` package is responsible for persisting and retrieving decks and their cards. An interface defines how to interact with any kind of store. A memory store implementation is used as a proof of concept, though in production we'd likely use an implementation that talks to a database.

### Decks service

The `decks` package defines the service responsible for the business logic of all the current endpoints. If the application contained other domains then we could have more services to separate the logic.

A store is injected into the decks service constructor, allowing easy unit testing of the service (e.g. with the memory store) and production-scenario data persistence (e.g. with a database store).

### API

The `api` package is responsible for providing the interface for web clients to query the application and receive results. It basically only processes HTTP requests and sends back responses, delegating actual business logic to other service packages (i.e. only the decks service right now).

## Future improvements

More time allowing, I could:
- Generalise the cards more, so it could have any type of card e.g. Uno cards.
- Add tests for decks package; it's already been designed so that it could be easily unit tested (by injecting the memory store implementation).
- The API may currently expose internal errors. A production service would only expose useful info but nothing about internal implementation.
- Lately I've been trying Go's standard `http` library to create web servers, so I did the same here; in production I'd probably go with something with more conveniences, like [chi](https://github.com/go-chi/chi) or similar.
