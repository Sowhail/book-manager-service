package db

import (
	"bookManagement/config"
	models "bookManagement/db/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Db struct {
	db       *gorm.DB
	dbConfig config.Config
}

func CreateDb(dbConfig config.Config) (*Db, error) {
	dsn := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=disable",
		dbConfig.Db.Host,
		dbConfig.Db.Port,
		dbConfig.Db.Name,
		dbConfig.Db.UserName,
		dbConfig.Db.Password,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{FullSaveAssociations: true})
	if err != nil {
		return nil, err
	}

	return &Db{
		db:       db,
		dbConfig: dbConfig,
	}, nil
}

func (db *Db) CreateSchemas() error {
	if err := db.db.AutoMigrate(&models.Author{}); err != nil {
		return err
	}
	if err := db.db.AutoMigrate(&models.Book{}); err != nil {
		return err
	}
	if err := db.db.AutoMigrate(&models.User{}); err != nil {
		return err
	}
	return nil
}
