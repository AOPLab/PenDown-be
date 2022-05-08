package service

import (
	"time"

	"github.com/AOPLab/PenDown-be/src/model"
	"github.com/AOPLab/PenDown-be/src/persistence"
	"gorm.io/gorm"
)

type EditNoteInput struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Course_id   int64  `json:"course_id" binding:"required"`
	Bean        int    `json:"bean" binding:"required"`
	Is_template bool   `json:"is_template"`
}

func EditNote(user_id int64, note_id int64, form EditNoteInput) error {
	var note model.Note
	res := persistence.DB.Where("User_id = ? AND ID = ?", user_id, note_id).First(&note)
	if res.Error != nil {
		return res.Error
	}

	courseIdBefore := note.Course_id
	courseIdAfter := form.Course_id
	if courseIdBefore != courseIdAfter {
		courseBefore := &model.Course{ID: courseIdBefore}
		db_err := persistence.DB.Model(&courseBefore).Update("Note_cnt", gorm.Expr("Note_cnt - ?", 1)).Error
		if db_err != nil {
			return db_err
		}
		courseAfter := &model.Course{ID: courseIdAfter}
		db_err = persistence.DB.Model(&courseAfter).Update("Note_cnt", gorm.Expr("Note_cnt + ?", 1)).Update("Last_updated_time", time.Now()).Error
		if db_err != nil {
			return db_err
		}
	}

	err := res.Updates(map[string]interface{}{"Title": form.Title, "Description": form.Description, "Course_id": form.Course_id, "Bean": form.Bean, "Is_template": form.Is_template}).Error
	if err != nil {
		return err
	} else {
		return nil
	}
}

func DeleteNote(user_id int64, note_id int64) error {
	db_err := persistence.DB.Unscoped().Where("User_id = ? AND ID = ?", user_id, note_id).Delete(&model.Note{}).Error
	if db_err != nil {
		return db_err
	} else {
		return nil
	}
}

func GetMyNote(user_id int64, note_id int64) (bool, error) {
	var myNote model.Note
	err := persistence.DB.Where("User_id = ?", user_id).Where("ID = ?", note_id).First(&myNote).Error
	if err != nil {
		if err.Error() == "record not found" {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
}
