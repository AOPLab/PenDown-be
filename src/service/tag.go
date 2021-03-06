package service

import (
	"strings"

	"github.com/AOPLab/PenDown-be/src/model"
	"github.com/AOPLab/PenDown-be/src/persistence"
)

func AddTag(tag_name string) (*model.Tag, error) {

	tag_name = strings.ToLower(tag_name)
	tag := &model.Tag{
		Tag_name: tag_name,
	}

	db_err := persistence.DB.Model(&model.Tag{}).Where("Tag_name = ?", tag_name).FirstOrCreate(&tag).Error
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

func GetTagsByBatch(tagIds *[]int64) (*[]model.Tag, error) {
	var tags []model.Tag
	db_err := persistence.DB.Select([]string{"id", "tag_name"}).Where("id IN ?", *tagIds).Find(&tags).Error
	if db_err != nil {
		return nil, db_err
	}
	return &tags, nil
}

func FindTag(tag_id int64) (*model.Tag, error) {

	var tag model.Tag
	if res := persistence.DB.Where("ID = ?", tag_id).First(&tag); res.Error != nil {
		return nil, res.Error
	}
	return &tag, nil
}

type SearchTagOutput struct {
	ID       int64  `json:"tag_id"`
	Tag_name string `json:"tag_name"`
}

func SearchTag(q string, offset int, limit int) ([]SearchTagOutput, int64, error) {
	var results []SearchTagOutput
	var count int64
	searchName := "%" + strings.ToLower(q) + "%"
	if err := persistence.DB.Limit(limit).Offset(offset).Table("tags").Select("ID, tag_name").Where("lower(tag_name) LIKE ?", searchName).Find(&results).Error; err != nil {
		return results, 0, err
	}
	persistence.DB.Table("tags").Where("lower(tag_name) LIKE ?", searchName).Count(&count)
	return results, count, nil
}
