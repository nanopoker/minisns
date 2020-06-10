package model

import (
	"github.com/jinzhu/gorm"
	"github.com/nanopoker/minisns/agent/db"
	"github.com/nanopoker/minisns/apps/types"
	"github.com/nanopoker/minisns/proto"
	"time"
)

func GetFollow(follower, followee uint32) (*types.FollowTab, error) {
	var follow types.FollowTab
	err := db.Db.Table("follow").Where("follower = ?  and followee = ? ", follower, followee).First(&follow).Error
	if gorm.IsRecordNotFoundError(err) {
		err = nil
	}
	return &follow, err
}

func FollowUser(follower, followee uint32) error {
	follow := types.FollowTab{
		Follower:   follower,
		Followee:   followee,
		Createtime: time.Now().Format("2006-01-02 15:04:05"),
	}
	err := db.Db.Table("follow").Create(&follow).Error
	return err
}

func Followlist(pageno, count, userid uint32) ([]*proto.SingleUser, error) {
	var follows []uint32
	var users []*proto.SingleUser
	offset := (pageno - 1) * count
	err := db.Db.Table("follow").Where("follower = ?", userid).Limit(count).Offset(offset).Order("createtime desc").Pluck("followee", &follows).Error
	if gorm.IsRecordNotFoundError(err) {
		return users, nil
	}

	err = db.Db.Table("user").Select("userid,username").Where("userid in (?) ", follows).Scan(&users).Error
	if gorm.IsRecordNotFoundError(err) {
		return users, nil
	}
	return users, err
}
