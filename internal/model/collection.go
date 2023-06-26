package model

import (
	"ahp-be/pkg/date"
	"gorm.io/gorm"
)

type Collection struct {
	Name                   string  `json:"name"`
	Description            string  `json:"description"`
	ScoreIsCalculated      bool    `json:"score_is_calculated"`
	FinalScoreIsCalculated bool    `json:"final_score_is_calculated"`
	Latitude               float64 `json:"latitude"`
	Longitude              float64 `json:"longitude"`
	Center                 float64 `json:"center"`
}

type CollectionModel struct {
	BaseModel
	Collection
	Alternatives []AlternativeModel `json:"alternatives" gorm:"foreignKey:CollectionID;constraint:OnDelete:CASCADE;"`
}

func (CollectionModel) TableName() string {
	return "collections"
}

func (m *CollectionModel) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = *date.DateTodayLocal()
	return
}

func (m *CollectionModel) BeforeUpdate(tx *gorm.DB) (err error) {
	m.ModifiedAt = date.DateTodayLocal()
	return
}
