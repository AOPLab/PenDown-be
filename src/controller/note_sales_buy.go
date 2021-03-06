package controller

import (
	"net/http"
	"strconv"

	"github.com/AOPLab/PenDown-be/src/service"

	"github.com/gin-gonic/gin"
)

// Buy Note
func BuyNote(c *gin.Context) {
	user_id := c.MustGet("user_id").(int64)
	id := c.Params.ByName("note_id")
	note_id, parse_err := strconv.ParseInt(id, 0, 64)

	if parse_err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Input format error",
		})
		return
	}

	note, err := service.BuyNote(user_id, note_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"pdf_filename":        note.Pdf_filename,
		"notability_filename": note.Notability_filename,
		"goodnotes_filename":  note.Goodnotes_filename,
	})

	return
}

// GET Note Sales
func GetNoteSales(c *gin.Context) {
	id := c.Params.ByName("note_id")
	note_id, pasre_err := strconv.ParseInt(id, 0, 64)
	if pasre_err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "input format error",
		})
		return
	}
	cnt, revenue, err := service.FindSales(note_id)

	if err != nil {
		if err.Error() == "record not found" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "note_id not exists",
			})
			return
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"sales_count": cnt,
			"revenue":     revenue,
		})
	}
	return
}

// // Cancel Save
// func DeleteSave(c *gin.Context) {
// 	user_id := c.MustGet("user_id").(int64)
// 	id := c.Params.ByName("note_id")
// 	note_id, parse_err := strconv.ParseInt(id, 0, 64)

// 	if parse_err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": "Input format error",
// 		})
// 		return
// 	}

// 	err := service.DeleteSave(user_id, note_id)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"success": true,
// 	})
// 	return
// }
