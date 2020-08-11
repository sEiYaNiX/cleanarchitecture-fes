package main

import (
	"cleanarchitecture-fes/src/adaptor/graphqlgen"
	"cleanarchitecture-fes/src/adaptor/repository"
	"cleanarchitecture-fes/src/usecase"
	"cleanarchitecture-fes/src/usecase/feseventinteractor"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/rs/cors"
)

func main() {
	args := os.Getenv("MYSQL_ARGS")
	if args == "" {
		// args = "cafes:cafes99@(localhost:3306)/cleanarchitecture_fes"
		panic("Please set MYSQL_ARGS(read the README.)")
	}

	var fesEventRepository usecase.FesEventRepository
	fesEventRepository, err := repository.New(args)
	if err != nil {
		panic(err)
	}
	fesEventInteractor := feseventinteractor.New(fesEventRepository)

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
					Resolvers: &graphQLResolver{
						fesEventInteractor,
					},
				},
			),
		)),
	)
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
