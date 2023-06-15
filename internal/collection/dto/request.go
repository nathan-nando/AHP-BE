package dto

import "ahp-be/internal/model"

type FindCollectionByIDRequest struct {
	ID string `json:"id" param:"id" validate:"required"`
}

type CreateCollectionRequest struct {
	model.Collection
}

type UpdateCollectionRequest struct {
	ID string `json:"id" param:"id" validate:"required"`
	model.Collection
}

type DeleteCollectionRequest struct {
	ID string `json:"id" param:"id" validate:"required"`
	model.Collection
}
