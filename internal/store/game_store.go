package store

import (
	"context"
	"sync"

	"github.com/dannylindquist/gopoints"
)

var _ gopoints.GameService = (*GameService)(nil)

type GameService struct {
	db    *Store
	id    int
	mutex sync.Mutex
}

func NewGameService(db *Store) *GameService {
	return &GameService{
		db: db,
		id: 1,
	}
}

func (gs *GameService) CreateGame(ctx context.Context, game gopoints.GameUpdate) *gopoints.Game {
	gs.mutex.Lock()
	defer gs.mutex.Unlock()
	newGame := &gopoints.Game{
		ID:      gs.id,
		Type:    "ASC",
		Players: make([]gopoints.Player, 0),
	}
	gs.db.SetItems("games", newGame.ID, newGame)
	gs.id++
	return newGame
}

func (gs *GameService) GetGame(ctx context.Context, id int) *gopoints.Game {
	gs.mutex.Lock()
	defer gs.mutex.Unlock()
	game, ok := gs.db.GetItems("games")[id]
	if !ok {
		return nil
	}
	casted, ok := game.(*gopoints.Game)
	if !ok {
		return nil
	}
	return casted
}
