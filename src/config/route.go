package config

import (
	"github.com/AOPLab/PenDown-be/src/controller"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	public := r.Group("")
	{
		public.POST("/account", controller.Register)
		public.POST("/login", controller.Login)
	}

	url := r.Group("")
	{
		url.POST("/api/v1/urls", controller.UploadUrl)
		url.GET("/:url_id", controller.RedirectUrl)
	}
}
