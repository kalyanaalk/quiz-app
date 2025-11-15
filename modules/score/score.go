package score

import (
	"context"

	"github.com/google/uuid"
)

type ScoreRepository interface {
	Create(ctx context.Context, s *Score) error
	Update(ctx context.Context, s *Score) error
	GetByID(ctx context.Context, id uuid.UUID) (*Score, error)
	GetByUserAndQuiz(ctx context.Context, userID, quizID uuid.UUID) ([]Score, error)
}

type ScoreUsecase interface {
	Start(ctx context.Context, userID, quizID uuid.UUID) (*Score, error)
	Submit(ctx context.Context, scoreID uuid.UUID, answers []AnswerSubmission) (*Score, error)
	GetDetail(ctx context.Context, scoreID uuid.UUID) (*ScoreDetailAggregated, error)
	GetResultsForQuiz(ctx context.Context, quizID, userID uuid.UUID) ([]Score, error)
}

type AnswerSubmission struct {
	QuestionID string   `json:"question_id"`
	AnswerIDs  []string `json:"answer_ids"`
}

type QuestionDetail struct {
	QuestionID     string   `json:"question_id"`
	Content        string   `json:"content"`
	CorrectAnswers []string `json:"correct_answers"`
	UserAnswers    []string `json:"user_answers"`
	IsCorrect      bool     `json:"is_correct"`
	EarnedScore    int      `json:"earned_score"`
}

type ScoreDetailAggregated struct {
	Score     *Score           `json:"score"`
	QuizID    string           `json:"quiz_id"`
	Questions []QuestionDetail `json:"questions"`
}
