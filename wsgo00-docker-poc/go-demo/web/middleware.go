package web

// 中间件
import (
	"fmt"

	"github.com/kataras/iris/v12"
)

func UseBefore(ctx iris.Context) {
	fmt.Println("中间件被调用，可以做为拦截器，对用户进行验证，权限分配等功能")
	ctx.Next() // 交由被调用的 Web 页面或服务进行后续处理
}
