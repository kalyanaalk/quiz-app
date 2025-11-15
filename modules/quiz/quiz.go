package quiz

import (
	"context"

	"github.com/google/uuid"
)

type QuizRepository interface {
	Create(ctx context.Context, q *Quiz) error
	Update(ctx context.Context, q *Quiz) error
	Delete(ctx context.Context, id uuid.UUID) error
	GetByID(ctx context.Context, id uuid.UUID) (*Quiz, error)
	GetAll(ctx context.Context) ([]Quiz, error)
}

type QuizUsecase interface {
	Create(ctx context.Context, input Quiz) (*Quiz, error)
	Update(ctx context.Context, id uuid.UUID, input Quiz) (*Quiz, error)
	Delete(ctx context.Context, id uuid.UUID) error
	GetByID(ctx context.Context, id uuid.UUID) (*Quiz, error)
	GetAll(ctx context.Context) ([]Quiz, error)
}
