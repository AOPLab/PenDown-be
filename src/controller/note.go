package controller

import (
	"fmt"
	"net/http"

	"github.com/AOPLab/PenDown-be/src/service"

	"github.com/gin-gonic/gin"
)

type AddNoteInput struct {
	Title       string  `json:"title" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Is_template *bool   `json:"is_template" binding:"required"`
	Course_id   int64   `json:"course_id" binding:"required"`
	Bean        int     `json:"bean" binding:"required"`
	Tags        []int64 `json:"tags" binding:"required"`
}

// Create Note and add tags
func AddNote(c *gin.Context) {
	var form AddNoteInput
	user_id := c.MustGet("user_id").(int64)
	fmt.Print(form.Is_template)

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

	for i := 0; i < len(form.Tags); i++ {
		add_tag_err := service.AddNoteTag(note.ID, form.Tags[i])
		if add_tag_err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "AddTagError",
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"note_id": note.ID,
	})
	return
}
