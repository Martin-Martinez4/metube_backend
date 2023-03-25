package main

import (
	db "github/Martin-Martinez4/metube_backend/config"
	"github/Martin-Martinez4/metube_backend/graph"
	directives "github/Martin-Martinez4/metube_backend/graph/directives"
	services "github/Martin-Martinez4/metube_backend/graph/services"
	customMiddleware "github/Martin-Martinez4/metube_backend/middleware"

	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

const defaultPort = "8080"

func main() {

	// Load .env file that is in the same directory
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	DB_URL := os.Getenv("DB_URL")

	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowOriginFunc:  AllowOriginFunc,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	DB := db.GetDB("postgres", DB_URL)
	defer DB.Close()

	authService := &services.AuthServiceSQL{DB: DB}
	profileService := &services.ProfileServiceSQL{DB: DB}
	videoService := &services.VideoServiceSQL{DB: DB}

	r.Use(customMiddleware.WithTokenCookie())
	r.Use(customMiddleware.WithWriter())

	c := graph.Config{Resolvers: &graph.Resolver{
		AuthService:    authService,
		VideoService:   videoService,
		ProfileService: profileService,
	}}

	c.Directives.Authorize = directives.Authorization

	queryHandler := handler.NewDefaultServer(graph.NewExecutableSchema(c))

	r.Handle("/", playground.Handler("GraphQL playground", "/query"))
	r.Handle("/query", graph.DatatloaderMiddleware(DB, queryHandler))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

func AllowOriginFunc(r *http.Request, origin string) bool {
	// if origin == "http://example.com" {
	// 	return true
	// }
	// return false

	return true
}
