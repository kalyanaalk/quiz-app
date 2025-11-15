package score_detail

import (
	"context"

	"github.com/google/uuid"
)

type ScoreDetailRepository interface {
	BulkInsert(ctx context.Context, details []ScoreDetail) error
	GetByScoreID(ctx context.Context, scoreID uuid.UUID) ([]ScoreDetail, error)
}

type ScoreDetailUsecase interface {
	SaveDetails(ctx context.Context, scoreID uuid.UUID, submissions []DetailSubmission) error
	GetDetails(ctx context.Context, scoreID uuid.UUID) ([]ScoreDetail, error)
}

type DetailSubmission struct {
	QuestionID string   `json:"question_id"`
	AnswerIDs  []string `json:"answer_ids"`
}
