package dto

import "ahp-be/internal/model"

type AHPByCollectionIDRequest struct {
	CollectionID string `json:"collection_id" param:"collection_id" validate:"required"`
}

type CriteriaUpdateRequest struct {
	Pairwise model.Matrix `json:"pairwise"`
}
