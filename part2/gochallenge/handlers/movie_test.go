package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/plutotv-go-challenge/models"
)

func TestGetMovies(t *testing.T) {
	loadTestCache()
	req, err := http.NewRequest("GET", "/movies", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetMovies)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var movies []models.Movie
	if err := json.NewDecoder(rr.Body).Decode(&movies); err != nil {
		t.Errorf("Could not decode response: %v", err)
	}

	if len(movies) != 2 {
		t.Errorf("Expected 2 movies, got %d", len(movies))
	}
}

func TestAddMovie(t *testing.T) {
	newMovie := models.Movie{Title: "My Title", Director: "Director Test"}
	body, _ := json.Marshal(newMovie)
	req, err := http.NewRequest("POST", "/movies", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(AddMovie)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var movie models.Movie
	if err := json.NewDecoder(rr.Body).Decode(&movie); err != nil {
		t.Errorf("Could not decode response: %v", err)
	}

	if movie.Title != "My Title" {
		t.Errorf("Expected movie title 'My Title', got %v", movie.Title)
	}

	if movie.Director != "Director Test" {
		t.Errorf("Expected movie director 'Director Test', got %v", movie.Director)
	}
}

func TestGetMovieByID(t *testing.T) {
	loadTestCache()
	req, err := http.NewRequest("GET", "/movies/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetMovieByID)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var movie models.Movie
	if err := json.NewDecoder(rr.Body).Decode(&movie); err != nil {
		t.Errorf("Could not decode response: %v", err)
	}

	if movie.ID != 1 {
		t.Errorf("Expected movie ID 1, got %v", movie.ID)
	}

	if movie.Title != "The Lord of the Rings: The Fellowship of the Ring" {
		t.Errorf("Expected movie title 'The Lord of the Rings: The Fellowship of the Ring', got %v", movie.Title)
	}

	if movie.Director != "Peter Jackson" {
		t.Errorf("Expected movie director 'Peter Jackson', got %v", movie.Director)
	}
}

func TestGetMovieByIDNotFound(t *testing.T) {
	req, err := http.NewRequest("GET", "/movies/100", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetMovieByID)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}

	expected := "Movie not found\n"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestGetMovieByIDInvalidID(t *testing.T) {
	req, err := http.NewRequest("GET", "/movies/invalid", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetMovieByID)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

	expected := "Invalid movie ID\n"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func loadTestCache() {
	moviesList := []models.Movie{{
		ID:       1,
		Title:    "The Lord of the Rings: The Fellowship of the Ring",
		Director: "Peter Jackson",
		Year:     "2001",
	}, {
		ID:       2,
		Title:    "The Lord of the Rings: The Return of the King",
		Director: "Peter Jackson",
		Year:     "2003",
	}}

	for _, mv := range moviesList {
		moviesCache.Store(mv.ID, mv)
	}
}
