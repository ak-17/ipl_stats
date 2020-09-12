package inmemory

import (
	"context"
	"errors"
	"strings"

	"github.com/ak-17/ipl_stats/model"
)

func (im *inMemory) GetMovieByTitle(ctx context.Context, name string) (model.Movie, error) {

	for _, val := range im.Movies {
		if strings.TrimSpace(val.MovieTitle) == strings.TrimSpace(name) {
			return val, nil
		}
	}

	return model.Movie{}, errors.New("movie not found")

}
