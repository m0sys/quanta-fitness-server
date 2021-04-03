package main

import (
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/mhd53/quanta-fitness-server/api/auth"
	"github.com/mhd53/quanta-fitness-server/database/psql"
	"github.com/mhd53/quanta-fitness-server/graph"
	"github.com/mhd53/quanta-fitness-server/graph/generated"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()
	router.Use(auth.Middleware())

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(graph.NewResolver()))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	// Making sure to defer closing datbase connection...
	defer psql.DbConn.Close()

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
