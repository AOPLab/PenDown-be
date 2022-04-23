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
		"tag_name": tag.Tag_name,
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
			interfaceSlice[i] = tag
		}
		c.JSON(200, interfaceSlice)
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
			interfaceSlice[i] = school
		}
		c.JSON(200, interfaceSlice)
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
			interfaceSlice[i] = course
		}
		c.JSON(200, interfaceSlice)
	}
	return
}
