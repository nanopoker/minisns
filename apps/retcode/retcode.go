package retcode

type Code struct {
	CodeNum uint32
	CodeMsg string
}

var (
	// CodeSucc          succ code
	CodeSucc = Code{CodeNum: 0, CodeMsg: "succ"}

	// tcp 1000 ~ 2000
	CodeTCPFailedGetUserInfo = Code{CodeNum: 1101, CodeMsg: "tcp server: failed to get userinfo"}
	// CodeTCPPasswdErr password error
	CodeTCPPasswdErr = Code{CodeNum: 1102, CodeMsg: "tcp server: wrong passwd"}
	// CodeTCPInvalidIdentity invalid identity
	CodeTCPInvalidIdentity = Code{CodeNum: 1200, CodeMsg: "tcp server: invalid identity format"}
	// CodeTCPIdentityExpired identity expired
	CodeTCPIdentityExpired = Code{CodeNum: 1201, CodeMsg: "tcp server: login identity expired"}
	// CodeTCPUserInfoNotMatch identity info not match userinfo
	CodeTCPUserInfoNotMatch = Code{CodeNum: 1202, CodeMsg: "tcp server: login identity cache info not match"}
	// CodeTCPFailedUpdateUserInfo update userinfo failed
	CodeTCPFailedUpdateUserInfo = Code{CodeNum: 1301, CodeMsg: "tcp server: failed to update userinfo"}
	// CodeTCPInternelErr internel error
	CodeTCPInternelErr = Code{CodeNum: 1401, CodeMsg: "tcp server: internel error"}
	// CodeSetIdentityCacheErr
	CodeSetIdentityCacheErr = Code{CodeNum: 1402, CodeMsg: "set identity cache failed"}
	// CodeDelIdentityCacheErr
	CodeDelIdentityCacheErr = Code{CodeNum: 1403, CodeMsg: "del identity cache failed"}
	// CodeGetIdentityCacheErr
	CodeGetIdentityCacheErr = Code{CodeNum: 1404, CodeMsg: "get identity cache failed"}
	// CodeFailedGenUserid  err
	CodeFailedGenUserid = Code{CodeNum: 1405, CodeMsg: "failed to generate userid"}
	// CodeGetUserFailed  err
	CodeGetUserFailed = Code{CodeNum: 1406, CodeMsg: "failed to get user info"}
	// CodeRegisterUserFailed  err
	CodeRegisterUserFailed = Code{CodeNum: 1407, CodeMsg: "failed to get user info"}
	// CodeEditUserFailed  err
	CodeEditUserFailed = Code{CodeNum: 1408, CodeMsg: "failed to modify user"}
	// CodeGetFollowInfoFailed  err
	CodeGetFollowInfoFailed = Code{CodeNum: 1409, CodeMsg: "failed to get follow info"}
	// CodeGetFollowlistFailed  err
	CodeGetFollowlistFailed = Code{CodeNum: 1410, CodeMsg: "failed to get followlist"}

	// HTTP 2000 ~ 3000
	// CodeInternalErr   internel err
	CodeInternalErr = Code{CodeNum: 2101, CodeMsg: "please try again!"}
	// CodeIdentityNotFound missing identity
	CodeIdentityNotFound = Code{CodeNum: 2102, CodeMsg: "identity not found"}
	// CodeInvalidIdentity  identity format is invalid
	CodeInvalidIdentity = Code{CodeNum: 2103, CodeMsg: "param error: identity not found"}
	// CodeErrBackend    failed to comm with backend server
	CodeErrBackend = Code{CodeNum: 2201, CodeMsg: "Error found!please try again!"}
	// CodeInvalidPasswd passwd format isn't right
	CodeInvalidPasswd = Code{CodeNum: 2301, CodeMsg: "username/passwd error!"}
	// CodeUsernameExists get error
	CodeUsernameExists = Code{CodeNum: 2401, CodeMsg: "Username already been registered"}
	// CodeUsernameNotExists get error
	CodeUsernameNotExists = Code{CodeNum: 2402, CodeMsg: "Username not Exists"}
	// CodeUserFollowSelf get error
	CodeUserFollowSelf = Code{CodeNum: 2403, CodeMsg: "You cannot follow yourself"}
	// CodeFollowExists get error
	CodeFollowExists = Code{CodeNum: 2404, CodeMsg: "You have already followed this person"}
)
