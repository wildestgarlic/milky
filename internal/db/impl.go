package db

import (
	"TelebotOne/internal/config"
	"database/sql"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/url"
	"sync"
)

var (
	cfg  = config.Config
	once sync.Once
	db   *gorm.DB

	dsn = url.URL{
		Scheme:   cfg.DB.Scheme,
		User:     url.UserPassword(cfg.DB.Username, cfg.DB.Password),
		Host:     fmt.Sprintf("%s:%d", cfg.DB.Host, cfg.DB.Port),
		Path:     cfg.DB.Name,
		RawQuery: (&url.Values{"sslmode": []string{cfg.DB.SSLMode}}).Encode(),
	}
)

func gormInit() *gorm.DB {
	var err error

	once.Do(
		func() {
			db, err = gorm.Open(postgres.Open(dsn.String()), &gorm.Config{})
			if err != nil {
				log.Fatalf("Gorm connect error: %v", err)
			}

			// Connection pool settings
			//sqlDB, err := db.DB()
			//if err != nil {
			//	log.Fatalf("Connection for pool settings error: %v", err)
			//}
			//sqlDB.SetMaxIdleConns(int(cfg.DB.MaxIdleConns)) // must be <= than max open
			//sqlDB.SetMaxOpenConns(int(cfg.DB.MaxOpenConns) ) // in-use + idle life in seconds
			//sqlDB.SetConnMaxLifetime(time.Duration(cfg.DB.MaxConnsLifeTime) * time.Second)
		})

	fmt.Println("Success DB connection")
	return db
}

type impl struct {
	db *gorm.DB
}

func New() Gormer {
	c := &impl{
		db: gormInit(),
	}
	return c
}

func (c *impl) GetConnection() *gorm.DB {
	return c.db
}

func (c *impl) Shutdown() {
	var sqlDB *sql.DB
	var err error

	sqlDB, err = c.db.DB()
	if err != nil {
		log.Fatalf("shutdown connection error: %v", err)
	}

	err = sqlDB.Close()
	if err != nil {
		log.Fatalf("shutdown connection error: %v", err)
	}
}

func (c *impl) Migrate(models ...interface{}) {
	c.migration(models)
}

func (c *impl) migration(models ...interface{}) {
	var err error

	for _, model := range models {
		err = c.db.AutoMigrate(model)
		if err != nil {
			c.db = c.db.Rollback()
			if c.db.Error != nil {
				log.Printf("Rollback at table >%s error: %s", model, err)
			}
			log.Printf("Migrate at table >%s error: %s", model, err)
		}
	}
}
