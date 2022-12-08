package users

import (
	"github.com/ncostamagna/go-graphql/internal/model"
)

type DB struct {
	addresses []*model.Address
	users     []*model.User
}
