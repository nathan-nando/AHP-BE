package ahp

import (
	"ahp-be/internal/model"
	"context"
)

type Repository interface {
	CreateScore(ctx context.Context, e []model.ScoreModel) ([]model.ScoreModel, error)
	CreateFinalScore(ctx context.Context, e []model.FinalScoreModel) ([]model.FinalScoreModel, error)

	FindScoreByCollectionID(ctx context.Context, collectionID *string) ([]model.AlternativeModel, error)
	FindFinalScoreByCollectionID(ctx context.Context, collectionID *string) ([]model.AlternativeModel, error)

	UpdateCollection(ctx context.Context, collectionID *string, m *model.CollectionModel) (*model.CollectionModel, error)

	DeleteScoresByCollectionID(ctx context.Context, collectionID *string) (*model.ScoreModel, error)
	DeleteFinalScoresByCollectionID(ctx context.Context, collectionID *string) (*model.FinalScoreModel, error)
}
