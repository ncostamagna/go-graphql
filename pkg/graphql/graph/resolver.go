package graph

import (
	"github.com/ncostamagna/go-graphql/internal"
	"github.com/ncostamagna/go-graphql/pkg/graphql/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	todos   []*model.Todo
	users   []*model.User
	Service internal.Service
}
