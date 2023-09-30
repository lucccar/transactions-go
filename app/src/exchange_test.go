package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func Test_getExchangeRate(t *testing.T) {

	_ = godotenv.Load("../.env")

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	mockTime1 := time.Date(2023, 8, 13, 0, 0, 0, 0, time.Local)
	mockCurrrency1 := "Real"
	mockTime2 := time.Date(2015, 8, 13, 0, 0, 0, 0, time.Local)
	mockCurrrency2 := "Peso"
	type args struct {
		date           *time.Time
		targetCurrency *string
	}
	tests := []struct {
		name                string
		args                args
		expected            float64
		mockBodyExternalAPI string
		pathUrl             string
		wantErr             bool
	}{
		{
			name:                "teste1",
			args:                args{date: &mockTime1, targetCurrency: &mockCurrrency1},
			expected:            4.97065,
			mockBodyExternalAPI: `{"data": [{"record_date":"2023-03-31","country":"Brazil","currency":"Real","country_currency_desc": "Brazil-Real","exchange_rate": "4.97065"},{"record_date":"2023-03-31","country":"Sweeden","currency":"Crown","country_currency_desc":"Sweeden-Crow","exchange_rate":"0.87224"}]}`,
			pathUrl:             "https://api.fiscaldata.treasury.gov/services/api/fiscal_service/v1/accounting/od/rates_of_exchange?filter=currency:eq:Real,record_date:gt:2023-02-13,record_date:lt:2023-08-13&sort=-record_date",
			wantErr:             false,
		},
		{
			name:                "teste2",
			args:                args{date: &mockTime2, targetCurrency: &mockCurrrency2},
			expected:            0.87224,
			mockBodyExternalAPI: `{"data": []}`,
			pathUrl:             "https://api.fiscaldata.treasury.gov/services/api/fiscal_service/v1/accounting/od/rates_of_exchange?filter=currency:eq:Peso,record_date:gt:2015-02-13,record_date:lt:2015-08-13&sort=-record_date",
			wantErr:             true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			httpmock.RegisterResponder("GET", tt.pathUrl,
				httpmock.NewStringResponder(200, tt.mockBodyExternalAPI))

			got, err := getExchangeRate(tt.args.date, tt.args.targetCurrency)
			if tt.wantErr == true {
				if err != nil {
					fmt.Println((err.Error()))
				}
				assert.NotNil(t, err)
			} else {
				assert.Equal(t, tt.expected, got, fmt.Sprintln(got))
			}
		})
	}
}
