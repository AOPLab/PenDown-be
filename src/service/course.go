package service

import (
	"github.com/AOPLab/PenDown-be/src/model"
	"github.com/AOPLab/PenDown-be/src/persistence"
)

// func FindSchools() ([]*model.School, error) {

// 	var schools []*model.School

// 	if res := persistence.DB.Find(&schools); res.Error != nil {
// 		return nil, res.Error
// 	}
// 	return schools, nil
// }

func FindSchoolCourse(school_id int64) ([]*model.Course, error) {

	var schoolCourse []*model.Course
	if res := persistence.DB.Where("School_id = ?", school_id).Find(&schoolCourse); res.Error != nil {
		return nil, res.Error
	}
	return schoolCourse, nil
}
