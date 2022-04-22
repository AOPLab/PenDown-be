package controller

import (
	"net/http"
	"strconv"

	"github.com/AOPLab/PenDown-be/src/service"

	"github.com/gin-gonic/gin"
)

type AddNoteInput struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Is_template *bool  `json:"is_template" binding:"required"`
	Course_id   int64  `json:"course_id" binding:"required"`
	Bean        int    `json:"bean" binding:"required"`
}

// Create Note and add tags
func AddNote(c *gin.Context) {
	var form AddNoteInput
	user_id := c.MustGet("user_id").(int64)

	bindErr := c.BindJSON(&form)
	if bindErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": bindErr.Error(),
		})
		return
	}

	note, err := service.AddNote(user_id, form.Title, form.Description, *form.Is_template, form.Course_id, form.Bean)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"note_id": note.ID,
	})
	return
}

func AddNoteTag(c *gin.Context) {
	id := c.Params.ByName("note_id")
	note_id, parse_err := strconv.ParseInt(id, 0, 64)
	if parse_err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Note_id not exists",
		})
	}

	id = c.Params.ByName("tag_id")
	tag_id, parse_err := strconv.ParseInt(id, 0, 64)
	if parse_err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Tag_id not exists",
		})
	}

	user_id := c.MustGet("user_id").(int64)

	err := service.AddNoteTag(user_id, note_id, tag_id)
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

func DeleteNoteTag(c *gin.Context) {
	id := c.Params.ByName("note_id")
	note_id, parse_err := strconv.ParseInt(id, 0, 64)
	if parse_err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Note_id not exists",
		})
	}

	id = c.Params.ByName("tag_id")
	tag_id, parse_err := strconv.ParseInt(id, 0, 64)
	if parse_err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Tag_id not exists",
		})
	}

	user_id := c.MustGet("user_id").(int64)

	err := service.DeleteNoteTag(user_id, note_id, tag_id)
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
