package usecase

import (
	"ahp-be/internal/alternative"
	"ahp-be/internal/alternative/dto"
	"ahp-be/internal/model"
	"ahp-be/pkg/response"
	"context"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type Service struct {
	repo alternative.Repository
	db   *gorm.DB
}

func New(repo alternative.Repository, db *gorm.DB) *Service {
	return &Service{
		repo: repo,
		db:   db,
	}
}

func (s *Service) Create(ctx context.Context, payload *dto.CreateAlternativeRequest) (*dto.CreateAlternativeResponse, error) {
	var result *dto.CreateAlternativeResponse
	var data *model.AlternativeModel

	data = &model.AlternativeModel{
		BaseModel:    model.BaseModel{ID: uuid.NewString()},
		Alternative:  payload.Alternative,
		CollectionID: payload.CollectionID,
	}

	find, err := s.repo.FindsByCollectionID(ctx, &payload.CollectionID)

	if len(find) == 0 {
		return nil, response.ErrorBuilder(&response.ErrorConstant.NotFound, err)
	}
	if err != nil {
		return nil, response.ErrorBuilder(&response.ErrorConstant.NotFound, err)
	}

	_, err = s.repo.Create(ctx, data)

	if err != nil {
		return nil, response.ErrorBuilder(&response.ErrorConstant.UnprocessableEntity, err)
	}

	result = &dto.CreateAlternativeResponse{
		AlternativeModel: *data,
	}

	return result, nil
}

func (s *Service) FindByID(ctx context.Context, payload *dto.FindAlternativeByIDRequest) (*dto.FindAlternativeByIDResponse, error) {
	var result *dto.FindAlternativeByIDResponse

	data, err := s.repo.FindByID(ctx, &payload.ID)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.ErrorBuilder(&response.ErrorConstant.NotFound, err)
		}
		return nil, response.ErrorBuilder(&response.ErrorConstant.UnprocessableEntity, err)
	}

	result = &dto.FindAlternativeByIDResponse{
		AlternativeModel: *data,
	}

	return result, nil
}

func (s *Service) FindByCollectionID(ctx context.Context, payload *dto.FindAlternativeByIDRequest) ([]model.AlternativeModel, error) {
	datas := make([]model.AlternativeModel, 0)

	datas, err := s.repo.FindsByCollectionID(ctx, &payload.ID)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return datas, response.ErrorBuilder(&response.ErrorConstant.NotFound, err)
		}
		return datas, response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
	}

	return datas, nil
}

func (s *Service) Update(ctx context.Context, payload *dto.UpdateAlternativeRequest) (*dto.UpdateAlternativeResponse, error) {
	var result *dto.UpdateAlternativeResponse
	var data *model.AlternativeModel

	find, err := s.repo.FindByID(ctx, &payload.ID)
	if find == nil {
		return nil, response.ErrorBuilder(&response.ErrorConstant.NotFound, err)
	}

	data = &model.AlternativeModel{
		BaseModel:   model.BaseModel{ID: payload.ID},
		Alternative: payload.Alternative,
	}

	_, err = s.repo.Update(ctx, data, &payload.ID)

	if err != nil {
		return nil, response.ErrorBuilder(&response.ErrorConstant.UnprocessableEntity, err)
	}

	result = &dto.UpdateAlternativeResponse{
		AlternativeModel: *data,
	}

	return result, nil
}

func (s *Service) Delete(ctx context.Context, payload *dto.DeleteAlternativeRequest) (*dto.DeleteAlternativeResponse, error) {
	var result *dto.DeleteAlternativeResponse

	data, err := s.repo.FindByID(ctx, &payload.ID)
	if data == nil {
		return nil, response.ErrorBuilder(&response.ErrorConstant.NotFound, err)
	}

	_, err = s.repo.Delete(ctx, data, &payload.ID)

	if err != nil {
		return nil, response.ErrorBuilder(&response.ErrorConstant.UnprocessableEntity, err)
	}

	result = &dto.DeleteAlternativeResponse{
		ID: &payload.ID,
	}

	return result, nil
}
