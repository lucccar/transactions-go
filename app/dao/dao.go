package dao

import (
	"time"

	_ "gorm.io/driver/postgres" // Replace with your database driver import
	"gorm.io/gorm"
)

// DataStore represents the data access layer.
type DataStore struct {
	db *gorm.DB
}

// YourDataModel represents your database table model using GORM tags
type YourDataModel struct {
	ID        uint `gorm:"primary_key"`
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	// Add more fields as needed
}

// CreateRecord creates a new record in the database.
func (ds *DataStore) CreateRecord(record *YourDataModel) error {
	return ds.db.Create(record).Error
}

// RetrieveRecordByID retrieves a record by its ID.
func (ds *DataStore) RetrieveRecordByID(id uint) (*YourDataModel, error) {
	var record YourDataModel
	if err := ds.db.First(&record, id).Error; err != nil {
		return nil, err
	}
	return &record, nil
}

// UpdateRecord updates a record in the database.
func (ds *DataStore) UpdateRecord(record *YourDataModel) error {
	return ds.db.Save(record).Error
}

// DeleteRecord deletes a record from the database.
func (ds *DataStore) DeleteRecord(record *YourDataModel) error {
	return ds.db.Delete(record).Error
}
