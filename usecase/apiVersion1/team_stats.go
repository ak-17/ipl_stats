package apiVersion1

import (
	"context"
	"fmt"

	"github.com/ak-17/ipl_stats/model"
)

func (api *apiVersion1) GetFantasyPointsByTeam(ctx context.Context, teamNames []string) []model.Team {
	teamMap := make(map[string]float64)

	if len(teamNames) == 0 {
		fmt.Println("no team names")
		fmt.Printf("%v", model.TeamNames)
		teamNames = model.TeamNames
	}

	for _, val := range teamNames {
		players := api.repo.GetPlayersByTeam(ctx, val)
		for _, player := range players {
			teamMap[val] = teamMap[val] + player.Points
		}
	}

	var teams []model.Team
	for key, value := range teamMap {
		teams = append(teams, model.Team{Points: value, Name: key})
	}

	return teams
}
