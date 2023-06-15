package usecase

import (
	"ahp-be/internal/collection"
	"ahp-be/internal/collection/dto"
	"ahp-be/internal/model"
	"ahp-be/pkg/response"
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Service struct {
	repo collection.Repository
	db   *gorm.DB
}

func New(repo collection.Repository, db *gorm.DB) *Service {
	return &Service{
		repo: repo,
		db:   db,
	}
}

func (s *Service) Create(ctx context.Context, payload *dto.CreateCollectionRequest) (*dto.CreateCollectionResponse, error) {
	var result *dto.CreateCollectionResponse
	var data *model.CollectionModel

	data = &model.CollectionModel{
		BaseModel:  model.BaseModel{ID: uuid.NewString()},
		Collection: payload.Collection,
	}

	_, err := s.repo.Create(ctx, data)

	if err != nil {
		return nil, response.ErrorBuilder(&response.ErrorConstant.UnprocessableEntity, err)
	}

	result = &dto.CreateCollectionResponse{
		CollectionModel: *data,
	}

	return result, nil
}

func (s *Service) Finds(ctx context.Context) ([]model.CollectionModel, error) {
	result := make([]model.CollectionModel, 0)

	data, err := s.repo.Finds(ctx)

	if err != nil {
		return nil, response.ErrorBuilder(&response.ErrorConstant.UnprocessableEntity, err)
	}

	result = append(result, data...)

	return result, nil
}

func (s *Service) FindByID(ctx context.Context, payload *dto.FindCollectionByIDRequest) (*dto.FindCollectionByIDResponse, error) {
	var result *dto.FindCollectionByIDResponse

	data, err := s.repo.FindByID(ctx, &payload.ID)

	if err != nil {
		return nil, response.ErrorBuilder(&response.ErrorConstant.UnprocessableEntity, err)
	}

	result = &dto.FindCollectionByIDResponse{
		CollectionModel: *data,
	}

	return result, nil
}

func (s *Service) Update(ctx context.Context, payload *dto.UpdateCollectionRequest) (*dto.UpdateCollectionResponse, error) {
	var result *dto.UpdateCollectionResponse
	var data *model.CollectionModel

	find, err := s.repo.FindByID(ctx, &payload.ID)
	if find == nil {
		return nil, response.ErrorBuilder(&response.ErrorConstant.NotFound, err)
	}

	data = &model.CollectionModel{
		BaseModel:  model.BaseModel{ID: payload.ID},
		Collection: payload.Collection,
	}

	_, err = s.repo.Update(ctx, &payload.ID, data)

	if err != nil {
		return nil, response.ErrorBuilder(&response.ErrorConstant.UnprocessableEntity, err)
	}

	result = &dto.UpdateCollectionResponse{
		CollectionModel: *data,
	}

	return result, nil
}

func (s *Service) Delete(ctx context.Context, payload *dto.DeleteCollectionRequest) (*dto.DeleteCollectionResponse, error) {
	var result *dto.DeleteCollectionResponse

	data, err := s.repo.FindByID(ctx, &payload.ID)
	if data == nil {
		return nil, response.ErrorBuilder(&response.ErrorConstant.NotFound, err)
	}

	_, err = s.repo.Delete(ctx, &payload.ID, data)

	if err != nil {
		return nil, response.ErrorBuilder(&response.ErrorConstant.UnprocessableEntity, err)
	}

	result = &dto.DeleteCollectionResponse{
		ID: &payload.ID,
	}

	return result, nil
}
