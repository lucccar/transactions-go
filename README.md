# Purchase Transaction API

The Purchase Transaction API is a Go (Golang) web application that allows you to store and retrieve purchase transactions with currency conversion capabilities.

## Table of Contents

-   [Purchase Transaction API](#purchase-transaction-api)
    -   [Table of Contents](#table-of-contents)
    -   [Introduction](#introduction)
    -   [Getting Started](#getting-started)
    -   [API Endpoints](#api-endpoints)
        -   [Store a Purchase Transaction](#store-a-purchase-transaction)
        -   [Response:](#response)
        -   [Retrieve a Purchase Transaction in a Specified Currency](#retrieve-a-purchase-transaction-in-a-specified-currency)
        -   [Response:](#response-1)
    -   [Environment Variables](#environment-variables)

## Introduction

This API provides a simple and easy-to-use interface for managing purchase transactions. It includes functionality to store purchase transactions with details such as description, transaction date, and purchase amount in United States dollars (USD). Additionally, it supports currency conversion using exchange rates obtained from an external source.

## Getting Started

1. Clone the repository:

    ```bash
    git clone https://github.com/yourusername/purchase-transaction-api.git
    ```

2. Run makefile:
    ```bash
    make build
    ```

You should see a message exposing the port 8080.

## API Endpoints

### Store a Purchase Transaction

-   URL: /store

-   Method: POST

-   Request Body:

```json
{
	"description": "transaction description",
	"transaction_date": "2023-09-25",
	"purchase_amount": 6000
}
```

### Response:

-   Status Code: 201 Created

-   Response Body:

```json
{
	"ID": 123
}
```

### Retrieve a Purchase Transaction in a Specified Currency

-   URL: /retrieve
-   Method: GET
-   Query Parameters:
    -   ID (integer): Purchase transaction ID
    -   Currency (string): Target currency (e.g., "EUR")

### Response:

-   Status Code: 200 OK
-   Response Body:

```json
{
	"ID": "123",
	"Description": "Purchase description",
	"TransactionDate": "2023-09-21T00:00:00Z",
	"PurchaseAmountUSD": 100.5,
	"ExchangeRate": 0.85,
	"PurchaseAmount": 85.42
}
```

## Environment Variables

-   DB_USER: PostgreSQL database username
-   DB_PASSWORD: PostgreSQL database password
-   DB_NAME: PostgreSQL database name
-   DB_HOST: PostgreSQL database host
-   DB_PORT: PostgreSQL database port
-   DATE_LAYOUT: Date layout for parsing (e.g., "2006-01-02")
