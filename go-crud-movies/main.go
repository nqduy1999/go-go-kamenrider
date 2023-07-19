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


type Movie struct{
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct{
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
}


var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}


func getMovie(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item:= range movies{
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func deleteMovie(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item:= range movies{
		if item.ID == params["id"]{
			movies = append(movies[:index], movies[index + 1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func createMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(10000000))
	movies  = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item:= range movies{
		if item.ID == params["id"]{
			movies = append(movies[:index], movies[index + 1:]...) // remove current movie
			var movie Movie
			item = movie
			movies = append(movies, item)
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func main(){
	r := mux.NewRouter()
	movies = append(movies, Movie{ID: "1", Title: "Transformer", Isbn: "3232323", Director: &Director{FirstName: "Duy", LastName: "Nguyen"}})
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	fmt.Printf("Starting server at port 8080\n")
	log.Fatal(http.ListenAndServe(":8080", r))
}