package entities

type Book struct {
	Name              string // go 语言 大写字母开始的都是公共变量，而不用标记 public
	Id                string
	name              string // go 语言 小写字母开始的都是私有变量，而不用标记 private
	id                string
	AuthorId          string
	AuthorDisplayName string
}

func (self *Book) GetId() string { //go 语言必须使用这种{}方式，否则编译运行报错
	return self.id
}
