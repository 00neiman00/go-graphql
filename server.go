package main

import (
	"github.com/go-pg/pg/v10"
	"github.com/neimen-95/go-graphql/postgres"
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

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	c := Config{Resolvers: &Resolver{
		meetupRepository: &postgres.MeetupRepository{DB: DB},
		userRepository:   &postgres.UserRepository{DB: DB},
	}}

	srv := handler.NewDefaultServer(NewExecutableSchema(c))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
