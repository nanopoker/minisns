package rpcclient

import (
	code "github.com/nanopoker/minisns/apps/retcode"
	"github.com/nanopoker/minisns/libs/logger"
	"github.com/nanopoker/minisns/libs/utility"
	"github.com/nanopoker/minisns/proto"
	"strconv"
)

func Login(username, password string) (code.Code, string) {
	rsp, _ := client.Login(ctx, &proto.LoginRequest{Username: username, Password: password})
	return code.Code{CodeNum: rsp.Code, CodeMsg: rsp.Msg}, rsp.Identity
}

func Logout(identity string) code.Code {
	rsp, _ := client.Logout(ctx, &proto.LogoutRequest{Identity: identity})
	return code.Code{CodeNum: rsp.Code, CodeMsg: rsp.Msg}
}

func Register(username, password, email, nickname string) code.Code {
	salt := utility.GenerateSalt()
	encrypted_password := utility.Crypto(password, salt)
	userid, err := strconv.Atoi(utility.GenerateUserid())
	if err != nil {
		logger.Info("err is:", err.Error())
		return code.CodeFailedGenUserid
	}
	rsp, _ := client.Register(ctx, &proto.RegisterRequest{Userid: uint32(userid), Username: username, Password: encrypted_password,
		Nickname: nickname, Email: email, Salt: salt})
	return code.Code{CodeNum: rsp.Code, CodeMsg: rsp.Msg}
}

func EditUser(nickname, password, email, identity string) (code.Code, uint32, string, string, string) {
	rsp, _ := client.EditUser(ctx, &proto.EditRequest{Nickname: nickname, Password: password, Email: email, Identity: identity})
	return code.Code{CodeNum: rsp.Code, CodeMsg: rsp.Msg}, rsp.Userid, rsp.Nickname, rsp.Email, rsp.Username
}

func Follow(identity string, followee uint32) code.Code {
	rsp, _ := client.Follow(ctx, &proto.FollowRequest{Identity: identity, Followee: followee})
	return code.Code{CodeNum: rsp.Code, CodeMsg: rsp.Msg}
}

func Followlist(pageno, count int, identity string) (code.Code, []*proto.SingleUser) {
	rsp, _ := client.Followlist(ctx, &proto.FollowlistRequest{Pageno: uint32(pageno), Count: uint32(count), Identity: identity})
	if rsp.Code != code.CodeSucc.CodeNum {
		return code.Code{CodeNum: rsp.Code, CodeMsg: rsp.Msg}, nil
	}
	return code.Code{CodeNum: rsp.Code, CodeMsg: rsp.Msg}, rsp.Users
}

func AuthIdentity(identity string) code.Code {
	rsp, _ := client.AuthIdentity(ctx, &proto.AuthIdentityRequest{Identity: identity})
	return code.Code{CodeNum: rsp.Code, CodeMsg: rsp.Msg}
}
