package main

import (
	"github.com/go-pg/pg/v10"
	"github.com/neimen-95/go-graphql/dataloader"
	"github.com/neimen-95/go-graphql/graphql"
	"github.com/neimen-95/go-graphql/postgres"
	"github.com/neimen-95/go-graphql/resolver"
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
	c := graphql.Config{Resolvers: &resolver.Resolver{
		MeetupRepository: &postgres.MeetupRepository{DB: DB},
		UserRepository:   &postgres.UserRepository{DB: DB},
	}}

	queryHandler := handler.NewDefaultServer(graphql.NewExecutableSchema(c))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", dataloader.DataloaderMiddleWare(DB, queryHandler))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
