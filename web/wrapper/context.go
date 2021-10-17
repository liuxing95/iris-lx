package wrapper

import (
	"sync"

	"github.com/kataras/iris/v12"
)

type Context struct {
	iris.Context
}

var contextPool = sync.Pool{New: func() interface{} {
	return &Context{}
}}

func CasbinRequired(ctx *Context) bool {
	// path := ctx.Path()
	// role := ctx.UserJwt.Role
	// method := ctx.Method()
	// res := micro_middleware.H2PermissionCheck(int(role), method, path)
	// if !res.OK {
	// 	SendApiResponse(ctx, &ApiResponse{OK: false, Msg: "无访问权限"})
	// }
	return true
}

func acquire(original iris.Context, login bool) *Context {
	ctx := contextPool.Get().(*Context)
	ctx.Context = original // set the context to the original one in order to have access to iris's implementation.
	// if login {
	// 	// ctx.UserJwt = GetJwtInfo(original, false, false) // reset the jwt token
	// 	// if ctx.UserJwt == nil {                          // 验证失败，终止请求
	// 	// 	ctx.StopExecution()
	// 	// }
	// }
	return ctx
}

func release(ctx *Context) {
	contextPool.Put(ctx)
}

// Handler will convert our handler of func(*Context) to an iris Handler,
// in order to be compatible with the HTTP API.
func Handler(h func(*Context)) iris.Handler {
	return func(original iris.Context) {
		ctx := acquire(original, true)
		if !ctx.IsStopped() && CasbinRequired(ctx) { // 请求被终止
			h(ctx)
		}
		release(ctx)
	}
}
