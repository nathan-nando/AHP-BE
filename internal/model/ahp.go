package model

type Matrix [][]float64

type Criteria struct {
	PairwiseFromJson        Matrix    `json:"pairwise"`
	PairwiseAfterCalculated Matrix    `json:"pairwise_after_calculated"`
	Weights                 []float64 `json:"weights"`
}
