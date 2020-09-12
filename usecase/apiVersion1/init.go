package apiVersion1

import (
	"github.com/ak-17/ipl_stats/repository"
	"github.com/ak-17/ipl_stats/usecase"
)

func New(repo repository.Repository) (usecase.Usecase, error) {
	api := apiVersion1{repo: repo}
	return api, nil
}
