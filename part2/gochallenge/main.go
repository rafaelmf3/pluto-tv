package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/plutotv-go-challenge/handlers"
)

func main() {
	err := godotenv.Load("../.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	http.HandleFunc("/movies", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.GetMovies(w, r)
		case http.MethodPost:
			handlers.AddMovie(w, r)
		}
	})

	http.HandleFunc("/movies/", handlers.GetMovieByID)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
