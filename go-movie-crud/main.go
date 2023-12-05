package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID   string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"direction"`
}

type Director struct{
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-type", "applicaton/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, request *http.Request){
	w.Header().Set("Content-type", "applicaton/json")
	params := mux.Vars(request);
	for index, item := range movies{
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[:index + 1]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, request *http.Request){
	w.Header().Set("Content-type", "applicaton/json")
	params := mux.Vars(request);
	for _, item := range movies{
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, request *http.Request){
	w.Header().Set("Content-type", "applicaton/json")
	var movie Movie
	_ = json.NewDecoder(request.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(1000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
	return
}

func updateMovie(w http.ResponseWriter, request *http.Request){
	w.Header().Set("Content-type", "applicaton/json")
	params := mux.Vars(request);
	for index, item := range movies{
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[:index + 1]...)
			var movie Movie
			_ = json.NewDecoder(request.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}


func main() {
	movies = append(movies, Movie{ID: "1", Isbn: "43771", Title: "Title 124", Director: &Director{
		FirstName: "Duy",
		LastName: "Nguyen",
	}})
	movies = append(movies, Movie{ID: "2", Isbn: "43772", Title: "Title 125", Director: &Director{
		FirstName: "Duy",
		LastName: "Nguyen Quoc",
	}})
	r := mux.NewRouter()
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movie/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movie", createMovie).Methods("POST")
	r.HandleFunc("/movie/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movie/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
