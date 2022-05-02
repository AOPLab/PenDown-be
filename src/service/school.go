package service

import (
	"github.com/AOPLab/PenDown-be/src/model"
	"github.com/AOPLab/PenDown-be/src/persistence"
)

func FindSchools() ([]*model.School, error) {

	var schools []*model.School

	if res := persistence.DB.Find(&schools); res.Error != nil {
		return nil, res.Error
	}
	return schools, nil
}

func FindSchool(school_id int64) (*model.School, error) {

	var school model.School
	if res := persistence.DB.Where("ID = ?", school_id).First(&school); res.Error != nil {
		return nil, res.Error
	}
	return &school, nil
}

type SearchSchoolOutput struct {
	ID          int64  `json:"school_id"`
	School_name string `json:"school_name"`
}

func SearchSchool(q string, offset int, limit int) ([]SearchSchoolOutput, int64, error) {
	var results []SearchSchoolOutput
	var count int64
	searchName := "%" + q + "%"
	if err := persistence.DB.Limit(limit).Offset(offset).Table("schools").Select("ID, school_name").Where("school_name LIKE ?", searchName).Find(&results).Error; err != nil {
		return results, 0, err
	}
	persistence.DB.Table("schools").Where("school_name LIKE ?", searchName).Count(&count)
	return results, count, nil
}
