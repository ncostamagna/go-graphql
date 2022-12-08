package tools

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/ncostamagna/go-graphql/internal/users"
	"github.com/ncostamagna/go-graphql/pkg/graphql/graph"
)

func NewTransport(ctl users.Controller) *handler.Server {
	return handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		UserControl: ctl,
	}}))
}
