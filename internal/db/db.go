package db

import "gorm.io/gorm"

type Gormer interface {
	GetConnection() *gorm.DB
	Migrate(models ...interface{})

	Shutdown()
}
