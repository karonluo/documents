package services

import (
	"encoding/json"
	"fmt"
	"go-demo/pkg/biz"
	"go-demo/pkg/entities"

	"github.com/kataras/iris/v12"
)

// 由于 go 语言没有重载，因此每个方法都要专门命名，本次主要涉及到 WebService 因此只实现 Web Service 方式，其他的服务模式，可以以后再加，如：API 模式和 gRPC 模式
func WSBuyBook(ctx iris.Context) {

	var book entities.Book
	fmt.Println("Web 服务接口被调用")
	// 此处应该访问业务层的 BuyBook 方法
	book = biz.BuyBook(ctx.FormValue("id"))
	res, _ := json.Marshal(book)
	ctx.HTML(string(res))
}

// API 模式的方法命名
func APIBuyBook(id string) entities.Book {

	book := biz.BuyBook(id)
	return book
}

// gRPC 模式的方法命名
func GRPCBuyBook() {

}
