package controller

import (
	"net/http"
	"strconv"

	"github.com/AOPLab/PenDown-be/src/service"

	"github.com/gin-gonic/gin"
)

type SchoolResult struct {
	School_id   int64  `json:"school_id"`
	School_name string `json:"school_name"`
}

// GET Schools
func GetSchools(c *gin.Context) {
	schools, err := service.FindSchools()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	} else {
		var interfaceSlice []interface{} = make([]interface{}, len(schools))
		for i, school := range schools {
			schoolInfo := &SchoolResult{
				School_id:   school.ID,
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
		return
	}
	school, err := service.FindSchool(school_id)

	if err != nil {
		if err.Error() == "record not found" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "school_id not exists",
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
			"school_id":   school.ID,
			"school_name": school.School_name,
		})
	}
	return
}
