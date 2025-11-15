package usecase

import (
	"context"
	"errors"

	"github.com/google/uuid"

	"quiz-app/modules/question"
)

type questionUsecase struct {
	repo question.QuestionRepository
}

func NewQuestionUsecase(repo question.QuestionRepository) question.QuestionUsecase {
	return &questionUsecase{repo}
}

func (uc *questionUsecase) Create(ctx context.Context, input question.CreateQuestionInput) (*question.Question, error) {
	q := &question.Question{
		Content: input.Content,
		Type:    input.Type,
		QuizID:  input.QuizID,
	}

	err := uc.repo.Create(ctx, q)
	return q, err
}

func (uc *questionUsecase) Update(ctx context.Context, id uuid.UUID, input question.UpdateQuestionInput) (*question.Question, error) {
	q, err := uc.repo.GetByID(ctx, id)
	if err != nil || q == nil {
		return nil, errors.New("question not found")
	}

	q.Content = input.Content
	q.Type = input.Type

	err = uc.repo.Update(ctx, q)
	return q, err
}

func (uc *questionUsecase) Delete(ctx context.Context, id uuid.UUID) error {
	return uc.repo.Delete(ctx, id)
}

func (uc *questionUsecase) GetByID(ctx context.Context, id uuid.UUID) (*question.Question, error) {
	return uc.repo.GetByID(ctx, id)
}

func (uc *questionUsecase) GetByQuizID(ctx context.Context, quizID uuid.UUID) ([]question.Question, error) {
	return uc.repo.GetByQuizID(ctx, quizID)
}
