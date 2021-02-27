package models

import (
	"gorm.io/gorm"
)

type Exercise struct {
	gorm.Model

	ID            uint16 `gorm:"primaryKey"`
	Name          string `gorm:"type:VARCHAR(50);not null"`
	SetsQuantity  uint8  `gorm:"not null"`
	RepsInEachSet uint8  `gorm:"not null"`

	Description string `gorm:"type:TEXT"`
	MuscleGroup string `gorm:"type:VARCHAR(50)"`
}
