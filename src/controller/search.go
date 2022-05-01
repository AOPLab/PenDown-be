package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// q any string
// note_type: all, notability, goodnotes
// filter: people, tags, schools, courses, notes, templates
func Search(c *gin.Context) {
	q := c.Query("q")
	note_type := c.Query("type")
	filter := c.Query("filter")
	offset := c.Query("offset")

	switch filter {
	case "people":
		c.JSON(http.StatusOK, "Test")
		return
	case "tags":
		c.JSON(http.StatusOK, "Test")
		return
	case "schools":
		c.JSON(http.StatusOK, "Test")
		return
	case "courses":
		c.JSON(http.StatusOK, "Test")
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
