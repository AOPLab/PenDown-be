package config

import (
	"net/http"

	"github.com/AOPLab/PenDown-be/src/auth"
	"github.com/AOPLab/PenDown-be/src/controller"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	public := r.Group("api")
	{
		public.POST("/account", controller.Register)
		public.POST("/login", controller.Login)
		public.POST("/login/google", controller.GoogleLogin)
		public.GET("/tag", controller.GetTags)
		public.POST("/tag", controller.AddTag)
		public.GET("/school", controller.GetSchools)
		public.GET("/school/:school_id", controller.GetSchool)

	}

	// protected member router
	authorized := r.Group("/api")
	authorized.Use(auth.AuthRequired)
	{
		authorized.GET("/jwt/test", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"user_id": c.MustGet("user_id"),
			})
			return
		})
	}
}
