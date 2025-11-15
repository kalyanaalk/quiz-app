package score_detail

import (
	"github.com/google/uuid"
)

type ScoreDetail struct {
	ID        uuid.UUID  `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	ScoreID   uuid.UUID  `json:"score_id"`
	AnswerID  *uuid.UUID `json:"answer_id"`
	IsCorrect bool       `json:"is_correct"`
}
