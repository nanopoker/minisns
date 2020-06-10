package tcpserver

import (
	"context"
	"github.com/nanopoker/minisns/apps/model"
	code "github.com/nanopoker/minisns/apps/retcode"
	"github.com/nanopoker/minisns/libs/logger"
	"github.com/nanopoker/minisns/libs/utility"
	"github.com/nanopoker/minisns/proto"
)

type Server struct {
	proto.UnimplementedUserServiceServer
}

func (server *Server) Login(ctx context.Context, in *proto.LoginRequest) (out *proto.LoginResponse, err error) {
	user, err := model.GetUserByUsername(in.Username)
	if err != nil {
		logger.Error("get user info failed", err.Error())
		out = &proto.LoginResponse{Code: code.CodeGetUserFailed.CodeNum,
			Msg: code.CodeGetUserFailed.CodeMsg, Identity: "", Userid: 0, Nickname: "", Email: "", Username: ""}
		return out, nil
	}
	encrypted_password := utility.Crypto(in.Password, user.Salt)
	user, err = model.LoginAuth(in.Username, encrypted_password)
	if err != nil {
		out = &proto.LoginResponse{Code: code.CodeInvalidPasswd.CodeNum,
			Msg: code.CodeInvalidPasswd.CodeMsg, Identity: "", Userid: 0, Nickname: "", Email: "", Username: ""}
		return out, nil
	}
	identity := utility.GenerateIdentity()
	logger.Info("identity is", identity)
	err = model.SetIdentityCache(ctx, user.Userid, identity)
	if err != nil {
		out = &proto.LoginResponse{Code: code.CodeSetIdentityCacheErr.CodeNum,
			Msg: code.CodeSetIdentityCacheErr.CodeMsg, Identity: "",
			Userid: user.Userid, Nickname: user.Nickname, Email: user.Email, Username: user.Username}
		return out, nil
	}
	out = &proto.LoginResponse{Code: code.CodeSucc.CodeNum,
		Msg: code.CodeSucc.CodeMsg, Identity: identity,
		Userid: user.Userid, Nickname: user.Nickname, Email: user.Email, Username: user.Username}
	return out, nil
}

func (server *Server) Logout(ctx context.Context, in *proto.LogoutRequest) (out *proto.LogoutResponse, err error) {
	err = model.DelIdentityCache(ctx, in.Identity)
	if err != nil {
		logger.Error("logout DelIdentityCache error,", err.Error())
		out = &proto.LogoutResponse{Code: code.CodeDelIdentityCacheErr.CodeNum,
			Msg: code.CodeDelIdentityCacheErr.CodeMsg}
	} else {
		out = &proto.LogoutResponse{Code: code.CodeSucc.CodeNum,
			Msg: code.CodeSucc.CodeMsg}
	}
	return out, nil
}

func (server *Server) Register(ctx context.Context, in *proto.RegisterRequest) (out *proto.RegisterResponse, err error) {
	user, err := model.GetUserByUsername(in.Username)
	if err != nil {
		logger.Error("register service getuser error", err.Error())
		out = &proto.RegisterResponse{Code: code.CodeGetUserFailed.CodeNum,
			Msg: code.CodeGetUserFailed.CodeMsg}
		return out, nil
	}
	if user.Username != "" {
		out = &proto.RegisterResponse{Code: code.CodeUsernameExists.CodeNum,
			Msg: code.CodeUsernameExists.CodeMsg}
		return out, nil
	}

	user, err = model.RegisterUser(in.Userid, in.Username, in.Password, in.Nickname, in.Email, in.Salt)
	if err != nil {
		logger.Error("register user failed error,", err.Error())
		out = &proto.RegisterResponse{Code: code.CodeRegisterUserFailed.CodeNum,
			Msg: code.CodeRegisterUserFailed.CodeMsg}
	} else {
		out = &proto.RegisterResponse{Code: code.CodeSucc.CodeNum,
			Msg: code.CodeSucc.CodeMsg}
	}
	return out, nil
}

func (server *Server) EditUser(ctx context.Context, in *proto.EditRequest) (out *proto.EditResponse, err error) {
	userid, err := model.GetIdentityCache(ctx, in.Identity)
	if err != nil {
		logger.Error("get identity cache failed", err.Error())
		out = &proto.EditResponse{Code: code.CodeGetIdentityCacheErr.CodeNum,
			Msg: code.CodeGetIdentityCacheErr.CodeMsg, Userid: 0, Nickname: "", Email: "", Username: ""}
		return out, nil
	}
	user, err := model.GetUser(userid)
	if err != nil {
		logger.Error("get user info failed", err.Error())
		out = &proto.EditResponse{Code: code.CodeGetUserFailed.CodeNum,
			Msg: code.CodeGetUserFailed.CodeMsg, Userid: userid, Nickname: "", Email: "", Username: ""}
		return out, nil
	}
	encrypted_password := utility.Crypto(in.Password, user.Salt)
	err = model.EditUser(userid, in.Nickname, in.Email, encrypted_password)
	if err != nil {
		out = &proto.EditResponse{Code: code.CodeEditUserFailed.CodeNum,
			Msg: code.CodeEditUserFailed.CodeMsg, Userid: userid, Nickname: user.Nickname, Email: user.Email, Username: user.Username}
	} else {
		out = &proto.EditResponse{Code: code.CodeSucc.CodeNum,
			Msg: code.CodeSucc.CodeMsg, Userid: userid, Nickname: in.Nickname, Email: in.Email, Username: user.Username}
	}
	return out, nil
}

func (server *Server) Follow(ctx context.Context, in *proto.FollowRequest) (out *proto.FollowResponse, err error) {
	userid, err := model.GetIdentityCache(ctx, in.Identity)
	if err != nil {
		logger.Error("get identity cache failed", err.Error())
		out = &proto.FollowResponse{Code: code.CodeGetIdentityCacheErr.CodeNum,
			Msg: code.CodeGetIdentityCacheErr.CodeMsg}
		return out, nil
	}
	user, err := model.GetUser(userid)
	if err != nil {
		logger.Error("get user info failed", err.Error())
		out = &proto.FollowResponse{Code: code.CodeGetUserFailed.CodeNum,
			Msg: code.CodeGetUserFailed.CodeMsg}
		return out, nil
	}
	followee, err := model.GetUser(in.Followee)
	if err != nil {
		logger.Error("get user info failed", err.Error())
		out = &proto.FollowResponse{Code: code.CodeGetUserFailed.CodeNum,
			Msg: code.CodeGetUserFailed.CodeMsg}
		return out, nil
	}
	if followee.Username == "" {
		out = &proto.FollowResponse{Code: code.CodeUsernameNotExists.CodeNum,
			Msg: code.CodeUsernameNotExists.CodeMsg}
		return out, nil
	}
	if followee.Username == user.Username {
		out = &proto.FollowResponse{Code: code.CodeUserFollowSelf.CodeNum,
			Msg: code.CodeUserFollowSelf.CodeMsg}
		return out, nil
	}

	follow, err := model.GetFollow(user.Id, followee.Id)
	if err != nil {
		out = &proto.FollowResponse{Code: code.CodeGetFollowInfoFailed.CodeNum,
			Msg: code.CodeGetFollowInfoFailed.CodeMsg}
		return out, nil
	}
	if follow.Id != 0 {
		out = &proto.FollowResponse{Code: code.CodeFollowExists.CodeNum,
			Msg: code.CodeFollowExists.CodeMsg}
		return out, nil
	}
	err = model.FollowUser(user.Userid, followee.Userid)
	if err == nil {
		out = &proto.FollowResponse{Code: code.CodeSucc.CodeNum,
			Msg: code.CodeSucc.CodeMsg}
	} else {
		out = &proto.FollowResponse{Code: code.CodeInternalErr.CodeNum,
			Msg: code.CodeInternalErr.CodeMsg}
	}
	return out, nil
}

func (server *Server) Followlist(ctx context.Context, in *proto.FollowlistRequest) (out *proto.FollowlistResponse, err error) {
	userid, err := model.GetIdentityCache(ctx, in.Identity)
	if err != nil {
		logger.Error("get identity cache failed", err.Error())
		out = &proto.FollowlistResponse{Code: code.CodeGetIdentityCacheErr.CodeNum,
			Msg: code.CodeGetIdentityCacheErr.CodeMsg}
		return out, nil
	}
	users, err := model.Followlist(in.Pageno, in.Count, userid)
	if err == nil {
		out = &proto.FollowlistResponse{Code: code.CodeSucc.CodeNum,
			Msg: code.CodeSucc.CodeMsg, Users: users}
	} else {
		out = &proto.FollowlistResponse{Code: code.CodeGetFollowlistFailed.CodeNum,
			Msg: code.CodeGetFollowlistFailed.CodeMsg, Users: users}
	}
	return out, nil
}

func (server *Server) AuthIdentity(ctx context.Context, in *proto.AuthIdentityRequest) (out *proto.AuthIdentityResponse, err error) {
	userid, err := model.GetIdentityCache(ctx, in.Identity)
	if err != nil {
		logger.Error("get identity cache failed", err.Error())
		out = &proto.AuthIdentityResponse{Code: code.CodeGetIdentityCacheErr.CodeNum,
			Msg: code.CodeGetIdentityCacheErr.CodeMsg}
		return out, nil
	}
	if userid != 0 {
		out = &proto.AuthIdentityResponse{Code: code.CodeSucc.CodeNum,
			Msg: code.CodeSucc.CodeMsg}
		return out, nil
	} else {
		out = &proto.AuthIdentityResponse{Code: code.CodeIdentityNotFound.CodeNum,
			Msg: code.CodeIdentityNotFound.CodeMsg}
		return out, nil
	}
}
