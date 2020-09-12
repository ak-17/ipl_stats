package api

import (
	"context"
	"errors"
	"net/http"
)

func (api *Api) GetMovieByTitle(w http.ResponseWriter, r *http.Request) (interface{}, error) {

	title := r.URL.Query().Get("title")
	if title == "" {
		return nil, errors.New("empty title")
	}

	return api.usecase.GetMovieByTitle(context.Background(), title)

}
