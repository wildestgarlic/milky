package db

import "gorm.io/gorm"

type Gormer interface {
	Acquire() *gorm.DB
	Release()
	Migrate(models ...interface{})

	Shutdown()
}

