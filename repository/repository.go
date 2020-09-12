package repository

import (
	"context"

	"github.com/ak-17/ipl_stats/model"
)

type Repository interface {
	GetMovieByTitle(ctx context.Context, name string) (model.Movie, error)
}
