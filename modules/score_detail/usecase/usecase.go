package usecase

import (
	"context"

	"quiz-app/modules/score_detail"

	"github.com/google/uuid"
)

type scoreDetailUsecase struct {
	repo score_detail.ScoreDetailRepository
}

func NewScoreDetailUsecase(repo score_detail.ScoreDetailRepository) score_detail.ScoreDetailUsecase {
	return &scoreDetailUsecase{repo}
}

func (u *scoreDetailUsecase) SaveDetails(ctx context.Context, scoreID uuid.UUID, submissions []score_detail.DetailSubmission) error {
	var list []score_detail.ScoreDetail

	for _, s := range submissions {
		qID, _ := uuid.Parse(s.QuestionID)
		for _, ans := range s.AnswerIDs {
			aID, _ := uuid.Parse(ans)
			list = append(list, score_detail.ScoreDetail{
				ScoreID:    scoreID,
				QuestionID: qID,
				AnswerID:   &aID,
			})
		}
	}

	return u.repo.BulkInsert(ctx, list)
}

func (u *scoreDetailUsecase) GetDetails(ctx context.Context, scoreID uuid.UUID) ([]score_detail.ScoreDetail, error) {
	return u.repo.GetByScoreID(ctx, scoreID)
}
