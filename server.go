package main

import (
	"bocasay-gql/graph"
	"bocasay-gql/graph/generated"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/urfave/cli/v2"
)

var defaultPort = "8080"

func launchServer() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func main() {
	app := cli.NewApp()
	app.Name = "Custom GraphQL Server"
	app.Usage = "A little graphql server for meteo"
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:    "port",
			Aliases: []string{"p"},
			Value:   "8080",
			Usage:   "port in wich the server will listen",
		},
	}
	app.Action = func(c *cli.Context) error {
		defaultPort = c.String("port")
		launchServer()
		return nil
	}

	app.Run(os.Args)
}
