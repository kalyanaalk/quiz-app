package repository

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"quiz-app/modules/answer"
)

type answerRepository struct {
	db *gorm.DB
}

func NewAnswerRepository(db *gorm.DB) answer.AnswerRepository {
	return &answerRepository{db}
}

func (r *answerRepository) Create(ctx context.Context, a *answer.Answer) error {
	return r.db.WithContext(ctx).Create(a).Error
}

func (r *answerRepository) Update(ctx context.Context, a *answer.Answer) error {
	return r.db.WithContext(ctx).Save(a).Error
}

func (r *answerRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&answer.Answer{}, id).Error
}

func (r *answerRepository) GetByID(ctx context.Context, id uuid.UUID) (*answer.Answer, error) {
	var a answer.Answer
	err := r.db.WithContext(ctx).First(&a, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &a, err
}

func (r *answerRepository) GetByQuestionID(ctx context.Context, qid uuid.UUID) ([]answer.Answer, error) {
	var list []answer.Answer
	err := r.db.WithContext(ctx).Where("question_id = ?", qid).Find(&list).Error
	return list, err
}
