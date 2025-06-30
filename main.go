package main

import (
	"database/sql"
	"log"
	"movies/handlers"
	"movies/logger"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func initializeLogger() *logger.Logger {
	logInstance, err := logger.NewLogger("movies.log")
	if err != nil {
		log.Fatalf("Failed to initialize logger $v", err)
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

	// movie handler initializer
	movieHandler := handlers.MovieHandler{}
	http.HandleFunc("/api/movies/top", movieHandler.GetTopMovies)
	http.HandleFunc("/api/movies/random", movieHandler.GetRandomMovies)

	http.Handle("/", http.FileServer(http.Dir("public")))

	const addr = "localhost:8082"

	err = http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalf("Server failed: %v", err)
		logInstance.Error("Server failed", err)
	}

}
