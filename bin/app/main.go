package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"quiz-app/config"
	"quiz-app/seeder"

	"quiz-app/modules/user"
	user_handler "quiz-app/modules/user/handler"
	user_repository "quiz-app/modules/user/repository"
	user_usecase "quiz-app/modules/user/usecase"

	"quiz-app/modules/quiz"
	quiz_handler "quiz-app/modules/quiz/handler"
	quiz_repository "quiz-app/modules/quiz/repository"
	quiz_usecase "quiz-app/modules/quiz/usecase"

	"quiz-app/modules/score"
	score_handler "quiz-app/modules/score/handler"
	score_repository "quiz-app/modules/score/repository"
	score_usecase "quiz-app/modules/score/usecase"

	"quiz-app/modules/score_detail"
	score_detail_handler "quiz-app/modules/score_detail/handler"
	score_detail_repository "quiz-app/modules/score_detail/repository"
	score_detail_usecase "quiz-app/modules/score_detail/usecase"

	"quiz-app/modules/question"
	question_handler "quiz-app/modules/question/handler"
	question_repository "quiz-app/modules/question/repository"
	question_usecase "quiz-app/modules/question/usecase"

	"quiz-app/modules/answer"
	answer_handler "quiz-app/modules/answer/handler"
	answer_repository "quiz-app/modules/answer/repository"
	answer_usecase "quiz-app/modules/answer/usecase"
)

func main() {
	godotenv.Load()
	config.InitDB()

	config.DB.AutoMigrate(
		&user.User{},
		&quiz.Quiz{},
		&question.Question{},
		&answer.Answer{},
		&score.Score{},
		&score_detail.ScoreDetail{},
	)

	seeder.Seed()

	r := gin.Default()

	userRepo := user_repository.NewUserRepository(config.DB)
	userUC := user_usecase.NewUserUsecase(userRepo)
	userHandler := user_handler.NewUserHandler(userUC)
	user_handler.RegisterRoutes(r, userHandler)

	quizRepo := quiz_repository.NewQuizRepository(config.DB)
	quizUC := quiz_usecase.NewQuizUsecase(quizRepo)
	quizHandler := quiz_handler.NewQuizHandler(quizUC)
	quiz_handler.RegisterRoutes(r, quizHandler)

	scoreDetailRepo := score_detail_repository.NewScoreDetailRepository(config.DB)
	scoreDetailUC := score_detail_usecase.NewScoreDetailUsecase(scoreDetailRepo)
	scoreDetailHandler := score_detail_handler.NewScoreDetailHandler(scoreDetailUC)
	score_detail_handler.RegisterRoutes(r, scoreDetailHandler)

	questionRepo := question_repository.NewQuestionRepository(config.DB)
	questionUC := question_usecase.NewQuestionUsecase(questionRepo)
	questionHandler := question_handler.NewQuestionHandler(questionUC)
	question_handler.RegisterRoutes(r, questionHandler)

	answerRepo := answer_repository.NewAnswerRepository(config.DB)
	answerUC := answer_usecase.NewAnswerUsecase(answerRepo)
	answerHandler := answer_handler.NewAnswerHandler(answerUC)
	answer_handler.RegisterRoutes(r, answerHandler)

	scoreRepo := score_repository.NewScoreRepository(config.DB)
	scoreUC := score_usecase.NewScoreUsecase(scoreRepo, questionRepo, answerRepo, scoreDetailRepo)
	scoreHandler := score_handler.NewScoreHandler(scoreUC)
	score_handler.RegisterRoutes(r, scoreHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server running on port %s", port)
	r.Run(fmt.Sprintf(":%s", port))
}
