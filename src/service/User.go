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

func AddGoogleUser(google_id string, username string, full_name string, email string) (*model.User, error) {
	user := &model.User{
		Google_ID: google_id,
		Username:  username,
		Full_name: full_name,
		Email:     email,
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
	if res := persistence.DB.Where("username = ?", username).First(&user); res.Error != nil {
		return nil, res.Error
	}
	return &user, nil
}

func findUserByGoogleId(google_id string) (*model.User, error) {
	var user model.User
	if err := persistence.DB.Where("google_id = ?", google_id).First(&user).Error; err != nil {
		return nil, err
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

func VerifyGoogleLogin(google_id string) (*model.User, error) {
	user, err := findUserByGoogleId(google_id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

type SearchUserOutput struct {
	ID       int64  `json:"user_id"`
	Username string `json:"username"`
}

func SearchUser(q string, offset int, limit int) ([]SearchUserOutput, int64, error) {
	var results []SearchUserOutput
	var count int64
	searchName := "%" + q + "%"
	if err := persistence.DB.Limit(limit).Offset(offset).Table("users").Select("ID, username").Where("username LIKE ?", searchName).Find(&results).Error; err != nil {
		return results, 0, err
	}
	persistence.DB.Table("users").Where("username LIKE ?", searchName).Count(&count)
	return results, count, nil
}
