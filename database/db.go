package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
	err error
)

func Init() {
	dsn := "host=localhost user=root password=root dbname=personalities port=5432 sslmode=disable TimeZone=Europe/Prague"
	DB, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		panic(err)
	}
}
