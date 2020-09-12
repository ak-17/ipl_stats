package api

import (
	"errors"
	"net/http"
	"strings"
)

func (api *Api) GetPlayersByTeam(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	ctx := r.Context()

	team := r.FormValue("team")
	team = strings.TrimSpace(team)
	if team == "" {
		return nil, errors.New("empty team name")
	}

	players := api.usecase.GetPlayersByTeam(ctx, team)

	if len(players) == 0 {
		return nil, errors.New("team name not found")
	}

	return players, nil

}
