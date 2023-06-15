package collection

import (
	"ahp-be/internal/model"
	"context"
)

type Repository interface {
	Create(ctx context.Context, e *model.CollectionModel) (*model.CollectionModel, error)
	Finds(ctx context.Context) ([]model.CollectionModel, error)
	FindByID(ctx context.Context, id *string) (*model.CollectionModel, error)
	FindAlternatives(ctx context.Context, id *string) ([]model.AlternativeModel, error)
	Update(ctx context.Context, id *string, e *model.CollectionModel) (*model.CollectionModel, error)
	Delete(ctx context.Context, id *string, e *model.CollectionModel) (*string, error)
}
