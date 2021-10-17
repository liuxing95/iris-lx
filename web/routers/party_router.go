package routers

import (
	"iris-lx/web/controllers"
	"iris-lx/web/wrapper"
)

// 配置所以 路由入口
func InitUserParty() {
	authParty := v1Router.Party("/user")
	authParty.Handle("GET", "/login", wrapper.Handler(controllers.UserController{}.Login))
}
