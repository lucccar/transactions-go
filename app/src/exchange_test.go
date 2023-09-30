package main

import (
	"testing"
	"time"
)

func Test_getExchangeRate(t *testing.T) {
	type args struct {
		date           time.Time
		targetCurrency string
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getExchangeRate(tt.args.date, tt.args.targetCurrency)
			if (err != nil) != tt.wantErr {
				t.Errorf("getExchangeRate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getExchangeRate() = %v, want %v", got, tt.want)
			}
		})
	}
}
