package model

import (
	"time"

	"gorm.io/gorm"
)

// Note: filename should be "{create_time}_{random_string}"
type Note struct {
	gorm.Model
	ID                  int64     `gorm:"primary_key;auto_increment" json:"note_id"`
	User_id             int64     `gorm:"not null" json:"user_id"`
	Title               string    `gorm:"not null" json:"title"`
	Description         string    `json:"description"`
	View_cnt            int64     `gorm:"default:0" json:"view_cnt"`
	Is_template         bool      `gorm:"default:false" json:"is_template"`
	Course_id           int64     `gorm:"default:null" json:"course_id"`
	Bean                int       `gorm:"default:0" json:"bean"`
	Pdf_filename        string    `gorm:"default:null" json:"pdf_filename"`
	Preview_filename    string    `gorm:"default:null" json:"preview_filename"`
	Goodnotes_filename  string    `gorm:"default:null" json:"goodnotes_filename"`
	Notability_filename string    `gorm:"default:null" json:"notability_filename"`
	CreatedAt           time.Time `json:"created_at"`
	User                User      `gorm:"foreignKey:User_id;constraint:OnDelete:SET NULL;"`
	Course              Course    `gorm:"foreignKey:Course_id;constraint:OnDelete:SET NULL;"`
}

type Saved struct {
	gorm.Model
	User_id int64 `gorm:"not null;uniqueIndex:compositeindex_like;" json:"user_id"`
	Note_id int64 `gorm:"not null;uniqueIndex:compositeindex_like;" json:"note_id"`
	User    User  `gorm:"foreignKey:User_id;constraint:OnDelete:CASCADE;"`
	Note    Note  `gorm:"foreignKey:Note_id;constraint:OnDelete:CASCADE;"`
}

type Download struct {
	gorm.Model
	User_id int64 `gorm:"not null;uniqueIndex:compositeindex_download;" json:"user_id"`
	Note_id int64 `gorm:"not null;uniqueIndex:compositeindex_download;" json:"note_id"`
	User    User  `gorm:"foreignKey:User_id;constraint:OnDelete:CASCADE;"`
	Note    Note  `gorm:"foreignKey:Note_id;constraint:OnDelete:CASCADE;"`
}
