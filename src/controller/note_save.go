package controller

import (
	"net/http"
	"strconv"

	"github.com/AOPLab/PenDown-be/src/service"

	"github.com/gin-gonic/gin"
)

// Is Note Saved
func IsNoteSaved(c *gin.Context) {
	// var form AddTagInput
	user_id := c.MustGet("user_id").(int64)
	id := c.Params.ByName("note_id")
	note_id, parse_err := strconv.ParseInt(id, 0, 64)

	if parse_err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Input format error",
		})
		return
	}

	// bool saved
	saved, err := service.SavedNote(user_id, note_id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"is_saved": saved,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"is_saved": saved,
	})
	return
}

// Save Note
func SaveNote(c *gin.Context) {
	user_id := c.MustGet("user_id").(int64)
	id := c.Params.ByName("note_id")
	note_id, parse_err := strconv.ParseInt(id, 0, 64)

	if parse_err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Input format error",
		})
		return
	}

	err := service.SaveNote(user_id, note_id)
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

// Cancel Save
func DeleteSave(c *gin.Context) {
	user_id := c.MustGet("user_id").(int64)
	id := c.Params.ByName("note_id")
	note_id, parse_err := strconv.ParseInt(id, 0, 64)

	if parse_err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Input format error",
		})
		return
	}

	err := service.DeleteSave(user_id, note_id)
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
