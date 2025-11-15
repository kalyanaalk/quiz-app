package quiz

import (
	"github.com/google/uuid"
)

type Quiz struct {
	ID             uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	TotalQuestions int       `json:"total_questions"`
	Duration       int       `json:"duration"`
}
