package dao

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() (*DataStore, error) {
	dbHost := os.Getenv("DBHOST")
	dbPort := os.Getenv("DBPORT")
	dbUser := os.Getenv("DBUSER")
	dbPassword := os.Getenv("DBPASSWORD")
	dbName := os.Getenv("DBNAME")

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", dbHost, dbPort, dbUser, dbPassword, dbName)

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	ds := DataStore{db: db}
	if err != nil {
		log.Fatal(err)
		return &ds, err
	}

	ds.db.AutoMigrate(&Purchase{})

	return &ds, nil
}
