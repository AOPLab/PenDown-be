package service

import (
	"errors"

	"github.com/AOPLab/PenDown-be/src/model"
	"github.com/AOPLab/PenDown-be/src/persistence"
)

func EditNote(user_id int64, note_id int64) error {
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

func DeleteNote(user_id int64, note_id int64) error {
	myNote, myNote_err := GetMyNote(user_id, note_id)
	if myNote_err != nil {
		return myNote_err
	}
	if !myNote {
		return errors.New("Note doesn't exist.")
	}

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
