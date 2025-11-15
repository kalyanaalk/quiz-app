package repository

import (
	"context"

	"quiz-app/modules/quiz"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type quizRepository struct {
	db *gorm.DB
}

func NewQuizRepository(db *gorm.DB) quiz.QuizRepository {
	return &quizRepository{db}
}

func (r *quizRepository) Create(ctx context.Context, q *quiz.Quiz) error {
	return r.db.WithContext(ctx).Create(q).Error
}

func (r *quizRepository) Update(ctx context.Context, q *quiz.Quiz) error {
	return r.db.WithContext(ctx).Save(q).Error
}

func (r *quizRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&quiz.Quiz{}, id).Error
}

func (r *quizRepository) GetByID(ctx context.Context, id uuid.UUID) (*quiz.Quiz, error) {
	var q quiz.Quiz
	err := r.db.WithContext(ctx).First(&q, id).Error
	if err != nil {
		return nil, err
	}
	return &q, nil
}

func (r *quizRepository) GetAll(ctx context.Context) ([]quiz.Quiz, error) {
	var list []quiz.Quiz
	err := r.db.WithContext(ctx).Find(&list).Error
	return list, err
}
