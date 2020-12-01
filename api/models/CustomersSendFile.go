package models

import (
	"errors"
	"html"
	"itjournal/configs"
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

func (csf *CustomersSendFile) Validate(action string) error {
	switch action {
	case "update":
		return nil
	default:
		if csf.Topic == "" {
			return errors.New(configs.MessageErrReq + "หัวข้อ")
		}
		if csf.Job == "" {
			return errors.New(configs.MessageErrReq + "ตำแหน่ง")
		}
		if csf.GovSector == "" {
			return errors.New(configs.MessageErrReq + "หน่วยงานราชการ")
		}
		if csf.Phone == "" {
			return errors.New(configs.MessageErrReq + "เบอร์โทรศัพท์")
		}
		if csf.SendType == "" {
			return errors.New(configs.MessageErrSelectReq + "ประเภทบทความที่ขอส่ง")
		}
		if csf.StatusSurety == "" {
			return errors.New(configs.MessageErrSelectReq + "รับรองบทความ")
		}
		return nil
	}
}

func (af *CustomersSendFile) CustomerSaveSendFile(db *gorm.DB) (*CustomersSendFile, error) {
	var err error
	err = db.Debug().Create(&af).Error
	if err != nil {
		return &CustomersSendFile{}, err
	}
	return af, nil
}
