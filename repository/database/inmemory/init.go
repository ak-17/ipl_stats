package inmemory

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/ak-17/ipl_stats/model"
	"github.com/ak-17/ipl_stats/repository"
)

func New() (repository.Repository, error) {

	movies, err := initMovies()
	if err != nil {
		log.Printf("[inmemory][New][Error] error occured during initialization. err:%s", err.Error())
		return nil, err
	}

	mem := &inMemory{
		Movies: movies,
	}

	return mem, nil

}

func initMovies() ([]model.Movie, error) {
	var movies []model.Movie
	file, err := ioutil.ReadFile("repository/database/inmemory/data.json")
	if err != nil {
		fmt.Printf(" error occured %s", err.Error())
		return movies, err
	}

	err = json.Unmarshal(file, &movies)
	if err != nil {
		fmt.Printf(" error occured %s", err.Error())
		return movies, err
	}
	fmt.Println(len(movies))
	return movies, nil

}
