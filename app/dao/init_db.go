package dao

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() (*DataStore, error) {

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	ds := DataStore{db: db}
	if err != nil {
		log.Fatal(err)
		return &ds, err
	}

	ds.db.AutoMigrate(&Purchase{})

	return &ds, nil
}
