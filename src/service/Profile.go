package service

import (
	"github.com/AOPLab/PenDown-be/src/model"
	"github.com/AOPLab/PenDown-be/src/persistence"
)

var PublicProfileFields = []string{"username", "description", "status", "bean"}
var PrivateProfileFields = []string{"username", "description", "status", "bean", "password", "google_id"}

func FindPublicProfile(account_id int64) (*model.User, int64, int64, int64, error) {
	user := &model.User{}
	var followers_num int64
	var following_num int64
	var note_num int64
	err := persistence.DB.Select(PublicProfileFields).Where("ID=?", account_id).First(&user).Error
	if err != nil {
		return nil, 0, 0, 0, err
	} else {
		persistence.DB.Model(&model.Follow{}).Where("Followee_id = ?", account_id).Count(&followers_num)
		persistence.DB.Model(&model.Follow{}).Where("Follower_id = ?", account_id).Count(&following_num)
		persistence.DB.Model(&model.Note{}).Where("User_id = ?", account_id).Count(&note_num)
		return user, followers_num, following_num, note_num, nil
	}
}

func FindPrivateProfile(account_id int64) (*model.User, int64, int64, int64, error) {
	user := &model.User{}
	var followers_num int64
	var following_num int64
	var note_num int64
	err := persistence.DB.Select(PrivateProfileFields).Where("ID=?", account_id).First(&user).Error
	if err != nil {
		return nil, 0, 0, 0, err
	} else {
		persistence.DB.Model(&model.Follow{}).Where("Followee_id = ?", account_id).Count(&followers_num)
		persistence.DB.Model(&model.Follow{}).Where("Follower_id = ?", account_id).Count(&following_num)
		persistence.DB.Model(&model.Note{}).Where("User_id = ?", account_id).Count(&note_num)
		return user, followers_num, following_num, note_num, nil
	}
}
