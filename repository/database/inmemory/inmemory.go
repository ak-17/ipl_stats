package inmemory

import (
	"context"

	"github.com/ak-17/ipl_stats/model"
)

func (im *inMemory) GetAllPlayers(ctx context.Context) []model.Player {
	return im.Players
}

func (im *inMemory) GetPlayersByTeam(ctx context.Context, team string) []model.Player {
	var res []model.Player
	for _, player := range im.Players {
		if player.Team == team {
			res = append(res, player)
		}
	}
	return res
}
