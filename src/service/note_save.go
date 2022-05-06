package service

import (
	// "errors"

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

	db_err := persistence.DB.Where("User_id = ?", user_id).Where("Note_id = ?", note_id).FirstOrCreate(&save).Error
	if db_err != nil {
		return db_err
	} else {
		return nil
	}
}

func DeleteSave(user_id int64, note_id int64) error {
	// saved, save_err := GetSave(user_id, note_id)
	// if save_err != nil {
	// 	return save_err
	// }
	// if !saved {
	// 	return errors.New("Doesn't saved.")
	// }

	db_err := persistence.DB.Unscoped().Where("User_id = ? AND Note_id = ?", user_id, note_id).Delete(&model.Saved{}).Error
	if db_err != nil {
		return db_err
	} else {
		return nil
	}
}

func GetSave(user_id int64, note_id int64) (bool, error) {
	var save model.Saved
	err := persistence.DB.Where("User_id = ?", user_id).Where("Note_id = ?", note_id).First(&save).Error
	if err != nil {
		if err.Error() == "record not found" {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
}
