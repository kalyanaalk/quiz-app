package usecase

import (
	"context"

	"quiz-app/modules/quiz"

	"github.com/google/uuid"
)

type quizUsecase struct {
	repo quiz.QuizRepository
}

func NewQuizUsecase(repo quiz.QuizRepository) quiz.QuizUsecase {
	return &quizUsecase{repo}
}

func (uc *quizUsecase) Create(ctx context.Context, input quiz.Quiz) (*quiz.Quiz, error) {
	q := &quiz.Quiz{
		Title:          input.Title,
		Description:    input.Description,
		TotalQuestions: input.TotalQuestions,
		Duration:       input.Duration,
	}
	err := uc.repo.Create(ctx, q)
	return q, err
}

func (uc *quizUsecase) Update(ctx context.Context, id uuid.UUID, input quiz.Quiz) (*quiz.Quiz, error) {
	q, err := uc.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	q.Title = input.Title
	q.Description = input.Description
	q.TotalQuestions = input.TotalQuestions
	q.Duration = input.Duration

	err = uc.repo.Update(ctx, q)
	return q, err
}

func (uc *quizUsecase) Delete(ctx context.Context, id uuid.UUID) error {
	return uc.repo.Delete(ctx, id)
}

func (uc *quizUsecase) GetByID(ctx context.Context, id uuid.UUID) (*quiz.Quiz, error) {
	return uc.repo.GetByID(ctx, id)
}

func (uc *quizUsecase) GetAll(ctx context.Context) ([]quiz.Quiz, error) {
	return uc.repo.GetAll(ctx)
}
