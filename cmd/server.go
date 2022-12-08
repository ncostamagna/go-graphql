package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/ncostamagna/go-graphql/internal/users"
	tools "github.com/ncostamagna/go-graphql/pkg/graphql"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db := users.DB{}
	usrRepo := users.NewRepo(db)
	usrSrv := users.NewService(usrRepo)
	usrCtl := users.NewController(usrSrv)

	handler := tools.NewTransport(usrCtl)

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", handler)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
