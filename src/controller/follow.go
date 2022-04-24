package controller

import (
	"net/http"
	"strconv"

	"github.com/AOPLab/PenDown-be/src/service"

	"github.com/gin-gonic/gin"
)

type AddFollowInput struct {
	Followee_id int64 `json:"account_id" binding:"required"`
}

// POST /api/account/{account_id}/follow
func AddFollow(c *gin.Context) {
	follower_id, consistence_err := checkIdConsistence(c)
	if consistence_err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": consistence_err.Error(),
		})
		return
	}

	var form AddFollowInput
	bindErr := c.BindJSON(&form)
	if bindErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": bindErr.Error(),
		})
		return
	}

	err := service.AddFollow(follower_id, form.Followee_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func GetFollowers(c *gin.Context) {
	id := c.Params.ByName("account_id")
	account_id, pasre_err := strconv.ParseInt(id, 0, 64)
	if pasre_err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "account id format error",
		})
		return
	}

	followers, err := service.GetFollowers(account_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"followers": followers,
		})
	}
}

func GetFollowing(c *gin.Context) {
	id := c.Params.ByName("account_id")
	account_id, pasre_err := strconv.ParseInt(id, 0, 64)
	if pasre_err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "account id format error",
		})
		return
	}

	following, err := service.GetFollowing(account_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"followings": following,
		})
	}
}

func GetFollow(c *gin.Context) {
	id := c.Params.ByName("account_id")
	account_id, pasre_err := strconv.ParseInt(id, 0, 64)
	if pasre_err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "account id format error",
		})
		return
	}
	id = c.Params.ByName("following_id")
	following_id, pasre_err_2 := strconv.ParseInt(id, 0, 64)
	if pasre_err_2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "following_id not exists",
		})
		return
	}

	has_follow, err := service.GetFollow(account_id, following_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"following": has_follow,
		})
	}
}

// DELETE /api/account/{account_id}/follow
func DeleteFollow(c *gin.Context) {
	follower_id, consistence_err := checkIdConsistence(c)
	if consistence_err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": consistence_err.Error(),
		})
		return
	}

	var form AddFollowInput
	bindErr := c.BindJSON(&form)
	if bindErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": bindErr.Error(),
		})
		return
	}

	err := service.DeleteFollow(follower_id, form.Followee_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}
