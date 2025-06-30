package data

import "movies/models"

type MovirStorage interface {
	GetTopMovies() ([]models.Movie, error)
	GetRandomMovies() ([]models.Movie, error)
	//GetMovieById(id int) (models.Movie, error)
	//SearchMovies(name string) ([]models.Movie, error)
	//GetAllGenres() ([]models.Genre, error)
}