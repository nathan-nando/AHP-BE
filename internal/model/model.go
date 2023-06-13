package model

import "time"

type BaseModel struct {
	ID string `json:"id" gorm:"primaryKey"`

	CreatedAt  time.Time  `json:"created_at"`
	CreatedBy  string     `json:"created_by"`
	ModifiedAt *time.Time `json:"modified_at"`
	ModifiedBy *string    `json:"modified_by"`
}
