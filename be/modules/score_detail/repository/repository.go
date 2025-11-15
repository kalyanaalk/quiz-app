package repository

import (
	"context"

	"quiz-app/modules/score_detail"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type scoreDetailRepository struct {
	db *gorm.DB
}

func NewScoreDetailRepository(db *gorm.DB) score_detail.ScoreDetailRepository {
	return &scoreDetailRepository{db}
}

func (r *scoreDetailRepository) BulkInsert(ctx context.Context, details []score_detail.ScoreDetail) error {
	return r.db.WithContext(ctx).Create(&details).Error
}

func (r *scoreDetailRepository) GetByScoreID(ctx context.Context, scoreID uuid.UUID) ([]score_detail.ScoreDetail, error) {
	var out []score_detail.ScoreDetail
	err := r.db.WithContext(ctx).
		Where("score_id = ?", scoreID).
		Find(&out).Error
	return out, err
}
