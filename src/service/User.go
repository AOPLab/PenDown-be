package service

import (
	"golang.org/x/crypto/bcrypt"

	"github.com/AOPLab/PenDown-be/src/model"
	"github.com/AOPLab/PenDown-be/src/persistence"
)

func hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func verifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func AddUser(username string, full_name string, email string, password string) (*model.User, error) {
	hashedPassword, err := hash(password)
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

func findUserByUsername(username string) (*model.User, error) {
	var user model.User
	if res := persistence.DB.Where("username = ?", username).Find(&user); res.Error != nil {
		return nil, res.Error
	}
	return &user, nil
}

func findUserByGoogleId(google_id int64) (*model.User, error) {
	var user model.User
	if res := persistence.DB.Where("google_id = ?", google_id).Find(&user); res.Error != nil {
		return nil, res.Error
	}
	return &user, nil
}

func VerifyLogin(username string, password string) (*model.User, error) {
	user, err := findUserByUsername(username)

	if err != nil {
		return nil, err
	}
	ver_err := verifyPassword(user.Password, password)
	if ver_err != nil {
		return nil, ver_err
	}

	return user, nil
}

func VerifyGoogleLogin(google_id int64) (*model.User, error) {
	user, err := findUserByGoogleId(google_id)
	if err != nil {
		return nil, err
	}

	return user, nil
}
