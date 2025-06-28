package handlers

import (
	"encoding/json"
	"movies/models"
	"net/http"
)

type MovieHandler struct{}

func (h *MovieHandler) writeJSONResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (h *MovieHandler) GetTopMovies(w http.ResponseWriter, r *http.Request) {
	movies := []models.Movie{
		{
			ID:          1,
			TMDB_ID:     182,
			ReleaseYear: 2022,
			Genres:      []models.Genre{{ID: 1, Name: "Thriller"}},
			Keywords:    []string{},
			Casting:     []models.Actor{{ID: 1, FirstName: "Oscar"}},
		},
		{
			ID:          2,
			TMDB_ID:     189,
			ReleaseYear: 1996,
			Genres:      []models.Genre{{ID: 3, Name: "Green mile"}},
			Keywords:    []string{},
			Casting:     []models.Actor{{ID: 3, FirstName: "Koffee"}},
		},
	}

	h.writeJSONResponse(w, movies)
}

func (h *MovieHandler) GetRandomMovies(w http.ResponseWriter, r *http.Request) {
	movies := []models.Movie{
		{
			ID:          4,
			TMDB_ID:     165,
			ReleaseYear: 2021,
			Genres:      []models.Genre{{ID: 5, Name: "Alien"}},
			Keywords:    []string{},
			Casting:     []models.Actor{{ID: 7, FirstName: "Norris"}},
		},
	}

	h.writeJSONResponse(w, movies)
}