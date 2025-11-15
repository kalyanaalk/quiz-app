package repository

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"quiz-app/modules/question"
)

type questionRepository struct {
	db *gorm.DB
}

func NewQuestionRepository(db *gorm.DB) question.QuestionRepository {
	return &questionRepository{db}
}

func (r *questionRepository) Create(ctx context.Context, q *question.Question) error {
	return r.db.WithContext(ctx).Create(q).Error
}

func (r *questionRepository) Update(ctx context.Context, q *question.Question) error {
	return r.db.WithContext(ctx).Save(q).Error
}

func (r *questionRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&question.Question{}, id).Error
}

func (r *questionRepository) GetByID(ctx context.Context, id uuid.UUID) (*question.Question, error) {
	var q question.Question
	err := r.db.WithContext(ctx).First(&q, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &q, err
}

func (r *questionRepository) GetByQuizID(ctx context.Context, quizID uuid.UUID) ([]question.Question, error) {
	var list []question.Question
	err := r.db.WithContext(ctx).Where("quiz_id = ?", quizID).Find(&list).Error
	return list, err
}
