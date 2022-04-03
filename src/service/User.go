package service

import (
	"golang.org/x/crypto/bcrypt"

	"github.com/AOPLab/PenDown-be/src/model"
	"github.com/AOPLab/PenDown-be/src/persistence"
)

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func AddUser(username string, full_name string, email string, password string) (*model.User, error) {
	hashedPassword, err := Hash(password)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		Username:  username,
		Full_name: full_name,
		Email:     email,
		Password:  string(hashedPassword),
	}

	db_err := persistence.DB.Model(&model.User{}).Create(&user).Error
	if db_err != nil {
		return nil, db_err
	} else {
		return user, nil
	}
}

func FindUserByUsername(username string) (*model.User, error) {
	var user model.User
	if res := persistence.DB.Where("username = ?", username).Find(&user); res.Error != nil {
		return nil, res.Error
	}
	return &user, nil
}
