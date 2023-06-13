package collection

import (
	"ahp-be/internal/model"
	"context"
)

type Repository interface {
	CreateCollection(ctx context.Context, collection *model.CollectionModel) (*model.CollectionModel, error)
	GetCollection(ctx context.Context) ([]model.CollectionModel, error)
	DeleteCollection(ctx context.Context, id *string) (*string, error)
}
