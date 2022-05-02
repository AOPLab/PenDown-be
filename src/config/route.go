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

	public_account := r.Group("api/account")
	{
		public_account.GET("/:account_id/profile", controller.GetPublicProfile)
		public_account.GET("/:account_id/followers", controller.GetFollowers)
		public_account.GET("/:account_id/followings", controller.GetFollowing)
		public_account.GET("/:account_id/following/:following_id", controller.GetFollow)
	}

	public_tag := r.Group("api/tag")
	{
		public_tag.GET("", controller.GetTags)
		public_tag.POST("", controller.AddTag)
	}

	public_note := r.Group("api/notes")
	{
		public_note.GET("/:note_id", controller.GetNote)
		public_note.GET("/:note_id/tags", controller.GetNoteTag)
		public_note.GET("/hot", controller.GetHotNote)
		public_note.GET("/tag/:tag_id", controller.GetNoteByTag)
		public_note.GET("/course/:course_id", controller.GetNoteByCourse)
	}

	public_school := r.Group("api/school")
	{
		public_school.GET("", controller.GetSchools)
		public_school.GET("/:school_id", controller.GetSchool)
		public_school.GET("/:school_id/course", controller.GetSchoolCourse)
	}

	public_course := r.Group("api/course")
	{
		public_course.GET("/:course_id", controller.GetCourse)
	}

	public_file := r.Group("/api/file")
	{
		public_file.GET("/preview", controller.GetPreviewFile)
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
	}

	authorized_account := r.Group("/api/account")
	authorized_account.Use(auth.AuthRequired)
	{
		authorized_account.GET("", controller.GetPrivateProfile)
		authorized_account.PATCH("", controller.EditProfile)
		authorized_account.PUT("/:account_id/pass_hash", controller.EditPassword)
		authorized_account.POST("/:account_id/follow", controller.AddFollow)
		authorized_account.DELETE("/:account_id/follow", controller.DeleteFollow)
	}

	authorized_note := r.Group("/api/notes")
	authorized_note.Use(auth.AuthRequired)
	{
		authorized_note.POST("", controller.AddNote)
		authorized_note.POST("/:note_id/tags/:tag_id", controller.AddNoteTag)
		authorized_note.DELETE("/:note_id/tags/:tag_id", controller.DeleteNoteTag)
		authorized_note.POST("/:note_id/notability", controller.UploadNotability)
		authorized_note.POST("/:note_id/goodnote", controller.UploadGoodnote)
		authorized_note.POST("/:note_id/pdf", controller.UploadPdf)
		authorized_note.POST("/:note_id/preview", controller.UploadPreview)
	}

	authorized_file := r.Group("/api/file")
	authorized_file.Use(auth.AuthRequired)
	{
		authorized_file.GET("/", controller.GetNoteFile)
	}
}
