package consts

import "github.com/RanFeng/ierror"

const (
	DBError = -200001 - iota
	RedisError
)

const (
	Service = iota + 10000
	Params
	AuthorizationFailed
	UserAlreadyExist
	UserIsNotExist
	TokenInvalid
	UserNameExist
	PasswordWrong
	PhoneNumInvalid
	NoLogin
)

var (
	TokenIsInvalid = ierror.NewIError(TokenInvalid, "Token is invalid")
	ErrNoLogin     = ierror.NewIError(NoLogin, "未登录")
)
