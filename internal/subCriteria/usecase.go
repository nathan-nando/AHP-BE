package subCriteria

import (
	"ahp-be/internal/model"
	"ahp-be/internal/subCriteria/dto"
	"context"
)

type UseCase interface {
	FindSubCriteria(ctx context.Context) (*model.SubCriteria, error)
	UpdateSubCriteria(ctx context.Context, payload *model.SubCriteria) (*model.SubCriteria, error)
	CheckConsistencyRatio(ctx context.Context, mode *dto.CheckCRSubCriteria) (bool, error)
}
