package models

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
)

type CustomersStatus struct {
	Code        string    `json:"code"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
}

func (cs *CustomersStatus) TableName() string {
	return "customers_status"
}

func (cs *CustomersStatus) CustomersStatusFindByCode(code string, db *gorm.DB) (*CustomersStatus, error) {
	err := db.Debug().Where(&CustomersStatus{Code: code}).Find(&cs).Error
	if err != nil {
		return &CustomersStatus{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &CustomersStatus{}, errors.New("ไม่พบสถานะนี้")
	}
	return cs, nil
}

func (cs *CustomersStatus) CustomersStatusFindAll(db *gorm.DB) (*[]CustomersStatus, error) {
	css := []CustomersStatus{}
	err := db.Debug().Limit(20).Find(&css).Error
	if err != nil {
		return &[]CustomersStatus{}, err
	}
	return &css, nil
}
