package usecase

import (
	"context"

	"github.com/ak-17/ipl_stats/model"
)

type Usecase interface {
	GetMovieByTitle(ctx context.Context, name string) (model.Movie, error)
}
