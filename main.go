package main

import (
	"database/sql"
	"log"
	"movies/data"
	"movies/handlers"
	"movies/logger"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func initializeLogger() *logger.Logger {
	logInstance, err := logger.NewLogger("movies.log")
	if err != nil {
		log.Fatalf("Failed to initialize logger %v", err)
	}
	defer logInstance.Close()

	return logInstance
}

func main() {

	logInstance := initializeLogger()

	// environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file was available")
	}

	// connect to the DB
	dbConnStr := os.Getenv("DATABASE_URL")
	if dbConnStr == "" {
		log.Fatal("DATABASE_URL is empty")
	}
	db, err := sql.Open("postgres", dbConnStr)
	if err != nil {
		log.Fatalf("Failed to connect to the DB: %v", err)
	}
	defer db.Close()

	// initialiaze data repository for movies
	movieRepo, err := data.NewMovieRepository(db, logInstance)
	if err != nil {
		log.Fatalf("Fail to initialize repository %v", err)
	}

	movieHandler := handlers.NewMovieHandler(movieRepo, logInstance)

	// movie handler initializer

	http.HandleFunc("/api/movies/top", movieHandler.GetTopMovies)
	http.HandleFunc("/api/movies/random/", movieHandler.GetRandomMovies)
	http.HandleFunc("/api/movies/search/", movieHandler.SearchMovies)
	http.HandleFunc("/api/movies/", movieHandler.GetMovie) // api/movies/140
	http.HandleFunc("/api/genres/", movieHandler.GetGenres)

	http.Handle("/", http.FileServer(http.Dir("public")))

	const addr = "localhost:8082"

	err = http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalf("Server failed: %v", err)
		logInstance.Error("Server failed", err)
	}

}
