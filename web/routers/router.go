package routers

import (
	"github.com/kataras/iris/v12"
)

var v1Router iris.Party

func InitRouter(app *iris.Application) {
	// Group 分组
	v1Router = app.Party("/api")
	InitUserParty()
}
