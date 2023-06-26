package dto

import "ahp-be/internal/model"

type FindAlternativeByIDRequest struct {
	ID string `json:"id" param:"id" validate:"required"`
}

type CreateAlternativeRequest struct {
	model.Alternative
	CollectionID string `json:"id"`
}

type UpdateAlternativeRequest struct {
	ID string `json:"id" param:"id" validate:"required"`
	model.Alternative
}

type DeleteAlternativeRequest struct {
	ID string `json:"id" param:"id" validate:"required"`
	model.Alternative
}
