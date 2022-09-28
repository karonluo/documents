# Ubuntu 20.04 中进行 Go 语言 Iris Web 框架项目布局 
## 初始化 Go 语言项目目录
~~~
cd ~
mkdir go-demo
cd go-demo
go mod init go-demo
go get github.com/kataras/iris/v12 # 获取 iris web 开发框架
~~~

## 创建 pkg 目录

**注1：没有特殊 cd 命令 则默认在 Go 语言项目目录中 ，后续章节不再做提醒**

**注2：布局目录介绍，包含的 sample code 均放在 go-demo，阅读者可以进行下载**
~~~
mkdir pkg
~~~
pkg\
作用主要是存放组件包的，对于 Go 语言项目，我们需要将所有的包都存放在这个目录下面。

### pkg 目录下创建 services 目录
pkg/services\
放置各种类型的接口服务层相关代码

~~~
mkdir pkg/services
~~~

### pkg 目录下创建 dao 目录
pkg/dao\
放置数据访问层相关代码

~~~
mkdir pkg/dao
~~~

### pkg 目录下创建 entites 目录
pkg/entities\
放置实体对象层相关代码 
~~~
mkdir pkg/entities
~~~

### pkg 目录下创建 biz 目录
pkg/biz\
放置业务逻辑实现层相关代码 
~~~
mkdir pkg/biz
~~~

最终的目录树结构如下：

~~~
├─ go-demo
│  ├─ cmd
│  ├─ conf
│  ├─ pkg
│  │  ├─ dao
│  │  ├─ biz
│  │  ├─ services
│  │  ├─ entities
│  │  └─ tools
│  ├─ web
│  │  ├─ js
│  │  ├─ images
│  │  ├─ css
│  │  └─ views
│  ├─ go.mod
│  ├─ go.sum
└──└─ main.go

~~~

## 参阅
>  <a target="_blank" href="https://github.com/karonluo/documents/blob/main/Ubuntu%2020.04%20%E4%B8%AD%E5%AE%89%E8%A3%85%20go%20%E8%AF%AD%E8%A8%80%E7%8E%AF%E5%A2%83.md">Ubuntu 20.04 Docker 容器环境配置</a>
> 
> <a target="_blank" href="https://github.com/karonluo/documents/blob/main/Ubuntu%2020.04%20%E4%B8%AD%E5%AE%89%E8%A3%85%20go%20%E8%AF%AD%E8%A8%80%E7%8E%AF%E5%A2%83.md">Ubuntu 20.04 中安装 go 语言环境</a>

