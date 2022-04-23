package service

import (
	"github.com/AOPLab/PenDown-be/src/model"
	"github.com/AOPLab/PenDown-be/src/persistence"
)

func FindSchoolCourse(school_id int64) ([]*model.Course, error) {

	var schoolCourse []*model.Course
	if res := persistence.DB.Where("School_id = ?", school_id).Find(&schoolCourse); res.Error != nil {
		return nil, res.Error
	}
	return schoolCourse, nil
}

func FindCourse(course_id int64) (*model.Course, error) {

	var course model.Course
	if res := persistence.DB.Where("ID = ?", course_id).First(&course); res.Error != nil {
		return nil, res.Error
	}
	return &course, nil
}