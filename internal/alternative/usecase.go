package alternative

import (
	"ahp-be/internal/alternative/dto"
	"ahp-be/internal/model"
	"context"
)

type UseCase interface {
	Create(ctx context.Context, payload *dto.CreateAlternativeRequest) (*dto.CreateAlternativeResponse, error)
	FindByID(ctx context.Context, payload *dto.FindAlternativeByIDRequest) (*dto.FindAlternativeByIDResponse, error)
	FindByCollectionID(ctx context.Context, payload *dto.FindAlternativeByIDRequest) ([]model.AlternativeModel, error)
	Update(ctx context.Context, payload *dto.UpdateAlternativeRequest) (*dto.UpdateAlternativeResponse, error)
	Delete(ctx context.Context, payload *dto.DeleteAlternativeRequest) (*dto.DeleteAlternativeResponse, error)
}
