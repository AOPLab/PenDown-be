package service

import (
	// "errors"

	"github.com/AOPLab/PenDown-be/src/model"
	"github.com/AOPLab/PenDown-be/src/persistence"
)

func BuyNote(user_id int64, note_id int64) error {
	buy := &model.Saved{
		User_id: user_id,
		Note_id: note_id,
	}

	db_err := persistence.DB.Model(&model.Download{}).FirstOrCreate(&buy).Error
	if db_err != nil {
		return db_err
	} else {
		return nil
	}
}

// func GetSave(user_id int64, note_id int64) (bool, error) {
// 	var save model.Saved
// 	err := persistence.DB.Where("User_id = ?", user_id).Where("Note_id = ?", note_id).First(&save).Error
// 	if err != nil {
// 		if err.Error() == "record not found" {
// 			return false, nil
// 		} else {
// 			return false, err
// 		}
// 	}
// 	return true, nil
// }
