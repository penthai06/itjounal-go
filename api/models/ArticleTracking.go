package models

import "time"

type ArticleTracking struct {
	ID        uint      `json:"id"`
	Aid       uint      `json:"aid"`
	Afid      uint      `json:"afid"`
	Status    int       `json:"status"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
}
