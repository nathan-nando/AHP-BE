package subCriteria

import (
	"ahp-be/internal/model"
	"context"
)

type Repository interface {
	FindSubCriteria(ctx context.Context) (*model.SubCriteria, error)
}
