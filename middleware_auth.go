package main

import (
	"fmt"
	"net/http"

	"github.com/hmuir28/goRSS/internal/auth"
	"github.com/hmuir28/goRSS/internal/database"
)

type handleAuthenticatedUser func(http.ResponseWriter, *http.Request, database.User)

func (apiConfig *ApiConfig) middleAuth(authenticatedUserHandler handleAuthenticatedUser) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		apiKey, err := auth.GetAPIKey(r.Header)

		if err != nil {
			respondWithErr(w, 403, fmt.Sprintf("Auth error %v", err))
			return
		}

		user, err := apiConfig.DB.GetUserByAPIKey(r.Context(), apiKey)

		if err != nil {
			respondWithErr(w, 400, fmt.Sprintf("Couldn't get the user %v", err))
			return
		}

		authenticatedUserHandler(w, r, user)
	}
}
