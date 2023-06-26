package model

import (
	"ahp-be/pkg/date"
	"gorm.io/gorm"
)

type Alternative struct {
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type AlternativeModel struct {
	BaseModel
	Alternative
	CollectionID string          `json:"collection_id" gorm:"size:191"`
	Score        ScoreModel      `json:"scores" gorm:"foreignKey:AlternativeID;constraint:OnDelete:CASCADE;"`
	FinalScore   FinalScoreModel `json:"final_scores" gorm:"foreignKey:AlternativeID;constraint:OnDelete:CASCADE;"`
}

func (AlternativeModel) TableName() string {
	return "alternatives"
}

func (m *AlternativeModel) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = *date.DateTodayLocal()
	return
}

func (m *AlternativeModel) BeforeUpdate(tx *gorm.DB) (err error) {
	m.ModifiedAt = date.DateTodayLocal()
	return
}
