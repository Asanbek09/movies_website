package main

import (
	"log"
	"movies/handlers"
	"movies/logger"
	"net/http"
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

	movieHandler := handlers.MovieHandler{}
	http.HandleFunc("/api/movies/top", movieHandler.GetTopMovies)
	http.HandleFunc("/api/movies/random", movieHandler.GetRandomMovies)

	http.Handle("/", http.FileServer(http.Dir("public")))

	const addr = "localhost:8082"

	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalf("Server failed: %v", err)
		logInstance.Error("Server failed", err)
	}

}
