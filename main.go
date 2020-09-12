package main

import (
	"fmt"
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

	//repo, err := inmemory.New()
	//if err != nil {
	//	fmt.Print(err.Error())
	//	return
	//}
	//
	//uc, err := apiVersion1.New(repo)
	//if err != nil {
	//	fmt.Print(err.Error())
	//	return
	//}
	//
	//err = api.New(uc)
	//if err != nil {
	//	fmt.Print(err.Error())
	//	return
	//}
	//return

	fmt.Println("Hello World")
	http.HandleFunc("/",helloWorld)
	http.ListenAndServe(":5000",nil)

}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Akshay")
}