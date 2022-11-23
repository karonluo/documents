#GO语言 #T200 #数据库  #WEB/WEB后端  #DEV #关系型数据库 
# Go语言下的 ORM 框架 gorm 的介绍和使用

## 概述

前言:[gorm](https://link.zhihu.com/?target=https%3A//github.com/jinzhu/gorm)是Golang语言中一款性能极好的ORM库，对开发人员相对是比较友好的。当然还有另外一个[xorm](https://link.zhihu.com/?target=https%3A//github.com/go-xorm/xorm)库也是比较出名的，感兴趣的也可以看看这个库，接下来主要介绍下`gorm`库的一些基本使用。

完整教程：[GORM 指南 | GORM - The fantastic ORM library for Golang, aims to be developer friendly.](https://gorm.io/zh_CN/docs/index.html)

## GORM介绍和快速入门

### 功能概览

-   全功能ORM(无限接近)
-   关联(Has One, Has Many, Belongs To, Many To Many, 多态)
-   钩子函数Hook(在创建/保存/更新/删除/查找之前或之后)
-   预加载
-   事务
-   复合主键
-   SQL 生成器
-   数据库自动迁移
-   自定义日志
-   可扩展性, 可基于 GORM 回调编写插件
-   所有功能都被测试覆盖
-   开发者友好

### 安装

我们都知道，在golang中需要使用一些驱动包来对指定数据库进行操作，比如MySQL需要使用`github.com/go-sql-driver/mysql`库，而Sqlite需要使用`github.com/mattn/go-sqlite3`库来支持，不过好在gorm框架中对各个驱动包进行了简单包装，可以让我们在写程序时可以更方便的管理驱动库.

支持的数据库以及导入路径如下:

-   mysql: [http://github.com/jinzhu/gorm/dialects/mysql](https://link.zhihu.com/?target=http%3A//github.com/jinzhu/gorm/dialects/mysql)
-   postgres: [http://github.com/jinzhu/gorm/dialects/postgres](https://link.zhihu.com/?target=http%3A//github.com/jinzhu/gorm/dialects/postgres)
-   sqlite: [http://github.com/jinzhu/gorm/dialects/sqlite](https://link.zhihu.com/?target=http%3A//github.com/jinzhu/gorm/dialects/sqlite)
-   sqlserver: [http://github.com/jinzhu/gorm/dialects/mssql](https://link.zhihu.com/?target=http%3A//github.com/jinzhu/gorm/dialects/mssql)

>`注意:`gorm框架只是简单封装了数据库的驱动包，在安装时仍需要下载原始的驱动包

```Shell
# 由于在项目中使用mysql比较多，这里使用mysql进行数据存储
$ go get -u github.com/jinzhu/gorm
$ go get -u github.com/go-sql-driver/mysql
```

### 快速入门
```Shell
# 使用docker快速创建一个本地可连接的mysql实例
$ docker run -itd -e MYSQL_ROOT_PASSWORD='bgbiao.top' --name go-orm-mysql  -p 13306:3306 mysql:5.6

# 登陆mysql并创建一个测试库
$ docker exec -it go-orm-mysql mysql -uroot -pbgbiao.top
....
mysql> create database test_api;
Query OK, 1 row affected (0.00 sec)

mysql> show databases;
+--------------------+
| Database           |
+--------------------+
| information_schema |
| mysql              |
| performance_schema |
| test_api           |
+--------------------+
4 rows in set (0.00 sec)

# 运行一个简单示例
$ cat gorm-mysql-example.go
```

~~~Go
/*
Sample Code File: gorm-mysql-example.go
*/
package main
import (
    "fmt"
    "time"

    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)
// 定义一个数据模型(user表)
// 列名是字段名的蛇形小写(PassWd->pass_word)
type User struct {
    Id          uint        `gorm:"AUTO_INCREMENT"`
    Name        string      `gorm:"size:50"`
    Age         int         `gorm:"size:3"`
    Birthday    *time.Time
    Email       string      `gorm:"type:varchar(50);unique_index"`
    PassWord    string      `gorm:"type:varchar(25)"`
}
var db *gorm.DB
func main() {
    db,err := gorm.Open("mysql","root:bgbiao.top@(127.0.0.1:13306)/test_api?charset=utf8&parseTime=True&loc=Local")
    if err != nil {
        fmt.Errorf("创建数据库连接失败:%v",err)
    }
    defer db.Close()
    // 自动迁移数据结构(table schema)
    // 注意:在gorm中，默认的表名都是结构体名称的复数形式，比如User结构体默认创建的表为users
    // db.SingularTable(true) 可以取消表名的复数形式，使得表名和结构体名称一致
    db.AutoMigrate(&User{})
    // 添加唯一索引
    db.Model(&User{}).AddUniqueIndex("name_email", "id", "name","email")
    // 插入记录
    db.Create(&User{Name:"bgbiao",Age:18,Email:"bgbiao@bgbiao.top"})
    db.Create(&User{Name:"xxb",Age:18,Email:"xxb@bgbiao.top"})
    var user User
    var users []User
    // 查看插入后的全部元素
    fmt.Printf("插入后元素:\n")
    db.Find(&users)
    fmt.Println(users)
    // 查询一条记录
    db.First(&user,"name = ?","bgbiao")
    fmt.Println("查看查询记录:",user)
    // 更新记录(基于查出来的数据进行更新)
    db.Model(&user).Update("name","biaoge")
    fmt.Println("更新后的记录:",user)
    // 删除记录
    db.Delete(&user)
    // 查看全部记录
    fmt.Println("查看全部记录:")
    db.Find(&users)
    fmt.Println(users)
}
~~~

~~~Shell
# 运行gorm实例
$ go run gorm-mysql-example.go
插入后元素:
[{1 bgbiao 18 <nil> bgbiao@bgbiao.top } {2 xxb 18 <nil> xxb@bgbiao.top }]
查看查询记录: {1 bgbiao 18 <nil> bgbiao@bgbiao.top }
更新后的记录: {1 biaoge 18 <nil> bgbiao@bgbiao.top }
查看全部记录:
[{2 xxb 18 <nil> xxb@bgbiao.top }]
~~~

### GORM常用的功能函数
#### 自动迁移
>`注意:` 使用自动迁移模式将保持表的更新，但是不会更新索引以及现有列的类型或删除未使用的列
~~~Go
// 同时迁移多个模型
db.AutoMigrate(&User{}, &Product{}, &Order{})
// 创建表时增加相关参数
// 比如修改表的字符类型CHARSET=utf8
db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{})
~~~

#### 检查表
```Go
// 检查模型是否存在
db.HasTable(&User{})
// 检查表是否存在
db.HasTable("users")
```

#### 增、删、改表结构
>`注意:`在实际企业的生产环境中，通常数据库级别的变更操作，都需要转换成 sql 交由 DBA 兄弟帮忙进行线上库表变更，因此不论使用自动迁移，还是手动创建表，都是在开发环境阶段的
~~~Go
// 使用模型创建
db.CreateTable(&User{})
// 增加参数创建
db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&User{})
// 删除表
db.DropTable(&User{})
db.DropTable("users")
// 模型和表名的混搭
db.DropTableIfExists(&User{}, "products")
// 修改列(修改字段类型)
db.Model(&User{}).ModifyColumn("description", "text")
// 删除列
db.Model(&User{}).DropColumn("description")
// 指定表名创建表
db.Table("deleted_users").CreateTable(&User{})
// 指定表名查询
var deleted_users []User
db.Table("deleted_users").Find(&deleted_users)
~~~
#### 索引和约束
```Go
// 添加外键
// 1st param : 外键字段
// 2nd param : 外键表(字段)
// 3rd param : ONDELETE
// 4th param : ONUPDATE
db.Model(&User{}).AddForeignKey("city_id", "cities(id)", "RESTRICT", "RESTRICT")
// 单个索引
db.Model(&User{}).AddIndex("idx_user_name", "name")
// 多字段索引
db.Model(&User{}).AddIndex("idx_user_name_age", "name", "age")
// 添加唯一索引(通常使用多个字段来唯一标识一条记录)
db.Model(&User{}).AddUniqueIndex("idx_user_name", "name")
db.Model(&User{}).AddUniqueIndex("idx_user_name_age", "name", "id","email")
// 删除索引
db.Model(&User{}).RemoveIndex("idx_user_name")
```
### GORM模型注意事项
#### gorm.Model 结构体
其实在`gorm`官方文档的示例中，会默认在模型的属性中增加一个[gorm.Model](https://link.zhihu.com/?target=https%3A//pkg.go.dev/github.com/jinzhu/gorm%3Ftab%3Ddoc%23Model)的属性，该属性的原始结构如下:
```Go
// 官方示例模型
type User struct {
  gorm.Model
  Name         string
	.....
}

// gorm.Model结构体
type Model struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
```
很明显，当我们在用户自定义的模型中增加`gorm.Model`时，会自动为我们的表增加`id,created_at,updated_at,deleted_at`四个字段。

同时需要注意的是，当我们的模型中有`CreatedAt`,`UpdatedAt`,`DeletedAt`属性并且类型为`time.Time`或者`*time.Time`类型时，当有数据操作时，会自动更新对应的时间。

所以，在定义模型时，可以根据实际的需求来决定是否引入`gorm.Model`结构

另外需要注意的是: 所有字段的零值, 比如`0, '', false 或者其它零值`都不会保存到数据库内，但会使用他们的默认值，因此一些非必须字段，可以使用`DEFAULT`的tag字段来声明列的默认值。

#### gorm模型支持的tags

>`注意:`gorm支持一些常见的tags来自定义模型字段的扩展信息

**结构体标记描述** column 列明(默认是字段的蛇形小写) type 数据类型 size 列长度 PRIMARY_KEY 声明主键 UNIQUE 声明唯一 DEFAULT 指定默认值 PRECISION 声明列精度 NOT NULL 将列指定为非 NULLAUTO_INCREMENT 声明自增列 INDEX 创建具有或不带名称的索引, 如果多个索引同名则创建复合索引 UNIQUE_INDEX 创建唯一索引 EMBEDDED 将结构设置为嵌入 EMBEDDED_PREFIX 设置嵌入结构的前缀-忽略此字段

>`注意:` gorm也支持一些关联的结构体标签，比如外键，关联外键，等操作，通常在复杂的企业环境中，建议在库表设计时将相关表都设计成孤立表，具体的关联逻辑由业务层去实现(可能增加了开发的成本，不过当业务发展比较复杂时，这样做无疑是方便后期做扩展和优化的)

### 详细的 CRUD 接口
#### 创建
**插入记录**
```Go
// 相当于insert into users(name,age,brithday) values("BGBiao",18,time.Now())
user := User{Name: "BGBiao", Age: 18, Birthday: time.Now()}
// 主键为空返回`true`
db.NewRecord(user)
db.Create(&user)
// 创建`user`后返回`false`
db.NewRecord(user)
```
**在Hooks中设置字段值**
>`注意`: 通常我们在设计模型时，有一些原始字段，希望在初始化模型后就拥有记录，此时可以使用hooks来插入初始记录

如果想再 `BeforeCreate` hook中修改字段的值，可以使用 `scope.SetColumn`:
```Go
func (user *User) BeforeCreate(scope *gorm.Scope) error {
  scope.SetColumn("ID", uuid.New())
  return nil
}
```
**创建扩展选项**
```Go
// 为Instert语句添加扩展SQL选项
// insert into produce(name,code) values("name","code") on conflict;
db.Set("gorm:insert_option", "ON CONFLICT").Create(&product)
```
#### 查询
**基本查询**

```Go
// 根据主键查询第一条记录
// SELECT * FROM users ORDER BY id LIMIT 1;
db.First(&user)

// 随机获取一条记录
// SELECT * FROM users LIMIT 1;
db.Take(&user)

// 根据主键查询最后一条记录
// SELECT * FROM users ORDER BY id DESC LIMIT 1;
db.Last(&user)

// 查询所有的记录
// SELECT * FROM users;
db.Find(&users)

// 查询指定的某条记录(仅当主键为整型时可用)
// SELECT * FROM users WHERE id = 10;
db.First(&user, 10)
```
**结构体方式查询**
```Go
// 结构体方式
// select * from users where name = 'bgbiao.top'
db.Where(&User{Name: "bgbiao.top", Age: 20}).First(&user)

// Map方式
// select * from users where name = 'bgbiao.top' and age = 20;
db.Where(map[string]interface{}{"name": "bgbiao.top", "age": 20}).Find(&users)

// 主键的切片
// select * from users where id in (20,21,22);
db.Where([]int64{20, 21, 22}).Find(&users)

```
**Where条件查询**
>`注意:` 使用了Where()方法，基本上就是基本的sql语法

```Go
// 使用条件获取一条记录 First()方法
db.Where("name = ?", "bgbiao.top").First(&user)

// 获取全部记录 Find()方法
db.Where("name = ?", "jinzhu").Find(&users)

// 不等于
db.Where("name <> ?", "jinzhu").Find(&users)

// IN
db.Where("name IN (?)", []string{"jinzhu", "bgbiao.top"}).Find(&users)

// LIKE
db.Where("name LIKE ?", "%jin%").Find(&users)

// AND
db.Where("name = ? AND age >= ?", "jinzhu", "22").Find(&users)

// Time
// select * from users where updated_at > '2020-03-06 00:00:00'
db.Where("updated_at > ?", lastWeek).Find(&users)

// BETWEEN
// select * from users where created_at between '2020-03-06 00:00:00' and '2020-03-14 00:00:00'
db.Where("created_at BETWEEN ? AND ?", lastWeek, today).Find(&users)
```

**Not条件**

```Go
// select * from users where name != 'bgbiao.top';
db.Not("name", "jinzhu").First(&user)

// Not In
// select * from users where name not in ("jinzhu","bgbiao.top");
db.Not("name", []string{"jinzhu", "bgbiao.top"}).Find(&users)

// 主键不在slice中的查询
// select * from users where id not in (1,2,3)
db.Not([]int64{1,2,3}).First(&user)

// select * from users;
db.Not([]int64{}).First(&user)

// 原生SQL
// select * from users where not(name = 'bgbiao.top');
db.Not("name = ?", "bgbiao.top").First(&user)

// struct方式查询
// select * from users where name != 'bgbiao.top'
db.Not(User{Name: "bgbiao.top"}).First(&user)
```

**Or条件**

```Go
// SELECT * FROM users WHERE role = 'admin' OR role = 'super_admin';
db.Where("role = ?", "admin").Or("role = ?", "super_admin").Find(&users)

// Struct 方式
// SELECT * FROM users WHERE name = 'jinzhu' OR name = 'bgbiao.top';
db.Where("name = 'jinzhu'").Or(User{Name: "bgbiao.top"}).Find(&users)

// Map 方式
// SELECT * FROM users WHERE name = 'jinzhu' OR name = 'bgbiao.top';
db.Where("name = 'jinzhu'").Or(map[string]interface{}{"name": "bgbiao.top"}).Find(&users)

```

**FirstOrCreate**

获取匹配的第一条记录, 否则根据给定的条件创建一个新的记录(仅支持 struct 和 map 条件)

```Go
// 未找到,就插入记录
if select * from users where name = 'non_existing') is null; insert into users(name) values("non_existing")
db.FirstOrCreate(&user, User{Name: "non_existing"})

// 找到
// select * from users where name = 'bgbiao.top'
db.Where(User{Name: "bgbiao.top"}).FirstOrCreate(&user)

## attrs参数:如果记录未找到，将使用参数创建 struct 和记录.
// 未找到(将条件前置，并将拆入数据填充到Attrs方法中)
db.Where(User{Name: "non_existing"}).Attrs(User{Age: 20}).FirstOrCreate(&user)

// 找到
db.Where(User{Name: "bgbiao.top"}).Attrs(User{Age: 30}).FirstOrCreate(&user)

## assgin参数: 不管记录是否找到，都将参数赋值给 struct 并保存至数据库
db.Where(User{Name: "non_existing"}).Assign(User{Age: 20}).FirstOrCreate(&user)


```

**子查询**

`*gorm.expr` 子查询

```Go
// SELECT * FROM "orders"  WHERE "orders"."deleted_at" IS NULL AND (amount > (SELECT AVG(amount) FROM "orders"  WHERE (state = 'paid')));
db.Where("amount > ?", DB.Table("orders").Select("AVG(amount)").Where("state = ?", "paid").QueryExpr()).Find(&orders)

```

**字段查询**

通常情况下，我们只想选择几个字段进行查询，指定你想从数据库中检索出的字段，默认会选择全部字段。

```Go
// SELECT name, age FROM users;
db.Select("name, age").Find(&users)

// SELECT name, age FROM users;
db.Select([]string{"name", "age"}).Find(&users)

// SELECT COALESCE(age,'42') FROM users;
db.Table("users").Select("COALESCE(age,?)", 42).Rows()
```

**排序(Order)**

```Go
// SELECT * FROM users ORDER BY age desc, name;
db.Order("age desc, name").Find(&users)

// 多字段排序
// SELECT * FROM users ORDER BY age desc, name;
db.Order("age desc").Order("name").Find(&users)

// 覆盖排序
//
db.Order("age desc").Find(&users1).Order("age", true).Find(&users2)
```

**限制输出**

```Go
// SELECT * FROM users LIMIT 3;
db.Limit(3).Find(&users)

// -1 取消 Limit 条件
// SELECT * FROM users LIMIT 10;
// SELECT * FROM users;
db.Limit(10).Find(&users1).Limit(-1).Find(&users2)
```

**统计count**

```Go
// SELECT count(*) from USERS WHERE name = 'jinzhu' OR name = 'bgbiao.top';
db.Where("name = ?", "jinzhu").Or("name = ?", "bgbiao.top").Find(&users).Count(&count)

// select count(*) from users where name = 'bgbiao.top'
db.Model(&User{}).Where("name = ?", "bgbiao.top").Count(&count)

// SELECT count(*) FROM deleted_users;
db.Table("deleted_users").Count(&count)

// SELECT count( distinct(name) ) FROM deleted_users;
db.Table("deleted_users").Select("count(distinct(name))").Count(&count)
```

**分组(grouo&having)**

```Go
rows, err := db.Table("orders").Select("date(created_at) as date, sum(amount) as total").Group("date(created_at)").Rows()
for rows.Next() {
  ...
}

rows, err := db.Table("orders").Select("date(created_at) as date, sum(amount) as total").Group("date(created_at)").Having("sum(amount) > ?", 100).Rows()
for rows.Next() {
  ...
}

type Result struct {
  Date  time.Time
  Total int64
}
db.Table("orders").Select("date(created_at) as date, sum(amount) as total").Group("date(created_at)").Having("sum(amount) > ?", 100).Scan(&results)
```

**连接查询**

```Go
rows, err := db.Table("users").Select("users.name, emails.email").Joins("left join emails on emails.user_id = users.id").Rows()
for rows.Next() {
  ...
}

db.Table("users").Select("users.name, emails.email").Joins("left join emails on emails.user_id = users.id").Scan(&results)

// 多连接及参数
db.Joins("JOIN emails ON emails.user_id = users.id AND emails.email = ?", "jinzhu@example.org").Joins("JOIN credit_cards ON credit_cards.user_id = users.id").Where("credit_cards.number = ?", "411111111111").Find(&user)
```

**pluck查询**

Pluck，查询 model 中的一个列作为`切片`，如果您想要查询多个列，您应该使用`Scan`。

```Go
var ages []int64
db.Find(&users).Pluck("age", &ages)

var names []string
db.Model(&User{}).Pluck("name", &names)

db.Table("deleted_users").Pluck("name", &names)

```

**scan扫描**

```Go
type Result struct {
  Name string
  Age  int
}

var result Result
db.Table("users").Select("name, age").Where("name = ?", "Antonio").Scan(&result)

// 原生 SQL
db.Raw("SELECT name, age FROM users WHERE name = ?", "Antonio").Scan(&result)
```

#### 更新

**更新所有字段**

`注意:`save会更新所有字段，即使没有赋值

```Go
db.First(&user)

user.Name = "bgbiao.top"
user.Age = 100
// update users set name = 'bgbiao.top',age=100 where id = user.id
db.Save(&user)

```

**更新修改字段**

使用 `Update` 和 `Updates` 方法

```Go
// 更新单个属性，如果它有变化
// update users set name = 'hello' where id = user.id
db.Model(&user).Update("name", "hello")

// 根据给定的条件更新单个属性
// update users set name = 'hello' where active = true
db.Model(&user).Where("active = ?", true).Update("name", "hello")

// 使用 map 更新多个属性，只会更新其中有变化的属性
// update users set name = 'hello',age=18,actived=false where id = user.id
db.Model(&user).Updates(map[string]interface{}{"name": "hello", "age": 18, "actived": false})

// 使用 struct 更新多个属性，只会更新其中有变化且为非零值的字段
db.Model(&user).Updates(User{Name: "hello", Age: 18})

// 警告：当使用 struct 更新时，GORM只会更新那些非零值的字段
// 对于下面的操作，不会发生任何更新，"", 0, false 都是其类型的零值
db.Model(&user).Updates(User{Name: "", Age: 0, Actived: false})
```

**更新选定字段**

如果想要更新或者忽略某些字段，可以先使用`Select`或`Omit`

```Go
// update users set name = 'hello' where id = user.id;
db.Model(&user).Select("name").Updates(map[string]interface{}{"name": "hello", "age": 18, "actived": false})

// Omit()方法用来忽略字段
// update users set age=18,actived=false where id = user.id
db.Model(&user).Omit("name").Updates(map[string]interface{}{"name": "hello", "age": 18, "actived": false})
```

**无hook更新**

上面的更新操作会会自动运行 model 的`BeforeUpdate`,`AfterUpdate`方法，来更新一些类似`UpdatedAt`的字段在更新时保存其 `Associations`, 如果你不想调用这些方法，你可以使用`UpdateColumn`,`UpdateColumns`

```Go
// 更新单个属性，类似于 `Update`
// update users set name = 'hello' where id = user.id;
db.Model(&user).UpdateColumn("name", "hello")

// 更新多个属性，类似于 `Updates`
// update users set name = 'hello',age=18 where id = user.id;
db.Model(&user).UpdateColumns(User{Name: "hello", Age: 18})

```

**批量更新**

`注意:` 使用struct实例更新时，只会更新非零值字段，弱项更新全部字段，建议使用map

```Go
// update users set name = 'hello',age=18 where id in (10,11)
db.Table("users").Where("id IN (?)", []int{10, 11}).Updates(map[string]interface{}{"name": "hello", "age": 18})

// 使用 struct 更新时，只会更新非零值字段，若想更新所有字段，请使用map[string]interface{}
db.Model(User{}).Updates(User{Name: "hello", Age: 18})

// 使用 `RowsAffected` 获取更新记录总数
db.Model(User{}).Updates(User{Name: "hello", Age: 18}).RowsAffected
```

**使用SQL计算表达式**

```Go
// update products set price = price*2+100 where id = product.id
DB.Model(&product).Update("price", gorm.Expr("price * ? + ?", 2, 100))

// update products set price = price*2+100 where id = product.id;
DB.Model(&product).Updates(map[string]interface{}{"price": gorm.Expr("price * ? + ?", 2, 100)})

// update products set quantity = quantity-1 where id = product.id;
DB.Model(&product).UpdateColumn("quantity", gorm.Expr("quantity - ?", 1))

// update products set quantity = quantity -1 where id = product.id and quantity > 1
DB.Model(&product).Where("quantity > 1").UpdateColumn("quantity", gorm.Expr("quantity - ?", 1))
```
#### 删除
**删除记录**

`注意:` 删除记录时，请确保主键字段有值，GORM 会通过主键去删除记录，如果主键为空，GORM 会删除该 model 的所有记录。

```Go
// 删除现有记录
// delete from emails where id = email.id;
db.Delete(&email)

// 为删除 SQL 添加额外的 SQL 操作
// delete from emails where id = email.id OPTION (OPTIMIZE FOR UNKNOWN)
db.Set("gorm:delete_option", "OPTION (OPTIMIZE FOR UNKNOWN)").Delete(&email)

```

**批量删除**

```Go
// delete from emails where email like '%jinzhu%'
db.Where("email LIKE ?", "%jinzhu%").Delete(Email{})

// 同上
db.Delete(Email{}, "email LIKE ?", "%jinzhu%")

```

**软删除**

如果一个model有`DeletedAt`字段，他将自动获得软删除的功能！
当调用`Delete`方法时， 记录不会真正的从数据库中被删除，只会将`DeletedAt`字段的值会被设置为当前时间。
在之前，我们可能会使用`isDelete`之类的字段来标记记录删除，不过在gorm中内置了`DeletedAt`字段，并且有相关hook来保证软删除。

```Go
// UPDATE users SET deleted_at="2020-03-13 10:23" WHERE id = user.id;
db.Delete(&user)

// 批量删除
// 软删除的批量删除其实就是把deleted_at改成当前时间,并且在查询时无法查到,所以底层用的是update的sql
db.Where("age = ?", 20).Delete(&User{})

// 查询记录时会忽略被软删除的记录
// SELECT * FROM users WHERE age = 20 AND deleted_at IS NULL;
db.Where("age = 20").Find(&user)

// Unscoped 方法可以查询被软删除的记录
// SELECT * FROM users WHERE age = 20;
db.Unscoped().Where("age = 20").Find(&users)

```

**物理删除**

`注意:` 使用`Unscoped().Delete()`方法才是真正执行sql中的`delete`语句.

```Go
// Unscoped 方法可以物理删除记录
// DELETE FROM orders WHERE id=10;
db.Unscoped().Delete(&order)
```

## 参考资料
