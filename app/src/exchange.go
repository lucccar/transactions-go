package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type ExchangeData struct {
	Record_date             string `json:record_date`
	Country                 string `json:country`
	Currency                string `json:currency`
	Country_currency_desc   string `json:country_currency_desc`
	Exchange_rate           string `json:exchange_rate`
	Effective_date          string `json:effective_date`
	Src_line_nbr            string `json:src_line_nbr`
	Record_fiscal_year      string `json:record_fiscal_year`
	Record_fiscal_quarter   string `json:record_fiscal_quarter`
	Record_calendar_year    string `json:record_calendar_year`
	Record_calendar_quarter string `json:record_calendar_quarter`
	Record_calendar_month   string `json:record_calendar_month`
	Record_calendar_day     string `json:record_calendar_day`
}
type TreasuryData struct {
	Data []ExchangeData `json:data`
}

func getExchangeRate(date *time.Time, targetCurrency *string) (float64, error) {
	var treasuryData TreasuryData

	apiURL := urlBuilder(date, targetCurrency)

	response, err := http.Get(apiURL)
	if err != nil {
		log.Fatal(err)
		return 0, err

	}
	defer response.Body.Close()

	jsonBody, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
		return 0, err

	}

	err = json.Unmarshal(jsonBody, &treasuryData)
	if err != nil {
		fmt.Printf("Error parsing JSON: %v\n", err)
		return 0, err
	}
	if len(treasuryData.Data) == 0 {
		return 0, errors.New("the purchase cannot be converted to the target currency")
	}
	var exchangeRate float64
	for _, currency := range treasuryData.Data {
		if currency.Currency == *targetCurrency {
			exchangeRate, err = strconv.ParseFloat(currency.Exchange_rate, 64)
			if err != nil {
				log.Fatal(err)
				return 0, err
			}
			break
		}
	}
	return exchangeRate, nil
}

func urlBuilder(date *time.Time, targetCurrency *string) string {

	currencyFilter := fmt.Sprintf("%s%s", "filter=currency:eq:", *targetCurrency)
	ltFilter := fmt.Sprintf("%s%s", ",record_date:lt:", date.Format("2006-01-02"))

	gtFilter := fmt.Sprintf("%s%s", ",record_date:gt:", date.AddDate(0, -6, 0).Format("2006-01-02"))

	sortParam := "&sort=-record_date"

	baseURL := os.Getenv("EXCHANGE_BASE_URL")

	return fmt.Sprintf("%s%s%s%s%s", baseURL, currencyFilter, gtFilter, ltFilter, sortParam)
}
