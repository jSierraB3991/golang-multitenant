package service

import (
	"context"

	"github.com/jSierraB3991/golang-multitenant/models"
	repositoryinterface "github.com/jSierraB3991/golang-multitenant/repository_interface"
	"github.com/jSierraB3991/golang-multitenant/request"
	"github.com/jSierraB3991/golang-multitenant/response"
)

type UserService struct {
	userRepo repositoryinterface.UserRepositoryInterface
}

func NewUserService(userRepo repositoryinterface.UserRepositoryInterface) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) GetAllUsers(ctx context.Context) ([]response.UserResponse, error) {
	modelData, err := s.userRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	var result []response.UserResponse
	for _, v := range modelData {
		result = append(result, response.UserResponse{
			ID:    v.ID,
			Name:  v.Name,
			Email: v.Email,
		})
	}

	return result, nil
}

func (s *UserService) SaveUser(ctx context.Context, req request.UserRequest) error {
	user := models.User{
		Name:  req.Name,
		Email: req.Email,
	}

	return s.userRepo.SaveUser(ctx, &user)
}
