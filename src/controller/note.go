package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/AOPLab/PenDown-be/src/auth"
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

type TagInfo struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type NoteOutput struct {
	ID                  int64     `json:"note_id"`
	Account_id          int64     `json:"account_id"`
	Username            string    `json:"username"`
	Title               string    `json:"title"`
	Description         string    `json:"description"`
	View_cnt            int64     `json:"view_cnt"`
	Saved_cnt           int64     `json:"saved_cnt"`
	Course_id           int64     `json:"course_id"`
	Course_name         string    `json:"course_name"`
	Course_no           string    `json:"course_no"`
	School_id           int64     `json:"school_id"`
	School_name         string    `json:"school_name"`
	Is_template         bool      `json:"is_template"`
	Note_type           string    `json:"note_type"`
	Bean                int       `json:"bean"`
	Pdf_filename        string    `json:"pdf_filename"`
	Preview_filename    string    `json:"preview_filename"`
	Goodnotes_filename  string    `json:"goodnotes_filename"`
	Notability_filename string    `json:"notability_filename"`
	CreatedAt           time.Time `json:"created_at"`
}

type NoteBrief struct {
	ID               int64     `json:"note_id"`
	Username         string    `json:"username"`
	Title            string    `json:"title"`
	View_cnt         int64     `json:"view_cnt"`
	Saved_cnt        int64     `json:"saved_cnt"`
	Note_type        string    `json:"note_type"`
	Preview_filename string    `json:"preview_filename"`
	CreatedAt        time.Time `json:"created_at"`
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

	var tags []*TagInfo
	for _, noteTag := range *noteTags {
		fmt.Println(noteTag.Tag_id)
		var tag TagInfo
		tag.ID = noteTag.Tag.ID
		tag.Name = noteTag.Tag.Tag_name
		tags = append(tags, &tag)
	}

	c.JSON(http.StatusOK, gin.H{
		"tags": tags,
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

func GetNote(c *gin.Context) {
	id := c.Params.ByName("note_id")

	note_id, parse_err := strconv.ParseInt(id, 0, 64)
	if parse_err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "NoteIdParseError",
		})
		return
	}

	if c.GetHeader("Authorization") != "" {
		// Get note with filename
		auth.AuthRequired(c)
		if c.Writer.Status() == 401 {
			return
		}
	}

	note, note_err := service.GetNoteById(note_id)
	if note_err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": note_err.Error(),
		})
		return
	}

	note_output := &NoteOutput{
		ID:               note.ID,
		Username:         note.User.Username,
		Account_id:       note.User_id,
		Title:            note.Title,
		Description:      note.Description,
		Course_name:      note.Course.Course_name,
		Course_id:        note.Course_id,
		Course_no:        note.Course.Course_no,
		School_name:      note.Course.School.School_name,
		School_id:        note.Course.School_id,
		Is_template:      note.Is_template,
		Bean:             note.Bean,
		Preview_filename: note.Preview_filename,
		View_cnt:         note.View_cnt,
		CreatedAt:        note.CreatedAt,
	}

	// add note type
	if note.Notability_filename != "" && note.Goodnotes_filename != "" {
		note_output.Note_type = "All"
	} else if note.Notability_filename != "" {
		note_output.Note_type = "Notability"
	} else if note.Goodnotes_filename != "" {
		note_output.Note_type = "Goodnotes"
	}

	// calculate saved cnt
	cnt, save_err := service.GetNoteSavedCnt(note.ID)
	if save_err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": note_err.Error(),
		})
		return
	}
	note_output.Saved_cnt = cnt

	if c.GetHeader("Authorization") != "" {
		// Get note with filename
		user_id := c.MustGet("user_id").(int64)

		// Check bought or not
		if user_id == note.User_id || service.CheckUserBuyNote(user_id, note.ID) {
			note_output.Pdf_filename = note.Pdf_filename
			note_output.Notability_filename = note.Notability_filename
			note_output.Goodnotes_filename = note.Goodnotes_filename
		}
	}

	service.UpdateNoteViewCnt(note.ID, note.View_cnt+1)

	c.JSON(http.StatusOK, note_output)
}

func GetHotNote(c *gin.Context) {
	offset := c.Query("offset")
	offset_num, parse_err := strconv.ParseInt(offset, 0, 64)
	if parse_err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Offset Parese Error",
		})
		return
	}
	notes, note_cnt, note_err := service.GetNotes("hot", offset_num)

	if note_err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": note_err.Error(),
		})
		return
	}

	var note_outputs []*NoteBrief
	for _, note := range notes {
		note_output := &NoteBrief{
			ID:               note.ID,
			Username:         note.User.Username,
			Title:            note.Title,
			Preview_filename: note.Preview_filename,
			View_cnt:         note.View_cnt,
			CreatedAt:        note.CreatedAt,
		}
		// add note type
		if note.Notability_filename != "" && note.Goodnotes_filename != "" {
			note_output.Note_type = "All"
		} else if note.Notability_filename != "" {
			note_output.Note_type = "Notability"
		} else if note.Goodnotes_filename != "" {
			note_output.Note_type = "Goodnotes"
		}
		// calculate saved cnt
		cnt, save_err := service.GetNoteSavedCnt(note.ID)
		if save_err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": note_err.Error(),
			})
			return
		}
		note_output.Saved_cnt = cnt
		note_outputs = append(note_outputs, note_output)
	}

	c.JSON(http.StatusOK, gin.H{
		"notes":     note_outputs,
		"total_cnt": note_cnt,
	})
}

func GetNoteByTag(c *gin.Context) {

	// offset := c.Query("offset")
	// filter := c.Query("filter")
	// noteType := c.Query("type")
	id := c.Params.ByName("tag_id")
	// offset_num, parse_err := strconv.ParseInt(offset, 0, 64)
	// if parse_err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"error": "Offset Parese Error",
	// 	})
	// 	return
	// }
	tag_id, parse_err := strconv.ParseInt(id, 0, 64)
	if parse_err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Tag Id Parse Error",
		})
		return
	}

	filter := "notability-recent"
	notes, note_cnt, note_err := service.GetNoteByTag(tag_id, filter, 0)

	if note_err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": note_err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"notes":       notes,
			"total_count": note_cnt,
		})
	}

	// var note_outputs []*NoteBrief
	// for _, note := range notes {
	// 	note_output := &NoteBrief{
	// 		ID:               note.ID,
	// 		Username:         note.User.Username,
	// 		Title:            note.Title,
	// 		Preview_filename: note.Preview_filename,
	// 		View_cnt:         note.View_cnt,
	// 		CreatedAt:        note.CreatedAt,
	// 	}
	// 	// add note type
	// 	if note.Notability_filename != "" && note.Goodnotes_filename != "" {
	// 		note_output.Note_type = "All"
	// 	} else if note.Notability_filename != "" {
	// 		note_output.Note_type = "Notability"
	// 	} else if note.Goodnotes_filename != "" {
	// 		note_output.Note_type = "Goodnotes"
	// 	}
	// 	// calculate saved cnt
	// 	cnt, save_err := service.GetNoteSavedCnt(note.ID)
	// 	if save_err != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{
	// 			"error": note_err.Error(),
	// 		})
	// 		return
	// 	}
	// 	note_output.Saved_cnt = cnt
	// 	note_outputs = append(note_outputs, note_output)
	// }

}
