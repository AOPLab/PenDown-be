package controller

import (
	"net/http"
	"strconv"
	"time"

	"github.com/AOPLab/PenDown-be/src/service"

	"github.com/gin-gonic/gin"
)

type NoteBrief struct {
	Note_ID          int64     `json:"note_id"`
	Account_ID       int64     `json:"user_id"`
	Username         string    `json:"username"`
	Title            string    `json:"title"`
	View_cnt         int64     `json:"view_cnt"`
	Saved_cnt        int64     `json:"saved_cnt"`
	Note_type        string    `json:"note_type"`
	Preview_filename string    `json:"preview_filename"`
	Preview_url      string    `json:"preview_url"`
	CreatedAt        time.Time `json:"created_at"`
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
			Note_ID:          note.ID,
			Account_ID:       note.User.ID,
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
		if note.Preview_filename != "" {
			path := strconv.Itoa(int(note.ID)) + "/" + note.Preview_filename
			file_url, _ := service.SignedFileUrl(path)
			note_output.Preview_url = file_url
		}
		note_outputs = append(note_outputs, note_output)
	}
	if note_outputs == nil {
		var output [0]int
		c.JSON(http.StatusOK, gin.H{
			"notes":     output,
			"total_cnt": note_cnt,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"notes":     note_outputs,
		"total_cnt": note_cnt,
	})
}

func GetNoteByTag(c *gin.Context) {

	offset := c.Query("offset")
	filter := c.Query("filter")
	noteType := c.Query("type")
	id := c.Params.ByName("tag_id")
	offset_num, parse_err := strconv.ParseInt(offset, 0, 64)
	if parse_err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Offset Parese Error",
		})
		return
	}
	tag_id, parse_err := strconv.ParseInt(id, 0, 64)
	if parse_err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Tag Id Parse Error",
		})
		return
	}

	full_filter := noteType + "-" + filter
	notes, note_cnt, note_err := service.GetNotesByTag(tag_id, full_filter, offset_num)

	if note_err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": note_err.Error(),
		})
		return
	}
	var note_outputs []*NoteBrief
	for _, note := range notes {
		note_output := &NoteBrief{
			Note_ID:          note.Note_ID,
			Account_ID:       note.ID,
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
			c.JSON(http.StatusBadRequest, gin.H{
				"error": note_err.Error(),
			})
			return
		}
		note_output.Saved_cnt = cnt
		if note.Preview_filename != "" {
			path := strconv.Itoa(int(note.ID)) + "/" + note.Preview_filename
			file_url, _ := service.SignedFileUrl(path)
			note_output.Preview_url = file_url
		}
		note_outputs = append(note_outputs, note_output)
	}
	if note_outputs == nil {
		var output [0]int
		c.JSON(http.StatusOK, gin.H{
			"notes":     output,
			"total_cnt": note_cnt,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"notes":     note_outputs,
		"total_cnt": note_cnt,
	})
}

func GetNoteByCourse(c *gin.Context) {

	offset := c.Query("offset")
	filter := c.Query("filter")
	noteType := c.Query("type")
	id := c.Params.ByName("course_id")
	offset_num, parse_err := strconv.ParseInt(offset, 0, 64)
	if parse_err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Offset Parese Error",
		})
		return
	}
	course_id, parse_err := strconv.ParseInt(id, 0, 64)
	if parse_err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Tag Id Parse Error",
		})
		return
	}

	full_filter := noteType + "-" + filter
	notes, note_cnt, note_err := service.GetNotesByCourse(course_id, full_filter, offset_num)

	if note_err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": note_err.Error(),
		})
		return
	}
	var note_outputs []*NoteBrief
	for _, note := range notes {
		note_output := &NoteBrief{
			Note_ID:          note.ID,
			Account_ID:       note.User.ID,
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
		if note.Preview_filename != "" {
			path := strconv.Itoa(int(note.ID)) + "/" + note.Preview_filename
			file_url, _ := service.SignedFileUrl(path)
			note_output.Preview_url = file_url
		}
		note_outputs = append(note_outputs, note_output)
	}
	if note_outputs == nil {
		var output [0]int
		c.JSON(http.StatusOK, gin.H{
			"notes":     output,
			"total_cnt": note_cnt,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"notes":     note_outputs,
		"total_cnt": note_cnt,
	})
}

func GetNotesByUserIdPublic(c *gin.Context) {
	offset := c.Query("offset")
	filter := c.Query("filter")
	noteType := c.Query("type")
	id := c.Params.ByName("account_id")
	offset_num, parse_err := strconv.ParseInt(offset, 0, 64)
	if parse_err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Offset Parese Error",
		})
		return
	}
	account_id, parse_err := strconv.ParseInt(id, 0, 64)
	if parse_err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Tag Id Parse Error",
		})
		return
	}

	full_filter := noteType + "-" + filter
	notes, note_cnt, note_err := service.GetNotesByUserId(account_id, full_filter, offset_num)

	if note_err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": note_err.Error(),
		})
		return
	}
	var note_outputs []*NoteBrief
	for _, note := range notes {
		note_output := &NoteBrief{
			Note_ID:          note.ID,
			Account_ID:       note.User.ID,
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
		if note.Preview_filename != "" {
			path := strconv.Itoa(int(note.ID)) + "/" + note.Preview_filename
			file_url, _ := service.SignedFileUrl(path)
			note_output.Preview_url = file_url
		}
		note_outputs = append(note_outputs, note_output)
	}
	if note_outputs == nil {
		var output [0]int
		c.JSON(http.StatusOK, gin.H{
			"notes":     output,
			"total_cnt": note_cnt,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"notes":     note_outputs,
		"total_cnt": note_cnt,
	})
}

func GetOwnNotes(c *gin.Context) {

	offset := c.Query("offset")
	filter := c.Query("filter")
	noteType := c.Query("type")
	offset_num, parse_err := strconv.ParseInt(offset, 0, 64)
	if parse_err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Offset Parese Error",
		})
		return
	}
	account_id, consistence_err := checkIdConsistence(c)
	if consistence_err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": consistence_err.Error(),
		})
		return
	}

	switch filter {
	case "uploaded":
		full_filter := noteType + "-recent"
		notes, note_cnt, note_err := service.GetNotesByUserId(account_id, full_filter, offset_num)

		if note_err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": note_err.Error(),
			})
			return
		}
		var note_outputs []*NoteBrief
		for _, note := range notes {
			note_output := &NoteBrief{
				Note_ID:          note.ID,
				Account_ID:       note.User.ID,
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
			if note.Preview_filename != "" {
				path := strconv.Itoa(int(note.ID)) + "/" + note.Preview_filename
				file_url, _ := service.SignedFileUrl(path)
				note_output.Preview_url = file_url
			}
			note_outputs = append(note_outputs, note_output)
		}
		if note_outputs == nil {
			var output [0]int
			c.JSON(http.StatusOK, gin.H{
				"notes":     output,
				"total_cnt": note_cnt,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"notes":     note_outputs,
			"total_cnt": note_cnt,
		})

	case "saved":

		notes, note_cnt, note_err := service.GetOwnSavedNotes(account_id, noteType, offset_num)

		if note_err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": note_err.Error(),
			})
			return
		}
		var note_outputs []*NoteBrief
		for _, note := range notes {
			note_output := &NoteBrief{
				Note_ID:          note.Note_id,
				Account_ID:       note.User.ID,
				Username:         note.User.Username,
				Title:            note.Note.Title,
				Preview_filename: note.Note.Preview_filename,
				View_cnt:         note.Note.View_cnt,
				CreatedAt:        note.Note.CreatedAt,
			}
			// add note type
			if note.Note.Notability_filename != "" && note.Note.Goodnotes_filename != "" {
				note_output.Note_type = "All"
			} else if note.Note.Notability_filename != "" {
				note_output.Note_type = "Notability"
			} else if note.Note.Goodnotes_filename != "" {
				note_output.Note_type = "Goodnotes"
			}
			// calculate saved cnt
			cnt, save_err := service.GetNoteSavedCnt(note.Note_id)
			if save_err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": note_err.Error(),
				})
				return
			}
			note_output.Saved_cnt = cnt
			if note.Note.Preview_filename != "" {
				path := strconv.Itoa(int(note.ID)) + "/" + note.Note.Preview_filename
				file_url, _ := service.SignedFileUrl(path)
				note_output.Preview_url = file_url
			}
			note_outputs = append(note_outputs, note_output)
		}
		if note_outputs == nil {
			var output [0]int
			c.JSON(http.StatusOK, gin.H{
				"notes":     output,
				"total_cnt": note_cnt,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"notes":     note_outputs,
			"total_cnt": note_cnt,
		})
	case "library":

		notes, note_cnt, note_err := service.GetOwnLibraryNotes(account_id, noteType, offset_num)

		if note_err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": note_err.Error(),
			})
			return
		}
		var note_outputs []*NoteBrief
		for _, note := range notes {
			note_output := &NoteBrief{
				Note_ID:          note.Note_id,
				Account_ID:       note.User.ID,
				Username:         note.User.Username,
				Title:            note.Note.Title,
				Preview_filename: note.Note.Preview_filename,
				View_cnt:         note.Note.View_cnt,
				CreatedAt:        note.Note.CreatedAt,
			}
			// add note type
			if note.Note.Notability_filename != "" && note.Note.Goodnotes_filename != "" {
				note_output.Note_type = "All"
			} else if note.Note.Notability_filename != "" {
				note_output.Note_type = "Notability"
			} else if note.Note.Goodnotes_filename != "" {
				note_output.Note_type = "Goodnotes"
			}
			// calculate saved cnt
			cnt, save_err := service.GetNoteSavedCnt(note.Note_id)
			if save_err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": note_err.Error(),
				})
				return
			}
			note_output.Saved_cnt = cnt
			if note.Note.Preview_filename != "" {
				path := strconv.Itoa(int(note.ID)) + "/" + note.Note.Preview_filename
				file_url, _ := service.SignedFileUrl(path)
				note_output.Preview_url = file_url
			}
			note_outputs = append(note_outputs, note_output)
		}
		if note_outputs == nil {
			var output [0]int
			c.JSON(http.StatusOK, gin.H{
				"notes":     output,
				"total_cnt": note_cnt,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"notes":     note_outputs,
			"total_cnt": note_cnt,
		})
		return
	default:
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "NoFilter",
		})
		return
	}

}
