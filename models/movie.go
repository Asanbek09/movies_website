package models

type Movie struct {
	ID int
	TMDB_ID int
	Title string
	Tagline string
	ReleaseYear int
	Genres []Genre
	Overview *string // nullable
	Score *float32 // nullable
	Popularity *float32 // nullable
	Keywords []string
	Language string // nullable
	PosterURL *string // nullable
	TrailerURL *string // nullable
	Casting []Actor
}