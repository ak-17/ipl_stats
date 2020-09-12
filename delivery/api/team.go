package api

import (
	"errors"
	"net/http"
	"strings"
)

func (api *Api) GetTeamPoints(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	ctx := r.Context()
	teamsQuery := r.FormValue("team")
	teamsQuery = strings.TrimSpace(teamsQuery)

	var teams []string
	if teamsQuery != "" {
		teams = strings.Split(teamsQuery, ",")
	}

	res := api.usecase.GetFantasyPointsByTeam(ctx, teams)
	if len(res) == 0 {
		return nil, errors.New("invalid team name")
	}

	return res, nil
}
