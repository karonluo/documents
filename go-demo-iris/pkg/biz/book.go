package biz

import (
	"fmt"
	"go-demo/pkg/entities"

	"github.com/google/uuid"
)

func BuyBook(id string) entities.Book {
	// 此处如果需要实例化将数据填充到 Book 对象则需要访问 数据访问层进行数据填充
	// TODO: 调用 DAO。
	fmt.Println("业务层，业务逻辑被调用")
	book := entities.Book{
		Id:                id,
		Name:              "《Go Language》",
		AuthorDisplayName: "Dr. Karon.Luo",
		AuthorId:          uuid.New().String(),
	}
	return book

}

func GetBookInformation(id string) entities.Book {
	fmt.Println("业务层，业务逻辑被调用")
	book := entities.Book{
		Id:                id,
		Name:              "《Go Language》",
		AuthorDisplayName: "Dr. Karon.Luo",
		AuthorId:          uuid.New().String(),
	}
	return book
}
