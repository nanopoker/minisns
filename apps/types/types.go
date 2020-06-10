package types

// Binding from JSON
type LoginParam struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Binding from JSON
type RegisterParam struct {
	Username string `json:"username" binding:"required"`
	Nickname string `json:"nickname" binding:"required"`
	Email    string `json:"email"    binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Binding from JSON
type EditParam struct {
	Nickname string `json:"nickname" binding:"required"`
	Email    string `json:"email"    binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Binding from JSON
type FollowParam struct {
	Followee uint32 `json:"followee" binding:"required"`
}

// Binding from JSON
type FollowlistParam struct {
	Pageno int `json:"pageno" form:"pageno"`
	Count  int `json:"count"  form:"count"`
}

type FollowTab struct {
	Id         uint32 "type:bigint;primary key"
	Follower   uint32 "type:bigint"
	Followee   uint32 "type:bigint"
	Createtime string "type:datetime"
}

type UserTab struct {
	Id         uint32 "type:bigint;primary key"
	Userid     uint32 "type:bigint"
	Username   string "type:varchar(100);unique;not null"
	Nickname   string "type:varchar(100)"
	Salt       string "type:varchar(10);not null"
	Password   string "type:varchar(64);not null"
	Email      string "type:varchar(100);not null"
	Createtime string "type:datetime"
	Updatetime string "type:datetime"
}
