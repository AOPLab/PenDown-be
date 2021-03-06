package model

import (
	"gorm.io/gorm"
)

// Note create enum type in db before running
// CREATE TYPE status_type AS ENUM (
//     'BASIC',
//     'VIP');

type statusType string

const (
	BASIC statusType = "BASIC"
	VIP   statusType = "VIP"
)

type User struct {
	gorm.Model
	ID          int64      `gorm:"primary_key;auto_increment" json:"user_id"`
	Google_ID   string     `gorm:"default null" json:"google_id"`
	Username    string     `gorm:"not null;unique" json:"username"`
	Full_name   string     `gorm:"not null" json:"full_name"`
	Email       string     `gorm:"not null" json:"email"`
	Password    string     `json:"password_hash"`
	Description string     `json:"description"`
	Status      statusType `sql:"status_type" gorm:"default:'BASIC';not null" json:"status"`
	Bean        int64      `gorm:"default:150" json:"bean"`
}
