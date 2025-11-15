package question

import "github.com/google/uuid"

type Question struct {
	ID      uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Content string    `gorm:"not null" json:"content"`
	Type    bool      `gorm:"not null" json:"type"`
	QuizID  uuid.UUID `gorm:"type:uuid;not null" json:"quiz_id"`
}

type CreateQuestionInput struct {
	Content string    `json:"content"`
	Type    bool      `json:"type"`
	QuizID  uuid.UUID `json:"quiz_id"`
}

type UpdateQuestionInput struct {
	Content string `json:"content"`
	Type    bool   `json:"type"`
}
