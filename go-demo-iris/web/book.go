package web

import (
	"fmt"
	"go-demo-iris/pkg/biz"
	"go-demo-iris/pkg/services"

	"github.com/kataras/iris/v12"
)

// 创建 book.go 在 go-demo/web/ 下，专门用于开发 book 相关对象的 WEB 视图
func BookIndex(ctx iris.Context) {
	ctx.HTML(ctx.FormValue("id"))
}

// 获取书籍信息并绑定到页面
func GetBookInformation(ctx iris.Context) {
	// 这里可以不用调用WEB接口服务层，需要绑定页面数据，属于Views层，只有当调用接口服务返回结果时才使用。
	// 如果不含有业务逻辑，则可以直接调用 DAO 数据访问层 填充 Book 实体，如果有业务逻辑，包括权限之类的，建议还是将其在 biz 层进行实现，这里按照标准来进行编写
	book := biz.GetBookInformation(ctx.FormValue("id"))
	ctx.ViewData("bookname", book.Name)
	ctx.ViewData("bookauthor", book.AuthorDisplayName)
	ctx.View("bookinfor.html")

}
func BuyBook(ctx iris.Context) {
	// 调用 services 服务层
	fmt.Println("Web 页面被调用")
	services.WSBuyBook(ctx)
}
