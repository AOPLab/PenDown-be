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
	Tag_id   int64  `json:"tag_id"`
	Tag_name string `json:"tag_name"`
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
		"tag_id": tag.ID,
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
		return
	} else {
		var interfaceSlice []interface{} = make([]interface{}, len(tags))
		for i, tag := range tags {
			tagInfo := &TagResult{
				Tag_id:   tag.ID,
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

// GET Tag
func GetTag(c *gin.Context) {
	id := c.Params.ByName("tag_id")
	tag_id, pasre_err := strconv.ParseInt(id, 0, 64)
	if pasre_err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "input format error",
		})
		return
	}
	tag, err := service.FindTag(tag_id)

	if err != nil {
		if err.Error() == "record not found" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "tag_id not exists",
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
			"tag_name": tag.Tag_name,
		})
	}
	return
}
