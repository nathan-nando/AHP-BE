package dto

type CheckCRSubCriteria struct {
	Mode string `json:"mode" param:"mode" validate:"required"`
}
