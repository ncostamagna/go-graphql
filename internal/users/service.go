package users

import (
	"context"
	"fmt"

	"github.com/ncostamagna/go-graphql/internal/model"
)

type (
	Service interface {
		Store(ctx context.Context, name, email string) (*model.User, error)
		GetAll(ctx context.Context, preload string) ([]*model.User, error)
	}

	service struct {
		repo Repository
	}
)

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) Store(ctx context.Context, name, email string) (*model.User, error) {
	fmt.Println("Service")
	user := &model.User{
		Name:  name,
		Email: email,
	}

	u, err := s.repo.Store(ctx, user)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (s *service) GetAll(ctx context.Context, preload string) ([]*model.User, error) {
	fmt.Println(preload)
	return s.repo.GetAll(ctx)
}
