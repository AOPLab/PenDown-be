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
		public.GET("/account/:account_id/profile", controller.GetPublicProfile)
		public.GET("/account/:account_id/followers", controller.GetFollowers)
		public.GET("/account/:account_id/followings", controller.GetFollowing)
		public.GET("/account/:account_id/following/:following_id", controller.GetFollow)
		public.GET("/tag", controller.GetTags)
		public.POST("/tag", controller.AddTag)
		public.GET("/school", controller.GetSchools)
		public.GET("/school/:school_id", controller.GetSchool)
		public.GET("/school/:school_id/course", controller.GetSchoolCourse)
		public.GET("/course/:course_id", controller.GetCourse)

	}

	// protected member router
	authorized := r.Group("/api")
	authorized.Use(auth.AuthRequired)
	{
		authorized.GET("/jwt/test", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"user_id": c.MustGet("user_id"),
			})
		})
		authorized.GET("account", controller.GetPrivateProfile)
		authorized.PATCH("account", controller.EditProfile)
		authorized.PUT("account/:account_id/pass_hash", controller.EditPassword)
		authorized.POST("/account/:account_id/follow", controller.AddFollow)
		authorized.DELETE("/account/:account_id/follow", controller.DeleteFollow)
		authorized.POST("/notes", controller.AddNote)
		authorized.POST("/notes/:note_id/tags/:tag_id", controller.AddNoteTag)
		authorized.DELETE("/notes/:note_id/tags/:tag_id", controller.DeleteNoteTag)
		authorized.POST("/notes/:note_id/notability", controller.UploadNotability)
		authorized.POST("/notes/:note_id/goodnote", controller.UploadGoodnote)
		authorized.POST("/notes/:note_id/pdf", controller.UploadPdf)
		authorized.POST("/notes/:note_id/preview", controller.UploadPreview)
		authorized.GET("/notes/:note_id/save", controller.IsNoteSaved)
		authorized.POST("/notes/:note_id/save", controller.SaveNote)
		authorized.DELETE("/notes/:note_id/save", controller.DeleteSave)
		authorized.PATCH("/notes/:note_id", controller.EditNote)
		authorized.DELETE("/notes/:note_id", controller.DeleteNote)

	}
}
