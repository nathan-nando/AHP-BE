package collection

import (
	"ahp-be/internal/collection/dto"
	"ahp-be/internal/model"
	"context"
)

type UseCase interface {
	Create(ctx context.Context, payload *dto.CreateCollectionRequest) (*dto.CreateCollectionResponse, error)
	Finds(ctx context.Context) ([]model.CollectionModel, error)
	FindByID(ctx context.Context, payload *dto.FindCollectionByIDRequest) (*dto.FindCollectionByIDResponse, error)
	Update(ctx context.Context, payload *dto.UpdateCollectionRequest) (*dto.UpdateCollectionResponse, error)
	Delete(ctx context.Context, payload *dto.DeleteCollectionRequest) (*dto.DeleteCollectionResponse, error)
}
