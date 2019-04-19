package main

import (
	"github.com/victorneuret/GitSync/app"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	"github.com/victorneuret/GitSync/database"
	"github.com/victorneuret/GitSync/generated"
	"github.com/victorneuret/GitSync/resolver"
	"github.com/victorneuret/GitSync/githubLogin"
	"github.com/victorneuret/GitSync/config"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	config.LoadConfiguration()

	database.ConnectDatabase()
	defer database.CloseDatabase()
	database.InitialMigration()

	if !config.Config.Prod {
		http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	}
	http.Handle("/query", handler.GraphQL(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver.Resolver{}})))

	githubLogin.Setup()
	http.HandleFunc("/hook", app.HandleWebHook)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}