package seeder

import (
	"encoding/json"
	"log"
	"os"

	"quiz-app/config"
	"quiz-app/modules/answer"
	"quiz-app/modules/question"
	"quiz-app/modules/quiz"

	"github.com/google/uuid"
)

type JsonAnswer struct {
	Content   string `json:"content"`
	IsCorrect bool   `json:"is_correct"`
}

type JsonQuestion struct {
	Content string       `json:"content"`
	Type    bool         `json:"type"`
	Answers []JsonAnswer `json:"answers"`
}

type JsonQuiz struct {
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Duration    int            `json:"duration"`
	Questions   []JsonQuestion `json:"questions"`
}

func Seed() {
	db := config.DB

	var count int64
	db.Model(&quiz.Quiz{}).Count(&count)
	if count > 0 {
		log.Println("Seed skipped: data exists")
		return
	}

	jsonFile, err := os.ReadFile("seeder/quiz_data.json")
	if err != nil {
		log.Fatal("Error reading JSON:", err)
	}

	var quizzes []JsonQuiz
	err = json.Unmarshal(jsonFile, &quizzes)
	if err != nil {
		log.Fatal("Invalid JSON format:", err)
	}

	log.Println("Seeding quiz data...")

	for _, qz := range quizzes {
		newQuiz := quiz.Quiz{
			ID:             uuid.New(),
			Title:          qz.Title,
			Description:    qz.Description,
			Duration:       qz.Duration,
			TotalQuestions: len(qz.Questions),
		}
		db.Create(&newQuiz)

		for _, q := range qz.Questions {
			newQuestion := question.Question{
				ID:      uuid.New(),
				Content: q.Content,
				Type:    q.Type,
				QuizID:  newQuiz.ID,
			}
			db.Create(&newQuestion)

			for _, a := range q.Answers {
				newAnswer := answer.Answer{
					ID:         uuid.New(),
					Content:    a.Content,
					IsCorrect:  a.IsCorrect,
					QuestionID: newQuestion.ID,
				}
				db.Create(&newAnswer)
			}
		}
	}

	log.Println("Seeder completed.")
}
