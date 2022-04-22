package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/AOPLab/PenDown-be/src/service"
	"github.com/gin-gonic/gin"
)

// Upload Notability File
func UploadNotability(c *gin.Context) {
	user_id := c.MustGet("user_id").(int64)
	id := c.Params.ByName("note_id")
	note_id, parse_err := strconv.ParseInt(id, 0, 64)
	if parse_err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Note_id not exists",
		})
	}

	file, file_err := c.FormFile("file")
	if file_err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": file_err.Error(),
		})
	}

	// Check Content-Type
	if file.Header.Get("Content-Type") != "application/octet-stream" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "FileTypeError",
		})
		return
	}

	blobFile, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Get Note
	note, note_err := service.GetNoteById(user_id, note_id)
	if note_err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": note_err.Error(),
		})
	}

	// Generate file path
	time := strconv.FormatInt(time.Now().Unix(), 10)
	filename := note.Title + "-" + time + ".note"
	path := strconv.Itoa(int(note.Course.School_id)) + "/" + strconv.Itoa(int(note.Course_id)) + "/" + filename
	fmt.Print(path)

	// Upload file
	upload_err := service.UploadFile(path, blobFile)
	if upload_err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": upload_err.Error(),
		})
	}

	// Update filename
	update_err := service.UpdateNotabilityFilename(note.ID, filename)
	if update_err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": update_err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"note_id":  note.ID,
		"filename": filename,
	})
	return
}
