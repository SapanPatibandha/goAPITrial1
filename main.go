package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Movie struct {
	Id       string   `json:"movieid"`
	Name     string   `json:"moviename"`
	Director Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var movies []Movie

func main() {

	fmt.Println("API starts")

	//Create router
	r := mux.NewRouter()
	fmt.Println("Router started")

	//seed testing data
	movies = append(movies,
		Movie{
			Id:   "1",
			Name: "Aavtar 2",
			Director: Director{
				FirstName: "Sapan",
				LastName:  "Patibandha",
			},
		},
		Movie{
			Id:   "2",
			Name: "Eagle Eye",
			Director: Director{
				FirstName: "Mark",
				LastName:  "Dunston",
			},
		},
	)

	fmt.Printf("%v", movies)

	//handle routes
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<h1>Welcome to Golang API</h1>"))
	})
	r.HandleFunc("/movies", getAllMovies).Methods("GET")
	r.HandleFunc("/movie/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movie", createMovie).Methods("POST")
	r.HandleFunc("/movie/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movie/{id}", deleteMovie).Methods("DELETE")

	//liston to the port.
	log.Fatal(http.ListenAndServe(":8001", r))

}

func getAllMovies(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all movies.")

	w.Header().Set("Content-Type", "applicaiton/json")

	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get movie by ID")

	w.Header().Set("Content-Type", "applicaiton/json")

	params := mux.Vars(r)

	for _, movie := range movies {
		if movie.Id == params["id"] {
			json.NewEncoder(w).Encode(movie)
			return
		}
	}

	json.NewEncoder(w).Encode("no record found")
	return
}

func createMovie(w http.ResponseWriter, r *http.Request) {

}

func updateMovie(w http.ResponseWriter, r *http.Request) {

}

func deleteMovie(w http.ResponseWriter, r *http.Request) {

}
