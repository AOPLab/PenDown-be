package controller

import (
	"net/http"
	"strconv"

	"github.com/AOPLab/PenDown-be/src/service"

	"github.com/gin-gonic/gin"
)

// Edit Note
func EditNote(c *gin.Context) {
	user_id := c.MustGet("user_id").(int64)
	id := c.Params.ByName("note_id")
	note_id, parse_err := strconv.ParseInt(id, 0, 64)

	if parse_err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Input format error",
		})
		return
	}

	var form service.EditNoteInput
	bindErr := c.BindJSON(&form)

	if bindErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": bindErr.Error(),
		})
		return
	}

	err := service.EditNote(user_id, note_id, form)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
		})
	}

	return
}

// Delete Note
func DeleteNote(c *gin.Context) {
	user_id := c.MustGet("user_id").(int64)
	id := c.Params.ByName("note_id")
	note_id, parse_err := strconv.ParseInt(id, 0, 64)

	if parse_err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Input format error",
		})
		return
	}

	err := service.DeleteNote(user_id, note_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
	return
}
