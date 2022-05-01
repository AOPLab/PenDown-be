package controller

import (
	"net/http"
	"strconv"

	"github.com/AOPLab/PenDown-be/src/service"
	"github.com/gin-gonic/gin"
)

// q any string
// note_type: all, notability, goodnotes
// filter: people, tags, schools, courses, notes, templates
func Search(c *gin.Context) {
	q := c.Query("q")
	note_type := c.Query("type")
	filter := c.Query("filter")
	offset_o := c.Query("offset")

	offset, err := strconv.Atoi(offset_o)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "OffsetParseError",
		})
		return
	}

	switch filter {
	case "people":
		result, total_cnt, err := service.SearchUser(q, offset, 12)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"people":    result,
			"total_cnt": total_cnt,
		})
		return
	case "tags":
		result, total_cnt, err := service.SearchTag(q, offset, 12)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"tags":      result,
			"total_cnt": total_cnt,
		})
		return
	case "schools":
		result, total_cnt, err := service.SearchSchool(q, offset, 12)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"schools":   result,
			"total_cnt": total_cnt,
		})
		return
	case "courses":
		result, total_cnt, err := service.SearchCourse(q, offset, 12)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"courses":   result,
			"total_cnt": total_cnt,
		})
		return
	case "notes":
		c.JSON(http.StatusOK, "Test")
		return
	case "templates":
		c.JSON(http.StatusOK, gin.H{
			"q":         q,
			"note_type": note_type,
			"filter":    filter,
			"offset":    offset,
		})
		return
	default:
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "NoFilter",
		})
		return
	}

}
