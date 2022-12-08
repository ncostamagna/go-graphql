package users

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/ncostamagna/go-graphql/internal/model"
)

type Repository interface {
	Store(ctx context.Context, user *model.User) (*model.User, error)
	GetAll(ctx context.Context) ([]*model.User, error)
}

type repo struct {
	db DB
}

func NewRepo(db DB) Repository {
	return &repo{
		db: db,
	}
}

func (r *repo) Store(ctx context.Context, user *model.User) (*model.User, error) {
	fmt.Println("repo")
	rand, _ := rand.Int(rand.Reader, big.NewInt(100))
	user.ID = fmt.Sprintf("T%d", rand)

	r.db.users = append(r.db.users, user)
	return user, nil
}

func (r *repo) GetAll(ctx context.Context) ([]*model.User, error) {
	return r.db.users, nil
}
