package store

import (
	"context"

	"github.com/dannylindquist/gopoints"
)

var _ gopoints.GameService = (*GameService)(nil)

type GameService struct {
}

func (gs *GameService) CreateGame(ctx context.Context, game gopoints.GameUpdate) {

}
