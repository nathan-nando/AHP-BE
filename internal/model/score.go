package model

import (
	"ahp-be/pkg/date"
	"gorm.io/gorm"
)

type Score struct {
}

type ScoreModel struct {
	BaseModel
	Score
	AlternativeID string `json:"alternative_id" gorm:"size:191"`
	CollectionID  string `json:"collection_id" gorm:"size:191"`
}

func (ScoreModel) TableName() string {
	return "scores"
}

func (m *ScoreModel) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = *date.DateTodayLocal()
	return
}

func (m *ScoreModel) BeforeUpdate(tx *gorm.DB) (err error) {
	m.ModifiedAt = date.DateTodayLocal()
	return
}
