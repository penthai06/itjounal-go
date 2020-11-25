package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Article struct {
	ID        uint      `json:"id"`
	Cid       uint      `json:"cid"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (a *Article) Prepare() {
	a.ID = 0
	a.Title = html.EscapeString(strings.TrimSpace(a.Title))
	a.CreatedAt = time.Now()
	a.UpdatedAt = time.Now()
}

func (a *Article) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if a.Title == "" {
			return errors.New("กรุณาใส่ไฟล์ หัวข้อบทความ")
		}
		return nil
	default:
		if a.Title == "" {
			return errors.New("กรุณาใส่ไฟล์ หัวข้อบทความ")
		}
		return nil
	}
}

func (a *Article) SaveArticle(db *gorm.DB) (*Article, error) {
	var err error
	err = db.Debug().Create(*a).Error
	if err != nil {
		return &Article{}, err
	}
	return a, nil
}
