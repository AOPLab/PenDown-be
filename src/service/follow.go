package service

import (
	"errors"

	"github.com/AOPLab/PenDown-be/src/model"
	"github.com/AOPLab/PenDown-be/src/persistence"
)

type Follow_Member_Detail struct {
	Account_id int64  `json:"account_id"`
	Username   string `json:"username"`
	Full_name  string `json:"full_name"`
}

func AddFollow(follower_id int64, followee_id int64) error {

	follow := &model.Follow{
		Follower_id: follower_id,
		Followee_id: followee_id,
	}

	db_err := persistence.DB.Model(&model.Follow{}).Create(&follow).Error
	if db_err != nil {
		return db_err
	} else {
		return nil
	}
}

func GetFollow(follower_id int64, followee_id int64) (bool, error) {
	var follow model.Follow
	err := persistence.DB.Where("Follower_id = ?", follower_id).Where("Followee_id = ?", followee_id).First(&follow).Error
	if err != nil {
		if err.Error() == "record not found" {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
}

func GetFollowers(followee_id int64) ([]*Follow_Member_Detail, error) {
	var follows []model.Follow
	var follow_members []*Follow_Member_Detail
	err := persistence.DB.Where("Followee_id = ?", followee_id).Find(&follows).Error
	if err != nil {
		return nil, err
	} else {
		for _, v := range follows {
			follower_id := v.Follower_id
			follower, follower_err := findUserByAccountID(follower_id)
			if follower_err == nil {
				var follow_member Follow_Member_Detail
				follow_member.Account_id = follower_id
				follow_member.Full_name = follower.Full_name
				follow_member.Username = follower.Username
				follow_members = append(follow_members, &follow_member)
			}
		}
	}
	return follow_members, nil
}

func GetFollowing(follower_id int64) ([]*Follow_Member_Detail, error) {
	var follows []model.Follow
	var follow_members []*Follow_Member_Detail
	err := persistence.DB.Where("Follower_id = ?", follower_id).Find(&follows).Error
	if err != nil {
		return nil, err
	} else {
		for _, v := range follows {
			followee_id := v.Followee_id
			followee, followee_err := findUserByAccountID(followee_id)
			if followee_err == nil {
				var follow_member Follow_Member_Detail
				follow_member.Account_id = followee_id
				follow_member.Full_name = followee.Full_name
				follow_member.Username = followee.Username
				follow_members = append(follow_members, &follow_member)
			}
		}
	}
	return follow_members, nil
}

func DeleteFollow(follower_id int64, followee_id int64) error {
	follow, follow_err := GetFollow(follower_id, followee_id)
	if follow_err != nil {
		return follow_err
	}
	if !follow {
		return errors.New("follow not exists")
	}

	db_err := persistence.DB.Unscoped().Where("Follower_id = ? AND Followee_id = ?", follower_id, followee_id).Delete(&model.Follow{}).Error
	if db_err != nil {
		return db_err
	} else {
		return nil
	}
}
