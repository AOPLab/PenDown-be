package controller

import (
	"net/http"
	"strconv"
	"time"

	"github.com/AOPLab/PenDown-be/src/service"
	"github.com/gin-gonic/gin"
)

type LazyNoteOutput struct {
	ID               int64     `json:"note_id"`
	User_id          int64     `json:"user_id"`
	Username         string    `json:"username"`
	Title            string    `json:"title"`
	View_cnt         int64     `json:"view_cnt"`
	Saved_cnt        int64     `json:"saved_cnt"`
	Note_type        string    `json:"note_type"`
	Preview_filename string    `json:"preview_filename"`
	CreatedAt        time.Time `json:"created_at"`
}

// q any string
// note_type: all, notability, goodnotes
// filter: people, tags, schools, courses, notes, templates
func Search(c *gin.Context) {
	q := c.Query("q")
	note_type := c.Query("type")
	filter := c.Query("filter")
	offset_o := c.Query("offset")

	offset, err := strconv.Atoi(offset_o)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "OffsetParseError",
		})
		return
	}

	switch filter {
	case "people":
		result, total_cnt, err := service.SearchUser(q, offset, 12)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"people":    result,
			"total_cnt": total_cnt,
		})
		return
	case "tags":
		result, total_cnt, err := service.SearchTag(q, offset, 12)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"tags":      result,
			"total_cnt": total_cnt,
		})
		return
	case "schools":
		result, total_cnt, err := service.SearchSchool(q, offset, 12)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"schools":   result,
			"total_cnt": total_cnt,
		})
		return
	case "courses":
		result, total_cnt, err := service.SearchCourse(q, offset, 12)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"courses":   result,
			"total_cnt": total_cnt,
		})
		return
	case "notes":
		result, total_cnt, err := service.SearchNote(q, offset, 12, note_type)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		var note_outputs []*LazyNoteOutput
		for _, note := range result {
			note_output := &LazyNoteOutput{
				ID:               note.ID,
				User_id:          note.User_id,
				Username:         note.Username,
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
				cnt = 0
			}
			note_output.Saved_cnt = cnt
			note_outputs = append(note_outputs, note_output)
		}
		if note_outputs == nil {
			var output [0]int
			c.JSON(http.StatusOK, gin.H{
				"notes":     output,
				"total_cnt": total_cnt,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"notes":     note_outputs,
			"total_cnt": total_cnt,
		})
		return
	case "templates":
		result, total_cnt, err := service.SearchTemplate(q, offset, 12, note_type)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		var note_outputs []*LazyNoteOutput
		for _, note := range result {
			note_output := &LazyNoteOutput{
				ID:               note.ID,
				User_id:          note.User_id,
				Username:         note.Username,
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
				cnt = 0
			}
			note_output.Saved_cnt = cnt
			note_outputs = append(note_outputs, note_output)
		}
		if note_outputs == nil {
			var output [0]int
			c.JSON(http.StatusOK, gin.H{
				"templates": output,
				"total_cnt": total_cnt,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"templates": note_outputs,
			"total_cnt": total_cnt,
		})
		return
	default:
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "NoFilter",
		})
		return
	}

}
