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

func (gs *GameService) CreateGame(ctx context.Context, game gopoints.GameUpdate) {
	gs.mutex.Lock()
	defer gs.mutex.Unlock()
	gs.db.SetItems("games", gs.id, game)
}
