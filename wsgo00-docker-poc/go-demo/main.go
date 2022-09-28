package main

import (
	"encoding/json"
	"fmt"
	"go-demo/conf"
	"go-demo/pkg/entities"
	"go-demo/web"

	"github.com/google/uuid"
	"github.com/kataras/iris/v12"
)

// go 语言程序入口
func main() {
	fmt.Println("Hello,World!!!")
	// 验证刚才建立的实体对象，调用。
	var book entities.Book
	book.Id = uuid.New().String()
	book.Name = "《C/C++ Language 100 days》"
	book.AuthorId = uuid.New().String()
	book.AuthorDisplayName = "Dr. Kaorn.Luo"
	res, _ := json.Marshal(book)
	fmt.Println(string(res))
	WebServerStart()

}

func WebServerStart() {
	app := iris.Default()
	app.Use(web.UseBefore)                              // 使用 Iris 全局中间件，可以做为拦截器，对用户进行验证，权限分配等功能
	app.RegisterView(iris.HTML("./web/views", ".html")) // 注册 HTML 视图模板地址
	app.Get("/book", web.BookIndex)
	app.Get("/book/buy", web.BuyBook)
	app.Get("/book/info", web.GetBookInformation)

	webconf := conf.LoadConfiguration("./conf/Web.Config")
	app.Listen(webconf.Port)

}
