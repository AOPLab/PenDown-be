package model

import (
	"gorm.io/gorm"
)

type Follow struct {
	gorm.Model
	Follower_id int64 `gorm:"not null;uniqueIndex:compositeindex;" json:"follower_id"`
	Followee_id int64 `gorm:"not null;uniqueIndex:compositeindex;" json:"followee_id"`
	Follower    User  `gorm:"foreignKey:Follower_id;constraint:OnDelete:CASCADE;"`
	Followee    User  `gorm:"foreignKey:Followee_id;constraint:OnDelete:CASCADE;"`
}
