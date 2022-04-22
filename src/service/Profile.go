package service

import (
	"errors"
	"github.com/AOPLab/PenDown-be/src/model"
	"github.com/AOPLab/PenDown-be/src/persistence"
)

var PublicProfileFields = []string{"user_id", "username", "description", "status", "bean"}
var PrivateProfileFields = []string{"user_id", "username", "description", "status", "bean", "password", "google_id"}

type EditAccountInput struct {
	Username    string `json:"username"`
	Full_name   string `json:"full_name"`
	Email       string `json:"email"`
	Description string `json:"description"`
}

type EditPasswordInput struct {
	Old_password string `json:"old_password"`
	New_password string `json:"new_password" binding:"required"`
}

func findUserByAccountID(account_id int64) (*model.User, error) {
	var user model.User
	if res := persistence.DB.Where("ID = ?", account_id).First(&user); res.Error != nil {
		return nil, res.Error
	}
	return &user, nil
}

func FindPublicProfile(account_id int64) (*model.User, int64, int64, int64, error) {
	var followers_num int64
	var following_num int64
	var note_num int64
	user, find_err := findUserByAccountID(account_id)
	if find_err != nil {
		return nil, 0, 0, 0, find_err
	} else {
		persistence.DB.Model(&model.Follow{}).Where("Followee_id = ?", account_id).Count(&followers_num)
		persistence.DB.Model(&model.Follow{}).Where("Follower_id = ?", account_id).Count(&following_num)
		persistence.DB.Model(&model.Note{}).Where("User_id = ?", account_id).Count(&note_num)
		return user, followers_num, following_num, note_num, nil
	}
}

func FindPrivateProfile(account_id int64) (*model.User, int64, int64, int64, error) {
	var followers_num int64
	var following_num int64
	var note_num int64
	user, find_err := findUserByAccountID(account_id)
	if find_err != nil {
		return nil, 0, 0, 0, find_err
	} else {
		persistence.DB.Model(&model.Follow{}).Where("Followee_id = ?", account_id).Count(&followers_num)
		persistence.DB.Model(&model.Follow{}).Where("Follower_id = ?", account_id).Count(&following_num)
		persistence.DB.Model(&model.Note{}).Where("User_id = ?", account_id).Count(&note_num)
		return user, followers_num, following_num, note_num, nil
	}
}

func EditProfile(account_id int64, form EditAccountInput) error {
	err := persistence.DB.Model(&model.User{}).Where("ID = ?", account_id).Updates(map[string]interface{}{"Username": form.Username, "Full_name": form.Full_name, "Email": form.Email, "Description": form.Description}).Error
	if err != nil {
		return err
	} else {
		return nil
	}
}

func EditPassword(account_id int64, form EditPasswordInput) error {
	user, find_err := findUserByAccountID(account_id)
	if find_err != nil {
		return find_err
	}

	if user.Google_ID != "" {
		if form.Old_password != "" {
			return errors.New("crypto/bcrypt: hashedPassword is not the hash of the given password")
		}
	} else {
		ver_err := verifyPassword(user.Password, form.Old_password)
		if ver_err != nil {
			return ver_err
		}
	}

	hashedPassword, hash_err := hash(form.New_password)
	if hash_err != nil {
		return hash_err
	}

	err := persistence.DB.Model(&model.User{}).Where("ID = ?", account_id).Update("Password", hashedPassword).Error
	if err != nil {
		return err
	} else {
		return nil
	}
}
