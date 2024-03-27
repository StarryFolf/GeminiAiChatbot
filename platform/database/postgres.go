package database

import (
	"fiber/app/models/entities"
	"fiber/pkg/config"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	*gorm.DB
}

var defaultDb = &DB{}

func (db *DB) Connect(cfg *config.DB) (err error) {
	dbUri := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.Name,
	)

	db.DB, err = gorm.Open(postgres.Open(dbUri))

	if err != nil {
		return err
	}

	err = db.AutoMigrate(
		&entities.Conversation{},
		&entities.Message{})
	if err != nil {
		return err
	}

	return nil
}

func ConnectDB() error {
	return defaultDb.Connect(config.DBCfg())
}

func GetDB() *DB {
	return defaultDb
}

func Migrate() error {
	db := GetDB()

	err := db.AutoMigrate(
		&entities.Conversation{},
		&entities.Message{})
	if err != nil {
		return err
	}

	return nil
}

func DropTables() error {
	db := GetDB()
	err := db.Migrator().DropTable(
		&entities.Conversation{},
		&entities.Message{})
	if err != nil {
		return err
	}

	return nil
}
