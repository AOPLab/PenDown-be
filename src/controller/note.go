package controller

import (
	"net/http"
	"strconv"

	"github.com/AOPLab/PenDown-be/src/service"

	"github.com/gin-gonic/gin"
)

type AddNoteInput struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Is_template *bool  `json:"is_template" binding:"required"`
	Course_id   int64  `json:"course_id" binding:"required"`
	Bean        *int   `json:"bean" binding:"required"`
}

// Create Note and add tags
func AddNote(c *gin.Context) {
	var form AddNoteInput
	user_id := c.MustGet("user_id").(int64)

	bindErr := c.BindJSON(&form)
	if bindErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": bindErr.Error(),
		})
		return
	}

	note, err := service.AddNote(user_id, form.Title, form.Description, *form.Is_template, form.Course_id, *form.Bean)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"note_id": note.ID,
	})
}

// TODO: get note tags
func GetNoteTag(c *gin.Context) {
	id := c.Params.ByName("note_id")
	note_id, parse_err := strconv.ParseInt(id, 0, 64)
	if parse_err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Note_id not exists",
		})
		return
	}

	noteTags, noteTag_err := service.GetNoteTag(note_id)
	if noteTag_err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": noteTag_err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": *noteTags,
	})
}

func AddNoteTag(c *gin.Context) {
	id := c.Params.ByName("note_id")
	note_id, parse_err := strconv.ParseInt(id, 0, 64)
	if parse_err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Note_id not exists",
		})
		return
	}

	id = c.Params.ByName("tag_id")
	tag_id, parse_err := strconv.ParseInt(id, 0, 64)
	if parse_err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Tag_id not exists",
		})
		return
	}

	user_id := c.MustGet("user_id").(int64)

	err := service.AddNoteTag(user_id, note_id, tag_id)
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

func DeleteNoteTag(c *gin.Context) {
	id := c.Params.ByName("note_id")
	note_id, parse_err := strconv.ParseInt(id, 0, 64)
	if parse_err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Note_id not exists",
		})
		return
	}

	id = c.Params.ByName("tag_id")
	tag_id, parse_err := strconv.ParseInt(id, 0, 64)
	if parse_err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Tag_id not exists",
		})
		return
	}

	user_id := c.MustGet("user_id").(int64)

	err := service.DeleteNoteTag(user_id, note_id, tag_id)
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

// TODO: GET NOTE
// CHECK USER DOWNLOAD AUTHENTICATION
func GetNote(c *gin.Context) {
	id := c.Params.ByName("note_id")
	note_id, parse_err := strconv.ParseInt(id, 0, 64)
	if parse_err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Note_id not exists",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"note_id": note_id,
	})
}

// response example
// {
//     "note_id":
//     "username":
//     "account_id":
//     "title":
//     "description":
//     "course_name":
//     "course_id":
//     "course_no":
//     "school_name":
//     "school_id":
//     "note_type":
//     "is_template":
//     "bean":
//     "preview_filename": # filename 前端拿 file_name 在打另外一支 API 去拿網址，再拿網址去開資料
//     "pdf_filename":
//     "notability_filename":
//     "goodnote_filename":
//     "view_cnt":
//     "saved_cnt":
//     "created_at":
//     "tags": [
//         {
//             "tag_id": ""
//             "tag_name": ""
//         }
//     ]
// }
