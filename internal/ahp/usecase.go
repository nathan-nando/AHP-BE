package ahp

import (
	"ahp-be/internal/ahp/dto"
	"ahp-be/internal/model"
	"context"
)

type UseCase interface {
	FindCriteria(ctx context.Context) (*model.Criteria, error)
	FindScoresByCollectionID(ctx context.Context, collectionID *string) ([]model.AlternativeModel, error)
	FindFinalScoresByCollectionID(ctx context.Context, collectionID *string) ([]model.AlternativeModel, error)

	UpdateCriteria(ctx context.Context, payload *dto.CriteriaUpdateRequest) (*model.Criteria, error)

	CalculateAlternativeToPoint(ctx context.Context, collectionID *string) (model.Matrix, error)
	CalculateScore(ctx context.Context, collectionID *string) ([]model.ScoreModel, error)
	CalculateFinalScore(ctx context.Context, collectionID *string) ([]model.FinalScoreModel, error)
}
