package model

import (
	"context"
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/nanopoker/minisns/agent/cache"
	"github.com/nanopoker/minisns/agent/db"
	"github.com/nanopoker/minisns/apps/types"
	"github.com/nanopoker/minisns/config"
	"strconv"
	"time"
)

const identityKeyPrefix = "identity_"

var conn *cache.RConn

func GetUser(userid uint32) (*types.UserTab, error) {
	var user types.UserTab
	err := db.Db.Table("user").Where("userid = ? ", userid).First(&user).Error
	if gorm.IsRecordNotFoundError(err) {
		err = nil
	}
	return &user, err
}

func GetUserByUsername(username string) (*types.UserTab, error) {
	var user types.UserTab
	err := db.Db.Table("user").Where("username = ? ", username).First(&user).Error
	if gorm.IsRecordNotFoundError(err) {
		err = nil
	}
	return &user, err
}

func LoginAuth(username, password string) (*types.UserTab, error) {
	var user types.UserTab
	err := db.Db.Table("user").Where("username = ? and password = ?", username, password).First(&user).Error
	return &user, err
}

func EditUser(userid uint32, nickname, email, password string) (err error) {
	err = db.Db.Table("user").Where("userid = ?", userid).
		Updates(map[string]interface{}{"nickname": nickname,
			"email": email, "password": password, "mtime": uint32(time.Now().Unix())}).Error
	return err
}

func RegisterUser(userid uint32, username, password, nickname, email, salt string) (*types.UserTab, error) {
	user := types.UserTab{
		Userid:     userid,
		Username:   username,
		Nickname:   nickname,
		Password:   password,
		Email:      email,
		Salt:       salt,
		Createtime: time.Now().Format("2006-01-02 15:04:05"),
		Updatetime: time.Now().Format("2006-01-02 15:04:05"),
	}
	err := db.Db.Table("user").Create(&user).Error
	return &user, err
}

func GetIdentityCache(ctx context.Context, identity string) (userid uint32, err error) {
	conn = cache.GetRConn()
	if conn == nil {
		return userid, errors.New("no redis conn")
	}
	redisKey := identityKeyPrefix + identity
	val, err := conn.Get(ctx, redisKey)
	if err != nil {
		return userid, err
	}
	useridstring := string(val)
	useridint, err := strconv.Atoi(useridstring)
	if err != nil {
		return userid, nil
	}
	return uint32(useridint), err
}

func SetIdentityCache(ctx context.Context, userid uint32, identity string) error {
	conn = cache.GetRConn()
	if conn == nil {
		return errors.New("no redis conn")
	}
	redisKey := identityKeyPrefix + identity
	useridstring := strconv.Itoa(int(userid))
	useridbyte := []byte(useridstring)
	return conn.SetEX(ctx, redisKey, config.REDIS_KEY_DURATION, useridbyte)
}

func DelIdentityCache(ctx context.Context, identity string) error {
	conn = cache.GetRConn()
	if conn == nil {
		return errors.New("no redis conn")
	}
	redisKey := identityKeyPrefix + identity
	return conn.Delete(ctx, redisKey)
}
