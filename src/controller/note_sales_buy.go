package controller

import (
	"net/http"
	"strconv"

	"github.com/AOPLab/PenDown-be/src/service"

	"github.com/gin-gonic/gin"
)

// Buy Note
func BuyNote(c *gin.Context) {
	// 要有足夠才能買
	// --> 用 note_id 查賣家
	// --> 更新雙方豆子(給賣家8成，要四捨五入)
	// --> New Download tuple (First or Create)
	// 共 5 支 API

	user_id := c.MustGet("user_id").(int64)
	id := c.Params.ByName("note_id")
	note_id, parse_err := strconv.ParseInt(id, 0, 64)

	if parse_err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Input format error",
		})
		return
	}

	err := service.BuyNote(user_id, note_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"is_bought": true,
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
