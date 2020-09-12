package usecase

import (
	"context"

	"github.com/ak-17/ipl_stats/model"
)

type Usecase interface {
	GetFantasyPointsByTeam(ctx context.Context, teams []string) []model.Team
	GetPlayersByTeam(ctx context.Context, team string) []model.Player
}
