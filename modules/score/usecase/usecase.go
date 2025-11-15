package usecase

import (
	"context"
	"errors"
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
	// ambil score
	s, err := uc.repo.GetByID(ctx, scoreID)
	if err != nil {
		return nil, err
	}
	if s == nil {
		return nil, errors.New("score not found")
	}

	// ambil semua pertanyaan terkait quiz
	questions, err := uc.qRepo.GetByQuizID(ctx, s.QuizID)
	if err != nil {
		return nil, err
	}

	// ubah answers slice jadi map[questionID] -> []answerID string
	subMap := make(map[string][]string)
	for _, sub := range answers {
		subMap[sub.QuestionID] = sub.AnswerIDs
	}

	var details []score_detail.ScoreDetail

	var totalScore float64
	var correctCount int
	var falseCount int

	for _, q := range questions {
		// ambil semua pilihan jawaban untuk question ini
		answerList, err := uc.aRepo.GetByQuestionID(ctx, q.ID)
		if err != nil {
			return nil, err
		}
		// kumpulkan ID jawaban yang benar
		correctIDsSet := map[uuid.UUID]struct{}{}
		for _, a := range answerList {
			if a.IsCorrect {
				correctIDsSet[a.ID] = struct{}{}
			}
		}
		totalCorrect := len(correctIDsSet)

		// jawaban user untuk question ini (string uuid)
		userAnswerStrs := subMap[q.ID.String()]

		// jika user tidak mengirim jawaban untuk question ini -> buat ScoreDetail dengan AnswerID = nil
		if len(userAnswerStrs) == 0 {
			details = append(details, score_detail.ScoreDetail{
				ScoreID:    scoreID,
				QuestionID: q.ID,
				AnswerID:   nil,
				IsCorrect:  false,
			})
			// question considered incorrect
			falseCount++
			continue
		}

		// untuk menghitung earned score per question:
		// hitung berapa banyak pilihan benar yang dipilih user (unique)
		correctSelectedCount := 0
		// track user selected ids to detect duplicates and extras
		userSelectedSet := map[uuid.UUID]struct{}{}
		for _, aStr := range userAnswerStrs {
			aID, err := uuid.Parse(aStr)
			if err != nil {
				// skip invalid uuid answer id (treat as wrong)
				continue
			}
			// prevent double counting duplicates
			if _, ok := userSelectedSet[aID]; ok {
				continue
			}
			userSelectedSet[aID] = struct{}{}

			_, isCorrect := correctIDsSet[aID]
			if isCorrect {
				correctSelectedCount++
			}
			// buat detail per answer user (AnswerID non-nil)
			tmpID := aID
			details = append(details, score_detail.ScoreDetail{
				ScoreID:    scoreID,
				QuestionID: q.ID,
				AnswerID:   &tmpID,
				IsCorrect:  isCorrect,
			})
		}

		// earned fraction: correctSelectedCount / totalCorrect (jika totalCorrect == 0 => 0)
		var earned float64
		if totalCorrect > 0 {
			earned = float64(correctSelectedCount) / float64(totalCorrect)
		} else {
			earned = 0
		}
		totalScore += earned

		// isQuestionCorrect: harus memilih semua correct AND tidak memilih pilihan incorrect (exact match)
		isQuestionCorrect := false
		if totalCorrect > 0 {
			if correctSelectedCount == totalCorrect && len(userSelectedSet) == totalCorrect {
				isQuestionCorrect = true
			}
		} else {
			// no correct answers configured => treat as incorrect
			isQuestionCorrect = false
		}
		if isQuestionCorrect {
			correctCount++
		} else {
			falseCount++
		}
	}

	// simpan semua details (bulk)
	if len(details) > 0 {
		if err := uc.sdRepo.BulkInsert(ctx, details); err != nil {
			return nil, err
		}
	}

	// update score
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
