package controller

import (
	"net/http"
	"strconv"

	"github.com/AOPLab/PenDown-be/src/service"

	"github.com/gin-gonic/gin"
)

type EditAccountInput struct {
	Username    string `json:"username"`
	Full_name   string `json:"full_name"`
	Email       string `json:"email"`
	Description string `json:"description"`
}

type EditPasswordInput struct {
	Old_password string `json:"old_password"`
	New_password string `json:"new_password" binding:"required"`
}

// GET /api/account/{account_id}/profile
// unfinished
func GetPublicProfile(c *gin.Context) {
	id := c.Params.ByName("account_id")
	account_id, _ := strconv.ParseInt(id, 0, 64)
	user, followers_num, following_num, note_num, err := service.FindPublicProfile(account_id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"username":      user.Username,
		"description":   user.Description,
		"status":        user.Status,
		"bean":          user.Point,
		"followers_num": followers_num,
		"following_num": following_num,
		"note_num":      note_num,
	})
}
