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
		authorized.POST("/notes", controller.AddNote)
		authorized.POST("/notes/:note_id/tags/:tag_id", controller.AddNoteTag)
		authorized.DELETE("/notes/:note_id/tags/:tag_id", controller.DeleteNoteTag)
		authorized.POST("/notes/:note_id/notability", controller.UploadNotability)
		authorized.POST("/notes/:note_id/goodnote", controller.UploadGoodnote)
	}
}
