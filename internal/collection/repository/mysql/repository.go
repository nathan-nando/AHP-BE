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

func (c *CollectionRepositoryImpl) Create(ctx context.Context, e *model.CollectionModel) (*model.CollectionModel, error) {
	err := c.db.Create(e).
		WithContext(ctx).Error
	if err != nil {
		return nil, err
	}

	return e, nil
}

func (c *CollectionRepositoryImpl) Finds(ctx context.Context) ([]model.CollectionModel, error) {
	var datas []model.CollectionModel

	err := c.db.Find(&datas).WithContext(ctx).Error

	if err != nil {
		return datas, err
	}
	return datas, nil
}

func (c *CollectionRepositoryImpl) FindByID(ctx context.Context, id *string) (*model.CollectionModel, error) {
	var data model.CollectionModel

	err := c.db.Where("id = ?", id).First(&data).
		WithContext(ctx).Error

	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (c *CollectionRepositoryImpl) FindAlternatives(ctx context.Context, id *string) ([]model.AlternativeModel, error) {
	var datas []model.AlternativeModel

	err := c.db.Where("collection_id = ?", id).Find(&datas).WithContext(ctx).Error

	if err != nil {
		return datas, err
	}

	return datas, nil
}

func (c *CollectionRepositoryImpl) Update(ctx context.Context, id *string, e *model.CollectionModel) (*model.CollectionModel, error) {
	err := c.db.Model(e).Where("id = ?", id).Updates(e).
		WithContext(ctx).Error
	if err != nil {
		return nil, err
	}
	err = c.db.Model(e).Updates(e).
		WithContext(ctx).Error
	if err != nil {
		return nil, err
	}
	return e, nil
}

func (c *CollectionRepositoryImpl) Delete(ctx context.Context, id *string, e *model.CollectionModel) (*string, error) {
	err := c.db.Model(e).Where("id = ?", id).Delete(e).
		WithContext(ctx).Error
	if err != nil {
		return nil, err
	}
	return id, nil
}
