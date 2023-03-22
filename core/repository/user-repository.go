package repository

import (
	"context"

	"api-google-keep/core/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user entity.UserCreate) (entity.UserCreate, error)
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

type userRepository struct {
	db *gorm.DB
}

func (r *userRepository) CreateUser(ctx context.Context, user entity.UserCreate) (entity.UserCreate, error) {
	rsl := r.db.Save(&user)
	return user, rsl.Error 
}