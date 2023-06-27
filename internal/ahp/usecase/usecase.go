package usecase

import (
	"ahp-be/internal/ahp"
	"ahp-be/internal/ahp/dto"
	"ahp-be/internal/model"
	"context"
	"gorm.io/gorm"
)

type Service struct {
	Repository ahp.Repository
	Db         *gorm.DB
}

func New(repo ahp.Repository, db *gorm.DB) *Service {
	return &Service{
		Repository: repo,
		Db:         db,
	}
}

func (s *Service) FindCriteria(ctx context.Context) (*model.Criteria, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) FindScoresByCollectionID(ctx context.Context, collectionID *string) ([]model.AlternativeModel, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) FindFinalScoresByCollectionID(ctx context.Context, collectionID *string) ([]model.AlternativeModel, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) UpdateCriteria(ctx context.Context, payload *dto.CriteriaUpdateRequest) (*model.Criteria, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) CalculateAlternativeToPoint(ctx context.Context, collectionID *string) (model.Matrix, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) CalculateScore(ctx context.Context, collectionID *string) ([]model.ScoreModel, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) CalculateFinalScore(ctx context.Context, collectionID *string) ([]model.FinalScoreModel, error) {
	//TODO implement me
	panic("implement me")
}
