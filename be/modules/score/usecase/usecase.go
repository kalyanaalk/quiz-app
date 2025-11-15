package usecase

import (
	"context"
	"errors"
	"fmt"
	"time"

	"quiz-app/modules/answer"
	"quiz-app/modules/question"
	"quiz-app/modules/score"
	"quiz-app/modules/score_detail"

	"github.com/google/uuid"
)

type scoreUsecase struct {
	repo   score.ScoreRepository
	qRepo  question.QuestionRepository
	aRepo  answer.AnswerRepository
	sdRepo score_detail.ScoreDetailRepository
}

func NewScoreUsecase(repo score.ScoreRepository, qRepo question.QuestionRepository, aRepo answer.AnswerRepository, sdRepo score_detail.ScoreDetailRepository) score.ScoreUsecase {
	return &scoreUsecase{
		repo:   repo,
		qRepo:  qRepo,
		aRepo:  aRepo,
		sdRepo: sdRepo,
	}
}

func (uc *scoreUsecase) Start(ctx context.Context, userID, quizID uuid.UUID) (*score.Score, error) {
	s := &score.Score{
		UserID:  userID,
		QuizID:  quizID,
		StartAt: time.Now(),
	}
	err := uc.repo.Create(ctx, s)
	return s, err
}

func (uc *scoreUsecase) Submit(ctx context.Context, scoreID uuid.UUID, answers []score.AnswerSubmission) (*score.Score, error) {
	s, err := uc.repo.GetByID(ctx, scoreID)
	if err != nil {
		return nil, err
	}
	if s == nil {
		return nil, errors.New("score not found")
	}

	questions, err := uc.qRepo.GetByQuizID(ctx, s.QuizID)
	if err != nil {
		return nil, err
	}

	fmt.Println("questions: ", questions)

	subMap := make(map[string][]string)
	for _, sub := range answers {
		subMap[sub.QuestionID] = sub.AnswerIDs
	}

	fmt.Println("answers: ", answers)

	var details []score_detail.ScoreDetail

	var totalScore float64
	var correctCount int
	var falseCount int

	for _, q := range questions {
		answerList, err := uc.aRepo.GetByQuestionID(ctx, q.ID)
		if err != nil {
			return nil, err
		}
		correctIDsSet := map[uuid.UUID]struct{}{}
		for _, a := range answerList {
			if a.IsCorrect {
				correctIDsSet[a.ID] = struct{}{}
			}
		}
		totalCorrect := len(correctIDsSet)

		userAnswerStrs := subMap[q.ID.String()]

		if len(userAnswerStrs) == 0 {
			details = append(details, score_detail.ScoreDetail{
				ScoreID:   scoreID,
				AnswerID:  nil,
				IsCorrect: false,
			})
			falseCount++
			continue
		}

		correctSelectedCount := 0
		userSelectedSet := map[uuid.UUID]struct{}{}
		for _, aStr := range userAnswerStrs {
			aID, err := uuid.Parse(aStr)
			if err != nil {
				continue
			}
			if _, ok := userSelectedSet[aID]; ok {
				continue
			}
			userSelectedSet[aID] = struct{}{}

			_, isCorrect := correctIDsSet[aID]
			if isCorrect {
				correctSelectedCount++
			}
			tmpID := aID
			details = append(details, score_detail.ScoreDetail{
				ScoreID:   scoreID,
				AnswerID:  &tmpID,
				IsCorrect: isCorrect,
			})

			fmt.Println("tmpid: ", tmpID)

		}

		var earned float64
		if totalCorrect > 0 {
			earned = float64(correctSelectedCount) / float64(totalCorrect)
		} else {
			earned = 0
		}
		totalScore += earned

		isQuestionCorrect := false
		if totalCorrect > 0 {
			if correctSelectedCount == totalCorrect && len(userSelectedSet) == totalCorrect {
				isQuestionCorrect = true
			}
		} else {
			isQuestionCorrect = false
		}
		if isQuestionCorrect {
			correctCount++
		} else {
			falseCount++
		}
	}

	if len(details) > 0 {
		if err := uc.sdRepo.BulkInsert(ctx, details); err != nil {
			return nil, err
		}
	}

	s.FinishedAt = time.Now()
	s.TotalScore = totalScore
	s.CorrectCount = correctCount
	s.FalseCount = falseCount

	if err := uc.repo.Update(ctx, s); err != nil {
		return nil, err
	}

	return s, nil
}

func (uc *scoreUsecase) GetDetail(ctx context.Context, scoreID uuid.UUID) (*score.ScoreDetailAggregated, error) {
	s, err := uc.repo.GetByID(ctx, scoreID)
	if err != nil {
		return nil, err
	}

	resp := &score.ScoreDetailAggregated{
		Score:     s,
		QuizID:    s.QuizID.String(),
		Questions: []score.QuestionDetail{},
	}

	return resp, nil
}

func (uc *scoreUsecase) GetResultsForQuiz(ctx context.Context, quizID, userID uuid.UUID) ([]score.Score, error) {
	return uc.repo.GetByUserAndQuiz(ctx, userID, quizID)
}
