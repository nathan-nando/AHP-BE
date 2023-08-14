package criteria

import (
	"ahp-be/internal/criteria/dto"
	"ahp-be/internal/model"
	"context"
)

type UseCase interface {
	FindCriteria(ctx context.Context) (*model.Pairwise, error)
	UpdateCriteria(ctx context.Context, payload *dto.CriteriaUpdateRequest) (*model.Pairwise, error)
	CheckConsistencyRatio(ctx context.Context) (bool, error)

	FindScoresByCollectionID(ctx context.Context, collectionID *string) ([]model.AlternativeModel, error)
	FindFinalScoresByCollectionID(ctx context.Context, collectionID *string) ([]model.AlternativeModel, error)

	CalculateAlternativeToPoint(ctx context.Context, collectionID *string) (model.Matrix, error)
	CalculateScore(ctx context.Context, collectionID *string) ([]model.ScoreModel, error)
	CalculateFinalScore(ctx context.Context, collectionID *string) ([]model.FinalScoreModel, error)
}
