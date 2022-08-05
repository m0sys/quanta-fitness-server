package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	// "github.com/m0sys/quanta-fitness-server/api/auth"
	// "github.com/m0sys/quanta-fitness-server/database/psql"
	"github.com/m0sys/quanta-fitness-server/internal/api/gql/graph"
	"github.com/m0sys/quanta-fitness-server/internal/api/gql/graph/generated"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()
	// router.Use(auth.Middleware())

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(graph.NewResolver()))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	// Making sure to defer closing datbase connection...
	// defer psql.DbConn.Close()

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
