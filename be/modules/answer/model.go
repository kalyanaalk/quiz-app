package answer

import "github.com/google/uuid"

type Answer struct {
	ID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Content    string    `gorm:"not null" json:"content"`
	IsCorrect  bool      `gorm:"not null" json:"is_correct"`
	QuestionID uuid.UUID `gorm:"type:uuid;not null" json:"question_id"`
}

type CreateAnswerInput struct {
	Content    string    `json:"content"`
	IsCorrect  bool      `json:"is_correct"`
	QuestionID uuid.UUID `json:"question_id"`
}

type UpdateAnswerInput struct {
	Content   string `json:"content"`
	IsCorrect bool   `json:"is_correct"`
}
