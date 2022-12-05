package tools

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/ncostamagna/gokit-graphql/internal"
	"github.com/ncostamagna/gokit-graphql/pkg/graphql/graph"
)

func NewServer(srv internal.Service) *handler.Server {
	return handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		Service: srv,
	}}))
}
