package model

import (
	"ahp-be/pkg/date"
	"gorm.io/gorm"
)

type Score struct {
	TimbulanSampah        float64 `json:"timbulan_sampah"`
	JarakTpa              float64 `json:"jarak_tpa"`
	JarakPemukiman        float64 `json:"jarak_pemukiman"`
	JarakSungai           float64 `json:"jarak_sungai"`
	PartisipasiMasyarakat float64 `json:"partisipasi_masyarakat"`
	CakupanRumah          float64 `json:"cakupan_rumah"`
	Aksesibilitas         float64 `json:"aksesibilitas"`
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
