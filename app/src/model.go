package main

import "time"

type PurchaseTransaction struct {
	ID                string    `json:"id"`
	Description       string    `json:"description"`
	TransactionDate   time.Time `json:"transaction_date"`
	PurchaseAmountUSD float64   `json:"purchase_amount_usd"`
	ExchangeRate      float64   `json:"exchange_rate"`
	PurchaseAmount    float64   `json:"purchase_amount"`
}

type TransactionRequest struct {
	Description     string  `json:"description"`
	TransactionDate string  `json:"transaction_date"`
	PurchaseAmount  float64 `json:"purchase_amount"`
}

type RetrieveRequest struct {
	ID       uint   `json:"id"`
	Currency string `json:"currency"`
}

const dateLayout = "2006-01-02"
