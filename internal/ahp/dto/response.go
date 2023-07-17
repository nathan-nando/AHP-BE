package dto

import "ahp-be/pkg/response"

type CheckConsistencyRatioResponse struct {
	bool
}

type CheckConsistencyRatioResponseDoc struct {
	Body struct {
		Meta response.Meta                 `json:"meta"`
		Data CheckConsistencyRatioResponse `json:"data"`
	} `json:"body"`
}
