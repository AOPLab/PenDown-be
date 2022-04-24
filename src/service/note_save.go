package service

import (
	"github.com/AOPLab/PenDown-be/src/model"
	"github.com/AOPLab/PenDown-be/src/persistence"
)

func SavedNote(user_id int64, note_id int64) (bool, error) {
	var saved model.Saved
	err := persistence.DB.Where("User_id = ?", user_id).Where("Note_id = ?", note_id).First(&saved).Error
	if err != nil {
		if err.Error() == "record not found" {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
}
