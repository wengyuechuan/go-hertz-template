package router

import (
	"github.com/cloudwego/hertz/pkg/route"
)

func UserRegister(r route.IRouter) {
	_user := r.Group("/user")
	{
		_user.GET("/login")
	}
}
