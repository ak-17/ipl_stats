package repository

import (
	"context"

	"github.com/ak-17/ipl_stats/model"
)

type Repository interface {
	GetAllPlayers(ctx context.Context) []model.Player
	GetPlayersByTeam(ctx context.Context, team string) []model.Player
}
