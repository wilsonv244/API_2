package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

///MOVIE STRUCT
type Movie struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"Isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

//creando las propiedades del autor
type Author struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

//inicializando la variable
var movies []Movie

//GET
//ENDPOINT: http:localhost:8080/movies
func allMoviesEndPoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Aplication/json")
	json.NewEncoder(w).Encode(movies)
}

//GET
//ENDPOINT: http:localhost:8080/movies/{id}
//ENDPOINT: http:localhost:8080/movies/4
func findMovieEndPoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Aplication/json")
	params := mux.Vars(r) ///obtiene los paramentros
	//se hara un for para, el recorrido
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Movie{})
}

// POST
// ENDPOINT: http:localhost:8080/movies
func createMovieEndPoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Aplication/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(1000))
	movies = append(movies, movie) //agregramos un libro
	json.NewEncoder(w).Encode(movie)

}

// PUT
// ENDPOINT: http:localhost:8080/movies
func updateMovieEndPoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Aplication/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movies)
			movie.ID = params["id"]
			movies = append(movies, movie) //agregramos un libro
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
	json.NewEncoder(w).Encode(movies)
}

// DELETE
// ENDPOINT: http:localhost:8080/movies/{id}
func deleteMovieEndPoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Aplication/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}

	}
	json.NewEncoder(w).Encode(movies)
}
func main() {
	r := mux.NewRouter()
	movies = append(movies, Movie{ID: "1", Isbn: "448765", Title: "Harry Potter", Author: &Author{FirstName: "Wilson", LastName: "Vargas"}})
	movies = append(movies, Movie{ID: "2", Isbn: "7898765", Title: "Dante", Author: &Author{FirstName: "Paul", LastName: "Charaja"}})
	r.HandleFunc("/movies", allMoviesEndPoint).Methods("GET")
	r.HandleFunc("/movies/{id}", findMovieEndPoint).Methods("GET")
	r.HandleFunc("/movies", createMovieEndPoint).Methods("POST")
	r.HandleFunc("/movies", updateMovieEndPoint).Methods("UPDATE")
	r.HandleFunc("/movies/{id}", deleteMovieEndPoint).Methods("DELETE")
	if err := http.ListenAndServe(":3100", r); err != nil {
		log.Fatal(err)
	}
}
