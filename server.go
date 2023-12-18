package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-pg/pg/v10"
	"github.com/neimen-95/go-graphql/dataloader"
	"github.com/neimen-95/go-graphql/graphql"
	customMiddleware "github.com/neimen-95/go-graphql/middleware"
	"github.com/neimen-95/go-graphql/postgres"
	"github.com/neimen-95/go-graphql/resolver"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	DB := postgres.New(&pg.Options{
		User:     "postgres",
		Password: "postgres",
		Database: "meetmeup",
	})

	defer DB.Close()
	DB.AddQueryHook(postgres.DbLogger{})
	userRepo := &postgres.UserRepository{DB: DB}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080"},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(customMiddleware.AuthMiddleware(userRepo))

	c := graphql.Config{Resolvers: &resolver.Resolver{
		MeetupRepository: &postgres.MeetupRepository{DB: DB},
		UserRepository:   userRepo,
	}}

	queryHandler := handler.NewDefaultServer(graphql.NewExecutableSchema(c))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", dataloader.DataloaderMiddleWare(DB, queryHandler))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
