package apiVersion1

import (
	"context"

	"github.com/ak-17/ipl_stats/model"
)

func (api *apiVersion1) GetPlayersByTeam(ctx context.Context, team string) []model.Player {
	players := api.repo.GetPlayersByTeam(ctx, team)

	return players
}
