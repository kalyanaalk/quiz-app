package question

import (
	"context"

	"github.com/google/uuid"
)

type QuestionRepository interface {
	Create(ctx context.Context, q *Question) error
	Update(ctx context.Context, q *Question) error
	Delete(ctx context.Context, id uuid.UUID) error
	GetByID(ctx context.Context, id uuid.UUID) (*Question, error)
	GetByQuizID(ctx context.Context, quizID uuid.UUID) ([]Question, error)
}

type QuestionUsecase interface {
	Create(ctx context.Context, input CreateQuestionInput) (*Question, error)
	Update(ctx context.Context, id uuid.UUID, input UpdateQuestionInput) (*Question, error)
	Delete(ctx context.Context, id uuid.UUID) error
	GetByID(ctx context.Context, id uuid.UUID) (*Question, error)
	GetByQuizID(ctx context.Context, quizID uuid.UUID) ([]Question, error)
}
