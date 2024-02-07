package gopoints

import "context"

type Game struct {
	ID      string   `json:"id"`
	Players []Player `json:"players"`
	Type    string   `json:"type"`
}

type GameService interface {
	// create a game
	CreateGame(ctx context.Context, update GameUpdate)
}

type GameUpdate struct {
	Name *string `json:"name"`
	Type *string `json:"type"`
}
