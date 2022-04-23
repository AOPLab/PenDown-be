package service

import (
	"github.com/AOPLab/PenDown-be/src/model"
	"github.com/AOPLab/PenDown-be/src/persistence"
)

// type result struct {
// 	ID       int64
// 	Tag_name string
// }

func AddTag(tag_name string) (*model.Tag, error) {

	tag := &model.Tag{
		Tag_name: tag_name,
	}

	db_err := persistence.DB.Model(&model.Tag{}).Create(&tag).Error
	if db_err != nil {
		return nil, db_err
	}

	return tag, nil
}

func FindTags() ([]*model.Tag, error) {

	var tags []*model.Tag

	if res := persistence.DB.Model(&model.Tag{}).Select("ID, Tag_name").Find(&tags); res.Error != nil {
		return nil, res.Error
	}
	return tags, nil
}
