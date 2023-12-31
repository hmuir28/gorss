package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/hmuir28/goRSS/internal/database"
)

func (apiConfig ApiConfig) handleUserCreation(w http.ResponseWriter, r *http.Request) {
	type Parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)

	params := Parameters{}

	err := decoder.Decode(&params)

	if err != nil {
		respondWithErr(w, 500, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	user, err := apiConfig.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})

	if err != nil {
		respondWithErr(w, 500, fmt.Sprintf("Error creating a new user: %v", err))
		return
	}

	respondWithJSON(w, 201, databaseUserToUser(user))
}

func (apiConfig ApiConfig) handleGetUserByAPIKey(w http.ResponseWriter, r *http.Request, user database.User) {
	respondWithJSON(w, 200, databaseUserToUser(user))
}
