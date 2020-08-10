package main

import (
	"cleanarchitecture-fes/src/adaptor/graphqlgen"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/rs/cors"
)

func main() {

	port := "19001"
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle(
		"/query",
		cors.New(cors.Options{
			AllowedOrigins:   []string{"*"},
			AllowedMethods:   []string{"GET", "POST"},
			AllowCredentials: false,
		}).Handler(handler.NewDefaultServer(
			graphqlgen.NewExecutableSchema(
				graphqlgen.Config{
					Resolvers: &graphQLResolver{},
				},
			),
		)),
	)
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
