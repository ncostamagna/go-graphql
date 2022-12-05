package tools

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/ncostamagna/go-graphql/internal"
	"github.com/ncostamagna/go-graphql/pkg/graphql/graph"
)

func NewServer(srv internal.Service) *handler.Server {
	return handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		Service: srv,
	}}))
}
