package gopoints

import "context"

type Game struct {
	ID      int      `json:"id"`
	Players []Player `json:"players"`
	Type    string   `json:"type"`
}

type GameService interface {
	// create a game
	CreateGame(ctx context.Context, update GameUpdate) *Game
	GetGame(ctx context.Context, id int) *Game
}

type GameUpdate struct {
	Name string `json:"name"`
	Type string `json:"type"`
}
