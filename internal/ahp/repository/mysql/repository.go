package mysql

import (
	"ahp-be/internal/model"
	"context"
	"gorm.io/gorm"
)

type AHPRepositoryImpl struct {
	Db *gorm.DB
}

func New(db *gorm.DB) *AHPRepositoryImpl {
	return &AHPRepositoryImpl{
		Db: db,
	}
}

func (r *AHPRepositoryImpl) CreateScore(ctx context.Context, e []model.ScoreModel) ([]model.ScoreModel, error) {
	err := r.Db.Create(&e).
		WithContext(ctx).Error
	if err != nil {
		return nil, err
	}

	return e, nil
}

func (r *AHPRepositoryImpl) CreateFinalScore(ctx context.Context, e []model.FinalScoreModel) ([]model.FinalScoreModel, error) {
	err := r.Db.Create(&e).
		WithContext(ctx).Error
	if err != nil {
		return nil, err
	}

	return e, nil
}

func (r *AHPRepositoryImpl) FindScoreByCollectionID(ctx context.Context, collectionID *string) ([]model.AlternativeModel, error) {
	var datas []model.AlternativeModel

	err := r.Db.Preload("Score").Where("collection_id = ?", collectionID).Find(&datas).WithContext(ctx).Error

	if err != nil {
		return datas, err
	}

	return datas, nil
}

func (r *AHPRepositoryImpl) FindFinalScoreByCollectionID(ctx context.Context, collectionID *string) ([]model.AlternativeModel, error) {
	var datas []model.AlternativeModel

	err := r.Db.Preload("FinalScore").Where("collection_id = ?", collectionID).Find(&datas).WithContext(ctx).Error

	if err != nil {
		return datas, err
	}

	return datas, nil
}

func (r *AHPRepositoryImpl) UpdateCollection(ctx context.Context, collectionID *string, m *model.CollectionModel) (*model.CollectionModel, error) {
	err := r.Db.Model(m).Where("id", collectionID).Updates(m).WithContext(ctx).Error

	if err != nil {
		return nil, err
	}

	return m, nil
}

func (r *AHPRepositoryImpl) DeleteScoresByCollectionID(ctx context.Context, collectionID *string) (*model.ScoreModel, error) {
	var data *model.ScoreModel

	err := r.Db.Where("collection_id = ?", collectionID).Delete(&data).WithContext(ctx).Error

	if err != nil {
		return data, err
	}

	return data, err
}

func (r *AHPRepositoryImpl) DeleteFinalScoresByCollectionID(ctx context.Context, collectionID *string) (*model.FinalScoreModel, error) {
	var data *model.FinalScoreModel

	err := r.Db.Where("collection_id = ?", collectionID).Delete(&data).WithContext(ctx).Error

	if err != nil {
		return data, err
	}

	return data, err
}
