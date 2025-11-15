package usecase

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"quiz-app/modules/user"
)

type userUsecase struct {
	repo user.UserRepository
}

func NewUserUsecase(repo user.UserRepository) user.UserUsecase {
	return &userUsecase{repo}
}

func (uc *userUsecase) Register(ctx context.Context, input user.RegisterInput) (*user.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	u := &user.User{
		Username: input.Username,
		Email:    input.Email,
		Password: string(hashedPassword),
	}

	err = uc.repo.Create(ctx, u)
	return u, err
}

func (uc *userUsecase) Login(ctx context.Context, input user.LoginInput) (*user.User, error) {
	u, err := uc.repo.GetByEmail(ctx, input.Email)
	if err != nil || u == nil {
		return nil, errors.New("invalid email or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(input.Password)); err != nil {
		return nil, errors.New("invalid email or password")
	}

	return u, nil
}

func (uc *userUsecase) GetProfile(ctx context.Context, id uuid.UUID) (*user.User, error) {
	return uc.repo.GetByID(ctx, id)
}
