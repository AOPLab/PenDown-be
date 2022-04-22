package controller

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/AOPLab/PenDown-be/src/service"
	"github.com/gin-gonic/gin"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Generate random string
var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

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
	filename := strconv.Itoa(int(note.ID)) + "_" + time + "_" + randStringRunes(5) + ".note"
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
		"note_id":             note.ID,
		"notability_filename": filename,
	})
	return
}

// Upload Goodnotes File
func UploadGoodnote(c *gin.Context) {
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
	filename := strconv.Itoa(int(note.ID)) + "_" + time + "_" + randStringRunes(5) + ".goodnote"
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
	update_err := service.UpdateGoodnoteFilename(note.ID, filename)
	if update_err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": update_err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"note_id":           note.ID,
		"goodnote_filename": filename,
	})
	return
}

// Upload PDF File
func UploadPdf(c *gin.Context) {
	// user_id := c.MustGet("user_id").(int64)
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
	if file.Header.Get("Content-Type") != "application/pdf" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "FileTypeError",
		})
		return
	}

	service.Fitz()

	// blobFile, err := file.Open()
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"error": err.Error(),
	// 	})
	// 	return
	// }

	// // Get Note
	// note, note_err := service.GetNoteById(user_id, note_id)
	// if note_err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"error": note_err.Error(),
	// 	})
	// }

	// // Generate file path
	// time := strconv.FormatInt(time.Now().Unix(), 10)
	// filename := strconv.Itoa(int(note.ID)) + "_" + time + "_" + randStringRunes(5) + ".goodnote"
	// path := strconv.Itoa(int(note.Course.School_id)) + "/" + strconv.Itoa(int(note.Course_id)) + "/" + filename
	// fmt.Print(path)

	// // Upload file
	// upload_err := service.UploadFile(path, blobFile)
	// if upload_err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"error": upload_err.Error(),
	// 	})
	// }

	// // Update filename
	// update_err := service.UpdateGoodnoteFilename(note.ID, filename)
	// if update_err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"error": update_err.Error(),
	// 	})
	// }

	c.JSON(http.StatusOK, gin.H{
		"note_id":      note_id,
		"pdf_filename": file.Header.Get("Content-Type"),
	})
	return
}
