package mysql

import (
	"ahp-be/internal/model"
	"context"
	"gorm.io/gorm"
)

type CollectionRepositoryImpl struct {
	db *gorm.DB
}

func New(db *gorm.DB) *CollectionRepositoryImpl {
	return &CollectionRepositoryImpl{db: db}
}

func (r *CollectionRepositoryImpl) CreateCollection(ctx context.Context, collection *model.CollectionModel) (*model.CollectionModel, error) {
	//TODO implement me
	panic("implement me")
}

func (r *CollectionRepositoryImpl) GetCollection(ctx context.Context) ([]model.CollectionModel, error) {
	//TODO implement me
	panic("implement me")
}

func (r *CollectionRepositoryImpl) DeleteCollection(ctx context.Context, id *string) (*string, error) {
	//TODO implement me
	panic("implement me")
}
