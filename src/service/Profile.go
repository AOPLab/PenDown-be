package service

import (
	"github.com/AOPLab/PenDown-be/src/model"
	"github.com/AOPLab/PenDown-be/src/persistence"
)

var PublicProfileFields = []string{"username", "description", "status", "bean"}

func FindPublicProfile(account_id int64) (*model.User, int64, int64, int64, error) {
	user := &model.User{}
	var followers_num int64
	var following_num int64
	var note_num int64
	err := persistence.DB.Select(PublicProfileFields).Where("account_id=?", account_id).Find(&user).Error
	if err != nil {
		return nil, 0, 0, 0, err
	} else {
		persistence.DB.Model(&model.Follow{}).Where("Followee = ?", account_id).Count(&followers_num)
		persistence.DB.Model(&model.Follow{}).Where("Follower = ?", account_id).Count(&following_num)
		persistence.DB.Model(&model.Note{}).Where("User_ID = ?", account_id).Count(&note_num)
		return user, followers_num, following_num, note_num, nil
	}
}
