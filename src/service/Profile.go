package service

import (
	"github.com/AOPLab/PenDown-be/src/model"
	"github.com/AOPLab/PenDown-be/src/persistence"
)

var PublicProfileFields = []string{"username", "description", "status", "bean"}

func FindPublicProfile(account_id int64) (*model.User, error) {
	user := &model.User{}
	err := persistence.DB.Select(PublicProfileFields).Where("account_id=?", account_id).Find(&user).Error
	if err != nil {
		return nil, err
	} else {
		return user, nil
	}
}
