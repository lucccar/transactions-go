package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"time"

	"app/dao"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var ds *dao.DataStore
var err error

func main() {

	err = godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	ds, err = dao.InitDB()
	if err != nil {
		log.Fatal(err)
		return
	}
	defer func() {
		dbInstance, _ := ds.GetDB().DB()
		_ = dbInstance.Close()
	}()

	router := gin.Default()

	router.POST("/store", storePurchase)
	router.GET("/retrieve", retrievePurchase)

	// Start the server
	port := ":8080"
	fmt.Printf("Server listening on %s\n", port)
	if err = router.Run(port); err != nil {
		log.Fatal(err)
	}
}

func storePurchase(c *gin.Context) {

	var purchaseRequest TransactionRequest
	if err := c.ShouldBindJSON(&purchaseRequest); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	date, err := time.Parse(dateLayout, purchaseRequest.TransactionDate)
	if err != nil {
		log.Fatal(err)
	}

	purchase := dao.Purchase{
		TransactionDate:   date,
		Description:       purchaseRequest.Description,
		PurchaseAmountUSD: purchaseRequest.PurchaseAmount,
	}
	fmt.Printf("%+v", purchase)
	createdID, err2 := ds.CreateRecord(&purchase)

	if err2 != nil {
		c.JSON(400, gin.H{"error": err2.Error()})
		return
	}

	c.JSON(201, createdID)
}

func retrievePurchase(c *gin.Context) {
	var retrieveRequest RetrieveRequest

	if err := c.ShouldBindJSON(&retrieveRequest); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	retrievedRecord, err2 := ds.RetrieveRecordByID(&retrieveRequest.ID)
	if err2 != nil {
		c.JSON(404, gin.H{"error": "Record not found"})
		return
	}

	exchangeRate, err := getExchangeRate(&retrievedRecord.TransactionDate, &retrieveRequest.Currency)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	convertedAmount := retrievedRecord.PurchaseAmountUSD * exchangeRate
	ratio := math.Pow(10, float64(2))
	convertedAmount = math.Floor(convertedAmount*ratio) / ratio

	purchaseID := strconv.FormatUint(uint64(retrievedRecord.ID), 10)
	convertedRecord := PurchaseTransaction{
		ID:                purchaseID,
		Description:       retrievedRecord.Description,
		TransactionDate:   retrievedRecord.TransactionDate,
		PurchaseAmountUSD: retrievedRecord.PurchaseAmountUSD,
		ExchangeRate:      exchangeRate,
		PurchaseAmount:    convertedAmount,
	}

	c.JSON(200, convertedRecord)
}
