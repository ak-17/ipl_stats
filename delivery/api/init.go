package api

import (
	"fmt"
	"net/http"

	"github.com/ak-17/ipl_stats/usecase"
	"github.com/gorilla/mux"
)

func New(usecase usecase.Usecase) (Api, error) {
	api := Api{
		usecase: usecase,
	}
	err := api.initHandlers()
	if err != nil {
		fmt.Printf("error initializing router. err:%s", err.Error())
		return api, err
	}
	return api, nil
}

func (api *Api) initHandlers() error {
	r := mux.NewRouter()

	r.Handle("/healthCheck", HandlerFunc(api.Health))
	r.Handle("/api/getTeamPoints", HandlerFunc(api.GetTeamPoints)).Methods(http.MethodGet)
	r.Handle("/api/getPlayersByTeam", HandlerFunc(api.GetPlayersByTeam)).Methods(http.MethodGet)
	r.Handle("/", HandlerFunc(api.Health))
	http.Handle("/", r)

	return nil
}
