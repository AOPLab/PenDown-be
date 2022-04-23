package controller

import (
	"net/http"
	"strconv"

	"github.com/AOPLab/PenDown-be/src/service"

	"github.com/gin-gonic/gin"
)

type CourseResult struct {
	School_id   int64  `json:"school_id"`
	Course_id   int64  `json:"course_id"`
	Course_name string `json:"course_name"`
	Course_no   string `json:"course_no"`
}

// GET School's Course
func GetSchoolCourse(c *gin.Context) {
	id := c.Params.ByName("school_id")
	school_id, pasre_err := strconv.ParseInt(id, 0, 64)
	if pasre_err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "input format error",
		})
		return
	}
	schoolCourse, err := service.FindSchoolCourse(school_id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	} else {
		var interfaceSlice []interface{} = make([]interface{}, len(schoolCourse))
		for i, course := range schoolCourse {
			courseInfo := &CourseResult{
				School_id:   course.School_id,
				Course_id:   course.ID,
				Course_name: course.Course_name,
				Course_no:   course.Course_no,
			}
			interfaceSlice[i] = courseInfo
		}

		c.JSON(200, gin.H{
			"courses": interfaceSlice,
		})
	}
	return
}

// GET Course
func GetCourse(c *gin.Context) {
	id := c.Params.ByName("course_id")
	course_id, pasre_err := strconv.ParseInt(id, 0, 64)
	if pasre_err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "input format error",
		})
		return
	}
	course, err := service.FindCourse(course_id)

	if err != nil {
		if err.Error() == "record not found" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "course_id not exists",
			})
			return
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"course_id":   course.ID,
			"course_no":   course.Course_no,
			"course_name": course.Course_name,
			"school_id":   course.School_id,
		})
	}
	return
}
