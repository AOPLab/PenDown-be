package controller

import (
	"net/http"

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
