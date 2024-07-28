package pkg

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/plutotv-go-challenge/models"
)

type imdbMovie struct {
	Title  string `json:"Title"`
	Year   string `json:"Year"`
	ImdbID string `json:"imdbID"`
}

type ImdbMovieDetail struct {
	Title      string `json:"Title"`
	Year       string `json:"Year"`
	ImdbID     string `json:"imdbID"`
	Director   string `json:"Director"`
	Runtime    string `json:"Runtime"`
	ImdbRating string `json:"imdbRating"`
}

type IMBDSearchResponse struct {
	Search []imdbMovie `json:"Search"`
}

func LoadMoviesFromIMDB(term string) ([]models.Movie, error) {

	imdbURL, _ := url.Parse("http://www.omdbapi.com")

	values := imdbURL.Query()
	values.Add("apikey", os.Getenv("IMDB_API_KEY"))
	values.Add("type", "movie")
	values.Add("s", term)

	movies := []models.Movie{}

	imdbURL.RawQuery = values.Encode()
	resp, err := http.Get(imdbURL.String())
	if err != nil {
		return movies, err
	}
	resBody, _ := io.ReadAll(resp.Body)
	response := &IMBDSearchResponse{}
	json.Unmarshal(resBody, response)

	for i, v := range response.Search {

		movieDetails, err := GetMovieDetailsFromIMDB(v.ImdbID)
		if err != nil {
			continue
		}
		movie := models.Movie{
			ID:       i + 1,
			Title:    v.Title,
			Year:     v.Year,
			Director: movieDetails.Director,
		}
		movies = append(movies, movie)
	}
	return movies, nil
}

func GetMovieDetailsFromIMDB(imdbID string) (models.Movie, error) {
	movies := models.Movie{}

	imdbURL, _ := url.Parse("http://www.omdbapi.com")
	values := imdbURL.Query()
	values.Add("apikey", os.Getenv("IMDB_API_KEY"))
	values.Add("i", imdbID)
	imdbURL.RawQuery = values.Encode()

	resp, err := http.Get(imdbURL.String())
	if err != nil {
		return movies, err
	}
	resBody, _ := io.ReadAll(resp.Body)
	response := &ImdbMovieDetail{}
	json.Unmarshal(resBody, response)

	movies.Director = response.Director

	return movies, nil
}
