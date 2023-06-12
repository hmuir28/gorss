package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/hmuir28/goRSS/internal/database"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

type ApiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load(".env")

	portString := os.Getenv("PORT")

	if portString == "" {
		log.Fatal("PORT was not found in the env vars")
	}

	dbURL := os.Getenv("DB_URL")

	if dbURL == "" {
		log.Fatal("DB_URL was not found in the env vars")
	}

	conn, err := sql.Open("postgres", dbURL)

	if err != nil {
		log.Fatal("Can't connect to database", err)
	}

	queries := database.New(conn)

	apiConfig := ApiConfig{
		DB: queries,
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/health", handleReadiness)
	v1Router.Get("/err", handleError)

	v1Router.Post("/user", apiConfig.handleUserCreation)
	v1Router.Get("/user", apiConfig.middleAuth(apiConfig.handleGetUserByAPIKey))

	v1Router.Post("/feed", apiConfig.middleAuth(apiConfig.handleFeedCreation))
	v1Router.Get("/feeds", apiConfig.handleGetFeeds)

	v1Router.Post("/feed_follows", apiConfig.middleAuth(apiConfig.handleFeedFollowCreation))
	v1Router.Get("/feed_follows", apiConfig.middleAuth(apiConfig.handleGetFeedFollows))
	v1Router.Delete("/feed_follows/{feedFollowId}", apiConfig.middleAuth(apiConfig.handleDeleteFeedFollow))

	router.Mount("/v1", v1Router)

	server := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	log.Printf("Server starting on port %v", portString)

	err = server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}
