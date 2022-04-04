package model

import (
	"time"

	"gorm.io/gorm"
)

type Note struct {
	gorm.Model
	ID             int64     `gorm:"primary_key;auto_increment" json:"note_id"`
	User_id        int64     `gorm:"not null" json:"user_id"`
	Title          string    `gorm:"not null" json:"title"`
	Description    string    `json:"description"`
	Is_template    bool      `gorm:"default:false" json:"is_template"`
	Course_id      int64     `json:"course_id"`
	Price          int       `gorm:"default:0" json:"price"`
	Pdf_Url        string    `gorm:"default:null" json:"pdf_url"`
	Pdf_Preview    string    `gorm:"default:null" json:"pdf_preview_url"`
	Goodnotes_Url  string    `gorm:"default:null" json:"goodnotes_url"`
	Notability_Url string    `gorm:"default:null" json:"notability_url"`
	CreatedAt      time.Time `json:"created_at"`
	User           User      `gorm:"foreignKey:User_id;constraint:OnDelete:SET NULL;"`
	Course         Course    `gorm:"foreignKey:Course_id;constraint:OnDelete:SET NULL;"`
}

type Liked struct {
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
