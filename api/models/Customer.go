package models

import (
	"errors"
	"html"
	"itjournal/configs"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	"github.com/jinzhu/gorm"
)

type Customer struct {
	ID           int64     `json:"id"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	Cid          int64     `json:"cid"`
	Fname        string    `json:"fname"`
	Lname        string    `json:"lname"`
	Phone        string    `json:"phone"`
	RefreshToken string    `json:"refresh_token"`
	Status       int       `json:"status"`
	Job          string    `json:"job"`
	Sector       string    `json:"sector"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (c *Customer) Prepare() {
	c.ID = 0
	c.Email = html.EscapeString(strings.TrimSpace(c.Email))
	c.Fname = html.EscapeString(strings.TrimSpace(c.Fname))
	c.Lname = html.EscapeString(strings.TrimSpace(c.Lname))
	c.Phone = html.EscapeString(strings.TrimSpace(c.Phone))
	c.Job = html.EscapeString(strings.TrimSpace(c.Job))
	c.Sector = html.EscapeString(strings.TrimSpace(c.Sector))
	c.CreatedAt = time.Now()
	c.UpdatedAt = time.Now()
}

func (c *Customer) Validate(action string) error {

	switch action {
	case "update":
		if c.Fname == "" {
			return errors.New(configs.MessageErrReq + "ชื่อ")
		}
		if c.Lname == "" {
			return errors.New(configs.MessageErrReq + "นามสกุล")
		}
		if c.Phone == "" {
			return errors.New(configs.MessageErrReq + "เบอร์โทรศัพท์")
		}
		if c.Job == "" {
			return errors.New(configs.MessageErrReq + "ตำแหน่ง")
		}
		if c.Sector == "" {
			return errors.New(configs.MessageErrReq + "หน่วยงาน")
		}
		return nil
	default:
		if c.Email == "" {
			return errors.New(configs.MessageErrReq + "อีเมล")
		}
		if err := checkmail.ValidateFormat(c.Email); err != nil {
			return errors.New("รูปแบบอีเมลไม่ถูกต้อง")
		}
		if c.Fname == "" {
			return errors.New(configs.MessageErrReq + "ชื่อ")
		}
		if c.Lname == "" {
			return errors.New(configs.MessageErrReq + "นามสกุล")
		}
		if c.Phone == "" {
			return errors.New(configs.MessageErrReq + "เบอร์โทรศัพท์")
		}
		if c.Job == "" {
			return errors.New(configs.MessageErrReq + "ตำแหน่ง")
		}
		if c.Sector == "" {
			return errors.New(configs.MessageErrReq + "หน่วยงาน")
		}
		return nil
	}
}

func (c *Customer) CustomerSave(db *gorm.DB) (*Customer, error) {
	err := db.Debug().Create(&c).Error
	if err != nil {
		return &Customer{}, err
	}
	return c, nil
}
