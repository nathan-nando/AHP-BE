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

func (r *CollectionRepositoryImpl) Create(ctx context.Context, e *model.CollectionModel) (*model.CollectionModel, error) {
	err := r.db.Create(e).
		WithContext(ctx).Error
	if err != nil {
		return nil, err
	}

	return e, nil
}

func (r *CollectionRepositoryImpl) Finds(ctx context.Context) ([]model.CollectionModel, error) {
	var datas []model.CollectionModel

	err := r.db.Find(&datas).WithContext(ctx).Error

	if err != nil {
		return datas, err
	}
	return datas, nil
}

func (r *CollectionRepositoryImpl) FindByID(ctx context.Context, id *string) (*model.CollectionModel, error) {
	var data model.CollectionModel

	err := r.db.Where("id = ?", id).First(&data).
		WithContext(ctx).Error

	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *CollectionRepositoryImpl) FindAlternatives(ctx context.Context, id *string) ([]model.AlternativeModel, error) {
	var datas []model.AlternativeModel

	err := r.db.Where("collection_id = ?", id).Find(&datas).WithContext(ctx).Error

	if err != nil {
		return datas, err
	}

	return datas, nil
}

func (r *CollectionRepositoryImpl) Update(ctx context.Context, id *string, e *model.CollectionModel) (*model.CollectionModel, error) {
	err := r.db.Model(e).Where("id = ?", id).Updates(e).
		WithContext(ctx).Error
	if err != nil {
		return nil, err
	}
	err = r.db.Model(e).Updates(e).
		WithContext(ctx).Error
	if err != nil {
		return nil, err
	}
	return e, nil
}

func (r *CollectionRepositoryImpl) Delete(ctx context.Context, id *string, e *model.CollectionModel) (*string, error) {
	err := r.db.Model(e).Where("id = ?", id).Delete(e).
		WithContext(ctx).Error
	if err != nil {
		return nil, err
	}
	return id, nil
}
