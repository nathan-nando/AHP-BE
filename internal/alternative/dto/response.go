package dto

import (
	"ahp-be/internal/model"
	"ahp-be/pkg/response"
)

type FindAlternativesByIDCollectionResponseDoc struct {
	Body struct {
		Meta response.Meta            `json:"meta"`
		Data []model.AlternativeModel `json:"data"`
	} `json:"body"`
}

type FindAlternativeByIDResponse struct {
	model.AlternativeModel
}

type FindAlternativeByIDResponseDoc struct {
	Body struct {
		Meta response.Meta               `json:"meta"`
		Data FindAlternativeByIDResponse `json:"data"`
	} `json:"body"`
}

type CreateAlternativeResponse struct {
	model.AlternativeModel
}

type CreateAlternativeResponseDoc struct {
	Body struct {
		Meta response.Meta             `json:"meta"`
		Data CreateAlternativeResponse `json:"data"`
	} `json:"body"`
}

type UpdateAlternativeResponse struct {
	model.AlternativeModel
}
type UpdateAlternativeResponseDoc struct {
	Body struct {
		Meta response.Meta             `json:"meta"`
		Data UpdateAlternativeResponse `json:"data"`
	} `json:"body"`
}

type DeleteAlternativeResponse struct {
	ID *string `json:"id"`
}

type DeleteAlternativeResponseDoc struct {
	Body struct {
		Meta response.Meta             `json:"meta"`
		Data DeleteAlternativeResponse `json:"data"`
	} `json:"body"`
}
