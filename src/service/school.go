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

func FindSchool()(school_id int64) (*model.School, error) {

	var school model.School
	if res := persistence.DB.Where("ID = ?", school_id).First(&school); res.Error != nil {
		return nil, res.Error
	}
	return &school, nil
}
