package models

import (
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type CustomersTimeline struct {
	ID                int64     `json:"id"`
	CsfId             int64     `json:"csf_id"`
	Status            string    `json:"status"`
	StatusDescription string    `json:"status_description"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

func (ct *CustomersTimeline) Prepare() {
	ct.ID = 0
	ct.StatusDescription = html.EscapeString(strings.TrimSpace(ct.StatusDescription))
	ct.CreatedAt = time.Now()
	ct.UpdatedAt = time.Now()
}

func (ct *CustomersTimeline) CustomerTimelineSave(db *gorm.DB) error {
	err := db.Debug().Create(&ct).Error
	if err != nil {
		return err
	}
	return nil
}
