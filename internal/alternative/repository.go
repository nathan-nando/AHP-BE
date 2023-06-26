package alternative

import (
	"ahp-be/internal/model"
	"context"
)

type Repository interface {
	Create(ctx context.Context, e *model.AlternativeModel) (*model.AlternativeModel, error)
	FindByID(ctx context.Context, id *string) (*model.AlternativeModel, error)
	FindsByCollectionID(ctx context.Context, collectionID *string) ([]model.AlternativeModel, error)
	Update(ctx context.Context, e *model.AlternativeModel, id *string) (*model.AlternativeModel, error)
	Delete(ctx context.Context, e *model.AlternativeModel, id *string) (*string, error)
}
