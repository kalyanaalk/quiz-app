package repository

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"quiz-app/modules/score"
)

type scoreRepository struct {
	db *gorm.DB
}

func NewScoreRepository(db *gorm.DB) score.ScoreRepository {
	return &scoreRepository{db}
}

func (r *scoreRepository) Create(ctx context.Context, s *score.Score) error {
	return r.db.WithContext(ctx).Create(s).Error
}

func (r *scoreRepository) Update(ctx context.Context, s *score.Score) error {
	return r.db.WithContext(ctx).Save(s).Error
}

func (r *scoreRepository) GetByID(ctx context.Context, id uuid.UUID) (*score.Score, error) {
	var s score.Score
	err := r.db.WithContext(ctx).First(&s, "id = ?", id).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &s, err
}

func (r *scoreRepository) GetByUserAndQuiz(ctx context.Context, userID, quizID uuid.UUID) ([]score.Score, error) {
	var result []score.Score
	err := r.db.WithContext(ctx).
		Where("user_id = ? AND quiz_id = ?", userID, quizID).
		Find(&result).Error
	return result, err
}
