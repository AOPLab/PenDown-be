package controller

import (
	"net/http"
	"strconv"

	"github.com/AOPLab/PenDown-be/src/service"

	"github.com/gin-gonic/gin"
)

type AddTagInput struct {
	Tag_name string `json:"tag_name" binding:"required"`
}

type TagResult struct {
	Tag_id   int64
	Tag_name string
}

type SchoolResult struct {
	School_id   int64
	School_name string
}

type CourseResult struct {
	School_id   int64
	Course_id   int64
	Course_name string
	Course_no   string
}

// Create Tag
func AddTag(c *gin.Context) {
	var form AddTagInput

	bindErr := c.BindJSON(&form)
	if bindErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": bindErr.Error(),
		})
		return
	}

	tag, err := service.AddTag(form.Tag_name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"tag_name": tag.ID,
	})
	return
}

// GET Tags
func GetTags(c *gin.Context) {
	tags, err := service.FindTags()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		var interfaceSlice []interface{} = make([]interface{}, len(tags))
		for i, tag := range tags {
			tagInfo := &TagResult{
				Tag_id:   tag.ID, // 改成小寫開頭會錯
				Tag_name: tag.Tag_name,
			}
			interfaceSlice[i] = tagInfo
		}

		c.JSON(200, gin.H{
			"tags": interfaceSlice,
		})
	}
	return

}

// GET Schools
func GetSchools(c *gin.Context) {
	schools, err := service.FindSchools()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		var interfaceSlice []interface{} = make([]interface{}, len(schools))
		for i, school := range schools {
			schoolInfo := &SchoolResult{
				School_id:   school.ID, // 改成小寫開頭會錯
				School_name: school.School_name,
			}
			interfaceSlice[i] = schoolInfo
		}

		c.JSON(200, gin.H{
			"schools": interfaceSlice,
		})
	}
	return
}

// GET School
func GetSchool(c *gin.Context) {
	id := c.Params.ByName("school_id")
	school_id, pasre_err := strconv.ParseInt(id, 0, 64)
	if pasre_err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "school_id not exists",
		})
	}
	school, err := service.FindSchool(school_id)

	if err != nil {
		if err.Error() == "record not found" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "school_id not exists",
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"school_id":   school.ID,
			"school_name": school.School_name,
		})
	}
	return
}

// GET School's Course
func GetSchoolCourse(c *gin.Context) {
	id := c.Params.ByName("school_id")
	school_id, pasre_err := strconv.ParseInt(id, 0, 64)
	if pasre_err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "school_id not exists",
		})
	}
	schoolCourse, err := service.FindSchoolCourse(school_id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		var interfaceSlice []interface{} = make([]interface{}, len(schoolCourse))
		for i, course := range schoolCourse {
			courseInfo := &CourseResult{
				School_id:   course.School_id, // 改成小寫開頭會錯
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
			"error": "course_id not exists",
		})
	}
	course, err := service.FindCourse(course_id)

	if err != nil {
		if err.Error() == "record not found" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "course_id not exists",
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
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
