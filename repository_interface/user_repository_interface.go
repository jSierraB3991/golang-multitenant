package repositoryinterface

import (
	"context"

	"github.com/jSierraB3991/golang-multitenant/models"
)

type UserRepositoryInterface interface {
	FindAll(ctx context.Context) ([]models.User, error)
	SaveUser(ctx context.Context, user *models.User) error
}
