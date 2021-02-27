package models

import (
	"gorm.io/gorm"
)

type Exercise struct {
	gorm.Model

	ID            uint16 `gorm:"primaryKey"`
	Name          string `gorm:"type:VARCHAR(50);type:NOT NULL"`
	SetsQuantity  uint8  `gorm:"type:NOT NULL"`
	RepsInEachSet uint8  `gorm:"type:NOT NULL"`

	Description string `gorm:"type:TEXT"`
	MuscleGroup string `gorm:"type:VARCHAR(50)"`
}
