package controller

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
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
		return
	}

	file, file_err := c.FormFile("file")
	if file_err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": file_err.Error(),
		})
		return
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
	note, note_err := service.GetUserNoteById(user_id, note_id)
	if note_err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": note_err.Error(),
		})
		return
	}

	// Generate file path
	time := strconv.FormatInt(time.Now().Unix(), 10)
	filename := strconv.Itoa(int(note.ID)) + "_" + time + "_" + randStringRunes(5) + ".note"
	path := strconv.Itoa(int(note.Course.School_id)) + "/" + strconv.Itoa(int(note.Course_id)) + "/" + filename
	// fmt.Print(path)

	// Upload file
	upload_err := service.UploadFile(path, blobFile)
	if upload_err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": upload_err.Error(),
		})
		return
	}

	// Update filename
	update_err := service.UpdateNotabilityFilename(note.ID, filename)
	if update_err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": update_err.Error(),
		})
		return
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
		return
	}

	file, file_err := c.FormFile("file")
	if file_err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": file_err.Error(),
		})
		return
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
	note, note_err := service.GetUserNoteById(user_id, note_id)
	if note_err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": note_err.Error(),
		})
		return
	}

	// Generate file path
	time := strconv.FormatInt(time.Now().Unix(), 10)
	filename := strconv.Itoa(int(note.ID)) + "_" + time + "_" + randStringRunes(5) + ".goodnote"
	path := strconv.Itoa(int(note.Course.School_id)) + "/" + strconv.Itoa(int(note.Course_id)) + "/" + filename
	// fmt.Print(path)

	// Upload file
	upload_err := service.UploadFile(path, blobFile)
	if upload_err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": upload_err.Error(),
		})
		return
	}

	// Update filename
	update_err := service.UpdateGoodnoteFilename(note.ID, filename)
	if update_err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": update_err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"note_id":           note.ID,
		"goodnote_filename": filename,
	})
	return
}

// Upload PDF File
func UploadPdf(c *gin.Context) {
	user_id := c.MustGet("user_id").(int64)
	id := c.Params.ByName("note_id")
	note_id, parse_err := strconv.ParseInt(id, 0, 64)
	if parse_err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Note_id not exists",
		})
		return
	}

	file, file_err := c.FormFile("file")
	if file_err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": file_err.Error(),
		})
		return
	}

	// Check Content-Type
	if file.Header.Get("Content-Type") != "application/pdf" {
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
	defer blobFile.Close()

	// Get Note
	note, note_err := service.GetUserNoteById(user_id, note_id)
	if note_err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": note_err.Error(),
		})
		return
	}

	// Generate pdf file path
	time := strconv.FormatInt(time.Now().Unix(), 10)
	filename := strconv.Itoa(int(note.ID)) + "_" + time + "_" + randStringRunes(5) + ".pdf"
	path := strconv.Itoa(int(note.Course.School_id)) + "/" + strconv.Itoa(int(note.Course_id)) + "/" + filename
	// preview_filename := strconv.Itoa(int(note.ID)) + "_" + time + "_" + randStringRunes(5) + ".jpg"
	// preview_path := strconv.Itoa(int(note.Course.School_id)) + "/" + strconv.Itoa(int(note.Course_id)) + "/" + preview_filename
	// fmt.Print(path)

	// Upload pdf file
	upload_err := service.UploadFile(path, blobFile)
	if upload_err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": upload_err.Error(),
		})
		return
	}

	// blobFile2, err := file.Open()
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"error": err.Error(),
	// 	})
	// 	return
	// }
	// defer blobFile2.Close()

	// // Upload pdf preview file
	// preview_err := service.Fitz(preview_path, blobFile2)
	// if preview_err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"error": preview_err.Error(),
	// 	})
	// 	return
	// }

	// Update filename (including pdf and preview)
	update_err := service.UpdatePdfFilename(note.ID, filename)
	if update_err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": update_err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"note_id":      note.ID,
		"pdf_filename": filename,
	})
	return
}

// Upload Preview File
func UploadPreview(c *gin.Context) {
	user_id := c.MustGet("user_id").(int64)
	id := c.Params.ByName("note_id")
	note_id, parse_err := strconv.ParseInt(id, 0, 64)
	if parse_err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Note_id not exists",
		})
		return
	}

	file, file_err := c.FormFile("file")
	if file_err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": file_err.Error(),
		})
		return
	}

	// Check Content-Type
	if file.Header.Get("Content-Type") != "image/jpeg" {
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
	defer blobFile.Close()

	// Get Note
	note, note_err := service.GetUserNoteById(user_id, note_id)
	if note_err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": note_err.Error(),
		})
		return
	}

	// Generate preview file path
	time := strconv.FormatInt(time.Now().Unix(), 10)
	filename := strconv.Itoa(int(note.ID)) + "_" + time + "_" + randStringRunes(5) + ".png"
	path := strconv.Itoa(int(note.Course.School_id)) + "/" + strconv.Itoa(int(note.Course_id)) + "/" + filename

	// Upload preview image file
	upload_err := service.UploadFile(path, blobFile)
	if upload_err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": upload_err.Error(),
		})
		return
	}

	// Update filename (preview)
	update_err := service.UpdatePreviewFilename(note.ID, filename)
	if update_err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": update_err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"note_id":          note.ID,
		"preview_filename": filename,
	})
	return
}

func GetPreviewFile(c *gin.Context) {
	filename := c.Query("filename")
	id := c.Query("note_id")
	note_id, parse_err := strconv.ParseInt(id, 0, 64)
	if parse_err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "NoteIdPareseError",
		})
		return
	}

	note, note_err := service.GetNoteByIdWithCourse(note_id)
	if note_err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": note_err.Error(),
		})
		return
	}

	path := strconv.Itoa(int(note.Course.School_id)) + "/" + strconv.Itoa(int(note.Course_id)) + "/" + filename
	fmt.Print(note.Course)

	// Check file is image
	contain := strings.Contains(filename, "jpg")
	if contain == false {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "FileNameError",
		})
		return
	}

	file_url, sign_err := service.SignedFileUrl(path)
	if sign_err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": sign_err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"file_url": file_url,
	})
}

func GetNoteFile(c *gin.Context) {
	user_id := c.MustGet("user_id").(int64)
	filename := c.Query("filename")
	id := c.Query("note_id")
	note_id, parse_err := strconv.ParseInt(id, 0, 64)
	if parse_err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "NoteIdPareseError",
		})
		return
	}

	note, note_err := service.GetNoteByIdWithCourse(note_id)
	if note_err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": note_err.Error(),
		})
		return
	}

	path := strconv.Itoa(int(note.Course.School_id)) + "/" + strconv.Itoa(int(note.Course_id)) + "/" + filename

	// check user can download file
	if note.User_id != user_id || !service.CheckUserBuyNote(user_id, note.ID) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "NoAuthorization",
		})
		return
	}

	file_url, sign_err := service.SignedFileUrl(path)
	if sign_err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": sign_err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"file_url": file_url,
	})
}
