package config

import (
	"os"
	"strconv"
)

type DB struct {
	Host     string
	Port     int
	Name     string
	User     string
	Password string
}

var db = &DB{}

// DBCfg returns the default DB configuration
func DBCfg() *DB {
	return db
}

func LoadDbCfg() {
	db.Host = os.Getenv("DB_HOST")
	db.Port, _ = strconv.Atoi(os.Getenv("DB_PORT"))
	db.User = os.Getenv("DB_USER")
	db.Password = os.Getenv("DB_PASSWORD")
	db.Name = os.Getenv("DB_NAME")
}
