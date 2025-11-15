package score

import (
	"time"

	"github.com/google/uuid"
)

type Score struct {
	ID           uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserID       uuid.UUID `json:"user_id"`
	QuizID       uuid.UUID `json:"quiz_id"`
	CorrectCount int       `json:"correct_count"`
	FalseCount   int       `json:"false_count"`
	TotalScore   float64   `json:"total_score"`
	StartAt      time.Time `json:"start_at"`
	FinishedAt   time.Time `json:"finished_at"`
}
