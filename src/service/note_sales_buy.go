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
	// findUserByAccountID(account_id)

	db_err := persistence.DB.Model(&model.Download{}).FirstOrCreate(&buy).Error
	if db_err != nil {
		return db_err
	} else {
		return nil
	}
}

func FindSales(note_id int64) (int64, int64, error) {
	var cnt int64
	var bean int64
	var revenue int64
	err := persistence.DB.Model(&model.Download{}).Where("Note_id = ?", note_id).Count(&cnt).Error
	if err != nil {
		return 0, 0, err
	}

	bean, err = findBeansByNoteID(note_id)
	if err != nil {
		return 0, 0, err
	}
	revenue = bean * cnt
	return cnt, revenue, nil

}

func findBeansByNoteID(note_id int64) (int64, error) {
	var note model.Note
	var bean int64
	if res := persistence.DB.Where("ID = ?", note_id).First(&note); res.Error != nil {
		return 0, res.Error
	}
	bean = int64(note.Bean)
	return bean, nil
}
