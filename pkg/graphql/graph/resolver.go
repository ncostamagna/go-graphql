package graph

import (
	"github.com/ncostamagna/go-graphql/internal/users"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	UserControl users.Controller
}
