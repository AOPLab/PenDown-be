package controller

import (
	"net/http"

	"github.com/AOPLab/PenDown-be/src/service"

	"github.com/gin-gonic/gin"
)

type AddTagInput struct {
	Tag_name string `json:"tag_name" binding:"required"`
}

type TagResult struct {
	Tag_id   int64
	Tag_name string
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
		"tag_name": tag.ID,
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
				Tag_id:   tag.ID, // 改成小寫開頭會錯
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
