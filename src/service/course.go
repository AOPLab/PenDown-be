package service

import (
	"strings"
	"time"

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

type SearchCourseOutput struct {
	ID                int64     `json:"course_id"`
	Course_no         string    `json:"course_no"`
	Course_name       string    `json:"course_name"`
	School_name       string    `json:"school_name"`
	School_ID         int64     `json:"school_id"`
	Note_cnt          int64     `json:"note_cnt"`
	Last_updated_time time.Time `json:"last_updated_time"`
}

func SearchCourse(q string, offset int, limit int) ([]*SearchCourseOutput, int64, error) {
	var results []*SearchCourseOutput
	var count int64
	var schoolCourse []*model.Course
	searchName := "%" + strings.ToLower(q) + "%"
	if err := persistence.DB.Limit(limit).Offset(offset).Table("courses").Preload("School").Where("lower(course_name) LIKE ?", searchName).Or("lower(course_no) LIKE ?", searchName).Find(&schoolCourse).Error; err != nil {
		return results, 0, err
	}
	for _, course := range schoolCourse {
		var result SearchCourseOutput
		result.ID = course.ID
		result.Course_no = course.Course_no
		result.Course_name = course.Course_name
		result.School_ID = course.School_id
		result.School_name = course.School.School_name
		result.Note_cnt = course.Note_cnt
		result.Last_updated_time = course.Last_updated_time
		results = append(results, &result)
	}
	persistence.DB.Table("courses").Where("lower(course_name) LIKE ?", searchName).Or("lower(course_no) LIKE ?", searchName).Count(&count)
	return results, count, nil
}
