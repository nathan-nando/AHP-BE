package model

import (
	"ahp-be/pkg/date"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID         string     `json:"id" gorm:"primaryKey"`
	CreatedAt  time.Time  `json:"created_at"`
	CreatedBy  string     `json:"created_by"`
	ModifiedAt *time.Time `json:"modified_at"`
	ModifiedBy *string    `json:"modified_by"`
}

func (e *BaseModel) BeforeCreate(db *gorm.DB) (err error) {
	e.CreatedAt = *date.DateTodayLocal()
	e.ID = uuid.NewString()
	return
}

func (e *BaseModel) BeforeUpdate(db *gorm.DB) (err error) {
	e.ModifiedAt = date.DateTodayLocal()
	return
}
