package model

import (
	"ahp-be/pkg/date"
	"gorm.io/gorm"
)

type FinalScore struct {
	FinalScore float64 `json:"final_score"`
	Rank       int8    `json:"rank"`
}

type FinalScoreModel struct {
	BaseModel
	FinalScore
	AlternativeID string `json:"alternative_id" gorm:"size:191"`
	CollectionID  string `json:"collection_id" gorm:"size:191"`
}

func (FinalScoreModel) TableName() string {
	return "final_scores"
}

func (m *FinalScoreModel) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = *date.DateTodayLocal()
	return
}

func (m *FinalScoreModel) BeforeUpdate(tx *gorm.DB) (err error) {
	m.ModifiedAt = date.DateTodayLocal()
	return
}
