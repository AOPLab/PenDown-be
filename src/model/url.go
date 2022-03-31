package model

import (
	"time"

	"gorm.io/gorm"
)

type Url struct {
	gorm.Model
	ID           int64     `gorm:"primary_key;auto_increment" json:"id"`
	Original_url string    `gorm:"not null" json:"original_url"`
	Expired_date time.Time `gorm:"not null" json:"time"`
	Url_id       string    `gorm:"not null" json:"url_id"`
}
