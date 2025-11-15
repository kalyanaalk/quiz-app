package usecase

import (
	"context"
	"errors"

	"github.com/google/uuid"

	"quiz-app/modules/answer"
)

type answerUsecase struct {
	repo answer.AnswerRepository
}

func NewAnswerUsecase(repo answer.AnswerRepository) answer.AnswerUsecase {
	return &answerUsecase{repo}
}

func (uc *answerUsecase) Create(ctx context.Context, input answer.CreateAnswerInput) (*answer.Answer, error) {
	a := &answer.Answer{
		Content:    input.Content,
		IsCorrect:  input.IsCorrect,
		QuestionID: input.QuestionID,
	}

	err := uc.repo.Create(ctx, a)
	return a, err
}

func (uc *answerUsecase) Update(ctx context.Context, id uuid.UUID, input answer.UpdateAnswerInput) (*answer.Answer, error) {
	a, err := uc.repo.GetByID(ctx, id)
	if err != nil || a == nil {
		return nil, errors.New("answer not found")
	}

	a.Content = input.Content
	a.IsCorrect = input.IsCorrect

	err = uc.repo.Update(ctx, a)
	return a, err
}

func (uc *answerUsecase) Delete(ctx context.Context, id uuid.UUID) error {
	return uc.repo.Delete(ctx, id)
}

func (uc *answerUsecase) GetByID(ctx context.Context, id uuid.UUID) (*answer.Answer, error) {
	return uc.repo.GetByID(ctx, id)
}

func (uc *answerUsecase) GetByQuestionID(ctx context.Context, qid uuid.UUID) ([]answer.Answer, error) {
	return uc.repo.GetByQuestionID(ctx, qid)
}
