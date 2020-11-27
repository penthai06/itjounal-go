package models

import (
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type CustomersSendFile struct {
	ID              int64     `json:"id"`
	Cid             int64     `json:"cid"`
	Job             string    `json:"job"`
	GovSector       string    `json:"gov_sector"`
	Phone           string    `json:"phone"`
	SendType        string    `json:"send_type"`
	Topic           string    `json:"topic"`
	StatusSurety    string    `json:"status_surety"`
	StatusCommittee string    `json:"status_committee"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

func (csf *CustomersSendFile) Prepare() {
	csf.ID = 0
	csf.Job = html.EscapeString(strings.TrimSpace(csf.Job))
	csf.GovSector = html.EscapeString(strings.TrimSpace(csf.GovSector))
	csf.Phone = html.EscapeString(strings.TrimSpace(csf.Phone))
	csf.SendType = html.EscapeString(strings.TrimSpace(csf.SendType))
	csf.Topic = html.EscapeString(strings.TrimSpace(csf.Topic))
	csf.StatusCommittee = html.EscapeString(strings.TrimSpace(csf.StatusCommittee))
	csf.StatusSurety = html.EscapeString(strings.TrimSpace(csf.StatusSurety))
	csf.CreatedAt = time.Now()
	csf.UpdatedAt = time.Now()
}

func (csf *CustomersSendFile) Validate() {
}

func (af *CustomersSendFile) CustomerSaveSendFile(db *gorm.DB) (*CustomersSendFile, error) {
	var err error
	err = db.Debug().Create(&af).Error
	if err != nil {
		return &CustomersSendFile{}, err
	}
	return af, nil
}
