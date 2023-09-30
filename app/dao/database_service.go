package dao

import (
	"time"

	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DataStore struct {
	db *gorm.DB
}

type Purchase struct {
	gorm.Model
	ID                uint   `gorm:"AUTO_INCREMENT;PRIMARY_KEY;not null"`
	Description       string `gorm:"not null;type:varchar(50)"`
	TransactionDate   time.Time
	PurchaseAmountUSD float64 `gorm:"not null"`
}

func (ds *DataStore) GetDB() *gorm.DB {
	return ds.db
}
func (ds *DataStore) CreateRecord(record *Purchase) (uint, error) {
	result := ds.db.Create(record)
	if result.Error != nil {

		return 0, result.Error
	}
	return record.ID, nil
}

func (ds *DataStore) RetrieveRecordByID(id *uint) (*Purchase, error) {
	var purchase Purchase
	if retrievedRecord := ds.db.Where("ID = ?", *id).First(&purchase); retrievedRecord.Error != nil {
		return nil, retrievedRecord.Error
	}

	return &purchase, nil
}
