package service

import (
	"errors"

	"github.com/AOPLab/PenDown-be/src/model"
	"github.com/AOPLab/PenDown-be/src/persistence"
	"gorm.io/gorm"
)

func AddTag(tag_name string) (*model.Tag, error) {

	tag := &model.Tag{
		Tag_name: tag_name,
	}

	var tag_exist model.Tag
	err := persistence.DB.Where("Tag_name = ?", tag_name).First(&tag_exist).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		db_err := persistence.DB.Model(&model.Tag{}).Create(&tag).Error
		if db_err != nil {
			return nil, db_err
		}
		return tag, nil
	} else if err != nil {
		return nil, err
	} else {
		return nil, err
	}

}

func FindTags() ([]*model.Tag, error) {

	var tags []*model.Tag

	if res := persistence.DB.Model(&model.Tag{}).Select("ID, Tag_name").Find(&tags); res.Error != nil {
		return nil, res.Error
	}
	return tags, nil
}
