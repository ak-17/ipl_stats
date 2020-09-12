package api

import (
	"fmt"
	"net/http"

	"github.com/ak-17/ipl_stats/usecase"
	"github.com/gorilla/mux"
)

func New(usecase usecase.Usecase) error {
	api := Api{
		usecase: usecase,
	}
	err := api.initHandlers()
	if err != nil {
		fmt.Printf("error initializing router. err:%s", err.Error())
		return err
	}
	return nil
}

func (api *Api) initHandlers() error {
	r := mux.NewRouter()

	r.Handle("/healthCheck", HandlerFunc(api.Health))
	r.Handle("/api/getMovieByTitle", HandlerFunc(api.GetMovieByTitle)).Methods(http.MethodGet)
	http.Handle("/", r)

	err := http.ListenAndServe(":3000", r)
	if err != nil {
		fmt.Printf("error initializing router. err:%s", err.Error())
		return err
	}
	return nil
}
