package repository

import (
	"context"

	"github.com/jSierraB3991/golang-multitenant/models"
)

func (r *Repository) FindAll(ctx context.Context) ([]models.User, error) {
	db, err := r.WithTenant(ctx)
	if err != nil {
		return nil, err
	}

	var users []models.User
	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *Repository) SaveUser(ctx context.Context, user *models.User) error {
	db, err := r.WithTenant(ctx)
	if err != nil {
		return err
	}

	if err := db.Create(user).Error; err != nil {
		return err
	}
	return nil
}
