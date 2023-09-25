package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	i "app/initialize"

	"github.com/gin-gonic/gin"

	// Use the appropriate driver for your database
	"gorm.io/gorm"

	_ "github.com/lib/pq"
)

var db *gorm.DB

type PurchaseTransaction struct {
	ID                string    `json:"id"`
	Description       string    `json:"description"`
	TransactionDate   time.Time `json:"transaction_date"`
	PurchaseAmountUSD float64   `json:"purchase_amount_usd"`
	ExchangeRate      float64   `json:"exchange_rate"`
	PurchaseAmount    float64   `json:"purchase_amount"`
}

func main() {
	// Initialize the database connection
	i.InitDB()
	defer func() {
		dbInstance, _ := db.DB()
		_ = dbInstance.Close()
	}()

	router := gin.Default()

	router.POST("/store", storePurchase)
	router.GET("/retrieve", retrievePurchase)

	// Start the server
	port := ":8080"
	fmt.Printf("Server listening on %s\n", port)
	if err := router.Run(port); err != nil {
		log.Fatal(err)
	}
}

func storePurchase(c *gin.Context) {
	// Parse request body into your model struct (e.g., YourDataModel)
	var purchaseRequest PurchaseTransaction
	if err := c.ShouldBindJSON(&purchaseRequest); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Create a new record in the database
	if result := db.Create(&purchaseRequest); result.Error != nil {
		c.JSON(400, gin.H{"error": result.Error.Error()})
		return
	}

	// Respond with the created record
	c.JSON(201, purchaseRequest)
}

func retrievePurchase(c *gin.Context) {
	// Retrieve a record by ID
	id := c.Param("id")
	targetCurrency := c.Param("currency")

	// retrievedRecord, err := db.RetrieveRecordByID(id)
	// Fetch the purchase from the database by ID
	var purchase i.Purchase
	if retrievedRecord := db.Where("id = ?", id).First(&purchase); retrievedRecord.Error != nil {
		c.JSON(404, gin.H{"error": "Record not found"})
		return
	}

	// Fetch the exchange rate for the purchase date from the Treasury API
	exchangeRate, err := getExchangeRate(purchase.TransactionDate, targetCurrency)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Calculate the converted purchase amount
	convertedAmount := purchase.PurchaseAmountUSD * exchangeRate
	purchaseID := strconv.FormatUint(uint64(purchase.ID), 10)
	// Create a response object
	convertedRecord := PurchaseTransaction{
		ID:                purchaseID,
		Description:       purchase.Description,
		TransactionDate:   purchase.TransactionDate,
		PurchaseAmountUSD: purchase.PurchaseAmountUSD,
		ExchangeRate:      exchangeRate,
		PurchaseAmount:    convertedAmount,
	}

	// Respond with the retrieved record
	c.JSON(200, convertedRecord)
}

func getExchangeRate(date time.Time, targetCurrency string) (float64, error) {
	// Implement fetching exchange rate from Treasury API based on date and targetCurrency
	// You should make an HTTP request to the Treasury API and parse the response JSON here.
	// Ensure the rate is within the last 6 months or return an error if not found.
	// Return the exchange rate as a float64.
	// Handle any potential errors during the process.
	// Here, we assume the exchange rate is fetched successfully.
	return 1.0, nil
}
