package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

var movieInfoURL = getEnv("MOVIE_INFO_URL", "http://default-movie-info-url")
var movieRatingsURL = getEnv("MOVIE_RATINGS_URL", "http://default-movie-ratings-url")

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func getMovieDetails(w http.ResponseWriter, r *http.Request) {
	movieIDStr := r.URL.Query().Get("movie_id")
	movieID, err := strconv.Atoi(movieIDStr)
	if err != nil {
		http.Error(w, "Invalid movie ID", http.StatusBadRequest)
		return
	}

	// Make a request to the Movie Info service
	responseInfo, err := http.Get(fmt.Sprintf("%s/%d", movieInfoURL, movieID))
	if err != nil {
		http.Error(w, "Failed to fetch movie info", http.StatusInternalServerError)
		return
	}
	defer responseInfo.Body.Close()

	var movieInfo map[string]interface{}
	if err := json.NewDecoder(responseInfo.Body).Decode(&movieInfo); err != nil {
		http.Error(w, "Failed to decode movie info response", http.StatusInternalServerError)
		return
	}

	// Make a request to the Movie Ratings service
	responseRatings, err := http.Get(fmt.Sprintf("%s/%d", movieRatingsURL, movieID))
	if err != nil {
		http.Error(w, "Failed to fetch movie ratings", http.StatusInternalServerError)
		return
	}
	defer responseRatings.Body.Close()

	var movieRatings map[string]interface{}
	if err := json.NewDecoder(responseRatings.Body).Decode(&movieRatings); err != nil {
		http.Error(w, "Failed to decode movie ratings response", http.StatusInternalServerError)
		return
	}

	// Combine movie info and ratings
	movieDetails := map[string]interface{}{
		"movie_info":    movieInfo,
		"movie_ratings": movieRatings,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movieDetails)
}

func main() {
	http.HandleFunc("/api/movie", getMovieDetails)

	// Start the HTTP server
	fmt.Println("Server listening on :5000")
	if err := http.ListenAndServe(":5000", nil); err != nil {
		fmt.Printf("Server error: %s\n", err)
	}
}
