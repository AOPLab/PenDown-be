package service

import (
	"errors"
	"math"

	"github.com/AOPLab/PenDown-be/src/model"
	"github.com/AOPLab/PenDown-be/src/persistence"
	"gorm.io/gorm"
)

func BuyNote(user_id int64, note_id int64) error {
	// Able to buy if enough beans
	var user_bean int64
	var price_bean int64
	user, find_err := findUserByAccountID(user_id)
	if find_err != nil {
		return find_err
	}
	user_bean = user.Bean

	price_bean, err := findBeansByNoteID(note_id)
	if err != nil {
		return err
	}

	if user_bean < price_bean {
		return errors.New("No enough beans!")
	}

	// Find seller by note ID
	seller_id, find_err := findUserByNoteID(note_id)
	if find_err != nil {
		return find_err
	}

	// Update both's beans, seller get 80% of the price
	err = BeanTransaction(seller_id, user_id, price_bean)
	if err != nil {
		return err
	}

	// New Download tuple
	buy := &model.Download{
		User_id: user_id,
		Note_id: note_id,
	}
	db_err := persistence.DB.Model(&model.Download{}).Create(&buy).Error
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

func findUserByNoteID(note_id int64) (int64, error) {
	var note model.Note
	var user_id int64
	if res := persistence.DB.Where("ID = ?", note_id).First(&note); res.Error != nil {
		return 0, res.Error
	}
	user_id = int64(note.User_id)
	return user_id, nil
}

func BeanTransaction(seller_id int64, buyer_id int64, bean int64) error {
	var revenue float64
	revenue = float64(bean)
	revenue = math.Round(revenue * 0.8)

	seller := &model.User{ID: seller_id}
	db_err := persistence.DB.Model(&seller).Update("Bean", gorm.Expr("Bean + ?", revenue)).Error
	if db_err != nil {
		return db_err
	}

	buyer := &model.User{ID: buyer_id}
	db_err = persistence.DB.Model(&buyer).Update("Bean", gorm.Expr("Bean - ?", bean)).Error
	if db_err != nil {
		return db_err
	}

	return nil
}
