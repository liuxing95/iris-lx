package main

import (
	stdContext "context"
	"flag"
	"fmt"
	"time"

	"iris-lx/config"
	"iris-lx/web/routers"

	"github.com/kataras/iris/v12"
)

func main() {
	flag.Parse()
	level := "info"

	app := newApp(level)
	iris.RegisterOnInterrupt(func() {
		timeout := 5 * time.Second
		ctx, cancel := stdContext.WithTimeout(stdContext.Background(), timeout)
		defer cancel()
		//关闭所有主机
		app.Shutdown(ctx)
	})
	err := app.Run(iris.Addr(fmt.Sprintf(":%d", config.WEB_PORT)), iris.WithoutInterruptHandler, iris.WithPostMaxMemory(int64(config.MAX_FILE_SIZE)))
	if err != nil {
		panic(err)
	}
}

func newApp(logLevel string) *iris.Application {
	app := iris.Default()
	app.Configure(
		iris.WithRemoteAddrHeader("X-Forwarded-For"),
		iris.WithRemoteAddrHeader("X-Real-Ip"))
	app.Logger().SetLevel(logLevel)
	routers.InitRouter(app)

	return app
}
