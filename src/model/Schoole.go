package model

import (
	"gorm.io/gorm"
)

type School struct {
	gorm.Model
	ID          int64  `gorm:"primary_key;auto_increment" json:"school_id"`
	School_name string `gorm:"not null" json:"school_name"`
}
