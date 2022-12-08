package users

import (
	"context"
	"fmt"

	"github.com/ncostamagna/go-graphql/internal/model"
)

type (
	Controller interface {
		Store(ctx context.Context, request model.NewUser) (*model.User, error)
		GetAll(ctx context.Context, preload string) ([]*model.User, error)
	}

	controller struct {
		service Service
	}
)

func NewController(service Service) Controller {
	return &controller{
		service: service,
	}
}

func (c *controller) Store(ctx context.Context, request model.NewUser) (*model.User, error) {
	fmt.Println("Controller")

	u, err := c.service.Store(ctx, request.Name, request.Email)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (c *controller) GetAll(ctx context.Context, preload string) ([]*model.User, error) {
	return c.service.GetAll(ctx, preload)
}
