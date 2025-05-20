package serviceinterface

import (
	"context"

	"github.com/jSierraB3991/golang-multitenant/request"
	"github.com/jSierraB3991/golang-multitenant/response"
)

type UserServiceInterface interface {
	GetAllUsers(ctx context.Context) ([]response.UserResponse, error)
	SaveUser(ctx context.Context, req request.UserRequest) error
}
