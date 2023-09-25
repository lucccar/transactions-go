package initialize

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	dbHost     = "db"
	dbPort     = 5432
	dbUser     = "myuser"
	dbPassword = "mypassword"
	dbName     = "mydb"
)

type Purchase struct {
	gorm.Model
	Description       string
	TransactionDate   time.Time
	PurchaseAmountUSD float64
}

func InitDB() {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)
	var err error

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	// AutoMigrate will create the "purchase" table based on the Purchase struct
	db.AutoMigrate(&Purchase{})
}
