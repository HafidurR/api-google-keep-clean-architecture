package module

import (
	"api-google-keep/core/entity"
	"api-google-keep/core/repository"
	"api-google-keep/libs/password"
	"context"
)

type UserUsecase interface {
	Register(ctx context.Context, user entity.UserCreate) (entity.UserCreate, error)
}

type userUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) UserUsecase {
	return &userUsecase{userRepo}
}


func (us *userUsecase) Register(ctx context.Context, user entity.UserCreate) (res entity.UserCreate, err error) {
	hash, _ := password.HashPassword(user.Password)
	user.Password = string(hash)

	res, err = us.userRepo.Register(ctx, user)
	return res, err
}