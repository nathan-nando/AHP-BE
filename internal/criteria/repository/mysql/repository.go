package mysql

import (
	"ahp-be/internal/model"
	"context"
	"gorm.io/gorm"
)

type CriteriaRepositoryIMPL struct {
	Db *gorm.DB
}

func New(db *gorm.DB) *CriteriaRepositoryIMPL {
	return &CriteriaRepositoryIMPL{
		Db: db,
	}
}

func (r *CriteriaRepositoryIMPL) CreateScore(ctx context.Context, e []model.ScoreModel) ([]model.ScoreModel, error) {
	err := r.Db.Create(&e).
		WithContext(ctx).Error
	if err != nil {
		return nil, err
	}

	return e, nil
}

func (r *CriteriaRepositoryIMPL) CreateFinalScore(ctx context.Context, e []model.FinalScoreModel) ([]model.FinalScoreModel, error) {
	err := r.Db.Create(&e).
		WithContext(ctx).Error
	if err != nil {
		return nil, err
	}

	return e, nil
}

func (r *CriteriaRepositoryIMPL) FindScoreByCollectionID(ctx context.Context, collectionID *string) ([]model.AlternativeModel, error) {
	var datas []model.AlternativeModel

	err := r.Db.Preload("Score").Where("collection_id = ?", collectionID).Find(&datas).WithContext(ctx).Error

	if err != nil {
		return datas, err
	}

	return datas, nil
}

func (r *CriteriaRepositoryIMPL) FindFinalScoreByCollectionID(ctx context.Context, collectionID *string) ([]model.AlternativeModel, error) {
	var datas []model.AlternativeModel

	err := r.Db.Preload("FinalScore").Where("collection_id = ?", collectionID).Find(&datas).WithContext(ctx).Error

	if err != nil {
		return datas, err
	}

	return datas, nil
}

func (r *CriteriaRepositoryIMPL) UpdateCollection(ctx context.Context, collectionID *string, m *model.CollectionModel) (*model.CollectionModel, error) {
	err := r.Db.Model(m).Where("id", collectionID).Updates(m).WithContext(ctx).Error

	if err != nil {
		return nil, err
	}

	return m, nil
}

func (r *CriteriaRepositoryIMPL) DeleteScoresByCollectionID(ctx context.Context, collectionID *string) (*model.ScoreModel, error) {
	var data *model.ScoreModel

	err := r.Db.Where("collection_id = ?", collectionID).Delete(&data).WithContext(ctx).Error

	if err != nil {
		return data, err
	}

	return data, err
}

func (r *CriteriaRepositoryIMPL) DeleteFinalScoresByCollectionID(ctx context.Context, collectionID *string) (*model.FinalScoreModel, error) {
	var data *model.FinalScoreModel

	err := r.Db.Where("collection_id = ?", collectionID).Delete(&data).WithContext(ctx).Error

	if err != nil {
		return data, err
	}

	return data, err
}
