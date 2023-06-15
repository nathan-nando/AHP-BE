package dto

import (
	"ahp-be/internal/model"
	"ahp-be/pkg/response"
)

type FindCollectionsResponseDoc struct {
	Body struct {
		Meta response.Meta           `json:"meta"`
		Data []model.CollectionModel `json:"data"`
	} `json:"body"`
}

type FindCollectionByIDResponse struct {
	model.CollectionModel
}

type FindCollectionByIDResponseDoc struct {
	Body struct {
		Meta response.Meta              `json:"meta"`
		Data FindCollectionByIDResponse `json:"data"`
	} `json:"body"`
}

type CreateCollectionResponse struct {
	model.CollectionModel
}

type CreateCollectionResponseDoc struct {
	Body struct {
		Meta response.Meta            `json:"meta"`
		Data CreateCollectionResponse `json:"data"`
	} `json:"body"`
}

type UpdateCollectionResponse struct {
	model.CollectionModel
}
type UpdateCollectionResponseDoc struct {
	Body struct {
		Meta response.Meta            `json:"meta"`
		Data UpdateCollectionResponse `json:"data"`
	} `json:"body"`
}

type DeleteCollectionResponse struct {
	ID *string `json:"id"`
}

type DeleteCollectionResponseDoc struct {
	Body struct {
		Meta response.Meta            `json:"meta"`
		Data DeleteCollectionResponse `json:"data"`
	} `json:"body"`
}
