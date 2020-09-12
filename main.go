package main

import (
	"fmt"
	"github.com/ak-17/ipl_stats/delivery/api"
	"github.com/ak-17/ipl_stats/repository/database/inmemory"
	"github.com/ak-17/ipl_stats/usecase/apiVersion1"
	"net/http"
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
	//
	_,err = api.New(uc)
	if err != nil {
		fmt.Print(err.Error())
		return
	}


	http.ListenAndServe(":5000",nil)

}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Akshay")
}

func Health(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	return "OK", nil
}
