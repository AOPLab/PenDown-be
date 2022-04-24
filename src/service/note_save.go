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

func SaveNote(user_id int64, note_id int64) error {
	save := &model.Saved{
		User_id: user_id,
		Note_id: note_id,
	}

	db_err := persistence.DB.Model(&model.Saved{}).Create(&save).Error
	if db_err != nil {
		return db_err
	} else {
		return nil
	}
}
