package answer

import (
	"context"

	"github.com/google/uuid"
)

type AnswerRepository interface {
	Create(ctx context.Context, a *Answer) error
	Update(ctx context.Context, a *Answer) error
	Delete(ctx context.Context, id uuid.UUID) error
	GetByID(ctx context.Context, id uuid.UUID) (*Answer, error)
	GetByQuestionID(ctx context.Context, qid uuid.UUID) ([]Answer, error)
}

type AnswerUsecase interface {
	Create(ctx context.Context, input CreateAnswerInput) (*Answer, error)
	Update(ctx context.Context, id uuid.UUID, input UpdateAnswerInput) (*Answer, error)
	Delete(ctx context.Context, id uuid.UUID) error
	GetByID(ctx context.Context, id uuid.UUID) (*Answer, error)
	GetByQuestionID(ctx context.Context, qid uuid.UUID) ([]Answer, error)
}
