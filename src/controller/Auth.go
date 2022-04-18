package controller

import (
	"net/http"

	"github.com/AOPLab/PenDown-be/src/auth"
	"github.com/AOPLab/PenDown-be/src/service"

	"github.com/gin-gonic/gin"
)

type AddAccountInput struct {
	Username  string `json:"username" binding:"required"`
	Full_name string `json:"full_name" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
}

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type GoogleLoginInput struct {
	GoogleToken string `json:"google_token" binding:"required"`
	Name        string `json:"name" binding:"required"`
}

// Register
func Register(c *gin.Context) {
	var form AddAccountInput
	bindErr := c.BindJSON(&form)
	if bindErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": bindErr.Error(),
		})
		return
	}

	user, err := service.AddUser(form.Username, form.Full_name, form.Email, form.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"id": user.ID,
	})
}

// Login
func Login(c *gin.Context) {
	var form LoginInput

	bindErr := c.BindJSON(&form)
	if bindErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": bindErr.Error(),
		})
		return
	}

	// verify user
	user, err := service.VerifyLogin(form.Username, form.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// set jwt
	token, jwt_err := auth.SetClaim(user.ID, false)
	if jwt_err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": jwt_err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"account_id": user.ID,
		"token":      token,
	})
	return
}

func GoogleLogin(c *gin.Context) {
	var form GoogleLoginInput

	bindErr := c.BindJSON(&form)
	if bindErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": bindErr.Error(),
		})
		return
	}

	// verify user
	googleUser, err := service.VerifyIdToken(form.GoogleToken)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	// fmt.Println(user)
	// "audience": "",
	// "email": "gary6658@ntu.im",
	// "expires_in": 3320,
	// "issued_to": "",
	// "user_id": "104599823526264245462",
	// "verified_email": true

	// TODO: Verify Google Login

	// set jwt
	// token, jwt_err := auth.SetClaim(user.ID)
	// if jwt_err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"error": jwt_err.Error(),
	// 	})
	// 	return
	// }
	c.JSON(http.StatusOK, gin.H{
		"account_id": googleUser,
	})
	return
}
