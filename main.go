package main

import (
	"fmt"

	"github.com/ak-17/ipl_stats/delivery/api"
	"github.com/ak-17/ipl_stats/usecase/apiVersion1"

	"github.com/ak-17/ipl_stats/repository/database/inmemory"
)

func main() {
	//file, err := ioutil.ReadFile("data.json")
	//if err != nil {
	//	fmt.Printf(" error occured %s", err.Error())
	//	return
	//}
	//
	//var movies []model.Movie
	//
	//err = json.Unmarshal(file, &movies)
	//if err != nil {
	//	fmt.Printf(" error occured %s", err.Error())
	//	return
	//}
	//
	//fmt.Print(len(movies))

	repo, err := inmemory.New()
	if err != nil {
		fmt.Print(err.Error())
		return
	}

	uc, err := apiVersion1.New(repo)
	if err != nil {
		fmt.Print(err.Error())
		return
	}

	err = api.New(uc)
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	return

}