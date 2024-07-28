package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"sync"

	"github.com/plutotv-go-challenge/models"
	"github.com/plutotv-go-challenge/pkg"
)

var movies = []models.Movie{}
var moviesCache sync.Map

func init() {
	mvs, _ := pkg.LoadMoviesFromIMDB("lord") //Example, but we can use any term to retrieve data from omdbapi

	movies = append(movies, mvs...)
	for _, movie := range movies {
		moviesCache.Store(movie.ID, movie)
	}
}

func GetMovies(w http.ResponseWriter, r *http.Request) {
	var moviesList []models.Movie
	moviesCache.Range(func(key, value interface{}) bool {
		moviesList = append(moviesList, value.(models.Movie))
		return true
	})
	json.NewEncoder(w).Encode(moviesList)
}

func GetMovieByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Path[len("/movies/"):])
	if err != nil {
		http.Error(w, "Invalid movie ID", http.StatusBadRequest)
		return
	}

	if movie, ok := moviesCache.Load(id); ok {
		json.NewEncoder(w).Encode(movie)
		return
	}
	http.Error(w, "Movie not found", http.StatusNotFound)
}

func AddMovie(w http.ResponseWriter, r *http.Request) {
	var newMovie models.Movie
	if err := json.NewDecoder(r.Body).Decode(&newMovie); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	newMovie.ID = len(movies) + 1
	movies = append(movies, newMovie)
	moviesCache.Store(newMovie.ID, newMovie)
	json.NewEncoder(w).Encode(newMovie)
}
