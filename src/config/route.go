package config

import (
	"PenDown/src/controller"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	url := r.Group("")
	{
		url.POST("/api/v1/urls", controller.UploadUrl)
		url.GET("/:url_id", controller.RedirectUrl)
	}
}
