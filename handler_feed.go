package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/hmuir28/goRSS/internal/database"
)

func (apiConfig ApiConfig) handleFeedCreation(w http.ResponseWriter, r *http.Request, user database.User) {
	type Parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}

	decoder := json.NewDecoder(r.Body)

	params := Parameters{}

	err := decoder.Decode(&params)

	if err != nil {
		respondWithErr(w, 500, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	feed, err := apiConfig.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.URL,
		UserID:    user.ID,
	})

	if err != nil {
		respondWithErr(w, 500, fmt.Sprintf("Error creating a new user: %v", err))
		return
	}

	respondWithJSON(w, 201, databaseFeedToFeed(feed))
}

func (apiConfig ApiConfig) handleGetFeeds(w http.ResponseWriter, r *http.Request) {

	feeds, err := apiConfig.DB.GetFeeds(r.Context())

	if err != nil {
		respondWithErr(w, 400, fmt.Sprintf("Couldn't get feeds %v", err))
		return
	}

	respondWithJSON(w, 200, databaseFeedsToFeeds(feeds))
}
