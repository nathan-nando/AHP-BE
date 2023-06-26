package mysql

import (
	"ahp-be/internal/model"
	"context"
	"gorm.io/gorm"
)

type AlternativeRepositoryImpl struct {
	db *gorm.DB
}

func New(db *gorm.DB) *AlternativeRepositoryImpl {
	return &AlternativeRepositoryImpl{db: db}
}

func (r *AlternativeRepositoryImpl) Create(ctx context.Context, e *model.AlternativeModel) (*model.AlternativeModel, error) {
	err := r.db.Create(e).
		WithContext(ctx).Error
	if err != nil {
		return nil, err
	}
	err = r.db.Model(e).First(e).
		WithContext(ctx).Error
	if err != nil {
		return nil, err
	}
	return e, nil
}

func (r *AlternativeRepositoryImpl) FindByID(ctx context.Context, id *string) (*model.AlternativeModel, error) {
	var data model.AlternativeModel

	err := r.db.Where("id = ?", id).First(&data).
		WithContext(ctx).Error

	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *AlternativeRepositoryImpl) FindsByCollectionID(ctx context.Context, collectionID *string) ([]model.AlternativeModel, error) {
	var datas []model.AlternativeModel

	err := r.db.Where("collection_id = ?", collectionID).Find(&datas).
		WithContext(ctx).Error
	if err != nil {
		return datas, err
	}

	return datas, nil
}

func (r *AlternativeRepositoryImpl) Update(ctx context.Context, e *model.AlternativeModel, id *string) (*model.AlternativeModel, error) {
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

func (r *AlternativeRepositoryImpl) Delete(ctx context.Context, e *model.AlternativeModel, id *string) (*string, error) {
	err := r.db.Model(e).Where("id = ?", id).Delete(e).
		WithContext(ctx).Error
	if err != nil {
		return nil, err
	}
	return id, nil
}
