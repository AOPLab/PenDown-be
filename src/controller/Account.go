package controller

import (
	"errors"
	_ "fmt"
	"net/http"
	"strconv"

	"github.com/AOPLab/PenDown-be/src/service"

	"github.com/gin-gonic/gin"
)

func checkIdConsistence(c *gin.Context) (int64, error) {
	token_account_id := c.MustGet("user_id").(int64)

	id := c.Params.ByName("account_id")
	account_id, pasre_err := strconv.ParseInt(id, 0, 64)

	if pasre_err != nil {
		return 0, pasre_err
	}

	if token_account_id != account_id {
		return 0, errors.New("token_account_id and account_id inconsistence")
	} else {
		return account_id, nil
	}
}

// GET /api/account/{account_id}/profile
func GetPublicProfile(c *gin.Context) {
	id := c.Params.ByName("account_id")
	account_id, pasre_err := strconv.ParseInt(id, 0, 64)
	if pasre_err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "account id format error",
		})
		return
	}
	user, followers_num, following_num, note_num, err := service.FindPublicProfile(account_id)

	if err != nil {
		if err.Error() == "record not found" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "account id format error",
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"account_id":    user.ID,
			"username":      user.Username,
			"full_name":     user.Full_name,
			"description":   user.Description,
			"status":        user.Status,
			"bean":          user.Bean,
			"followers_num": followers_num,
			"following_num": following_num,
			"note_num":      note_num,
		})
	}
}

// GET /api/account
func GetPrivateProfile(c *gin.Context) {
	id := c.MustGet("user_id")
	account_id, _ := id.(int64)
	user, followers_num, following_num, note_num, err := service.FindPrivateProfile(account_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"account_id":    user.ID,
			"username":      user.Username,
			"full_name":     user.Full_name,
			"email":         user.Email,
			"description":   user.Description,
			"status":        user.Status,
			"bean":          user.Bean,
			"followers_num": followers_num,
			"following_num": following_num,
			"note_num":      note_num,
			"is_google":     user.Google_ID != "",
			"has_password":  user.Password != "",
		})
	}
}

// PATCH /api/account
func EditProfile(c *gin.Context) {
	id := c.MustGet("user_id")
	account_id, _ := id.(int64)

	var form service.EditAccountInput
	bindErr := c.BindJSON(&form)

	if bindErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": bindErr.Error(),
		})
		return
	}

	err := service.EditProfile(account_id, form)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"account_id": account_id,
		})
	}
}

// PUT /api/account/{account_id}/pass_hash
func EditPassword(c *gin.Context) {
	account_id, consistence_err := checkIdConsistence(c)
	if consistence_err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": consistence_err.Error(),
		})
		return
	}

	var form service.EditPasswordInput
	bindErr := c.BindJSON(&form)

	if bindErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": bindErr.Error(),
		})
		return
	}

	err := service.EditPassword(account_id, form)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"account_id": account_id,
		})
	}

}
