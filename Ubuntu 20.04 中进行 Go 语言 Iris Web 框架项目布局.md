# Ubuntu 20.04 中进行 Go 语言 Iris Web 框架项目布局 
## 初始化 Go 语言项目目录
~~~
cd ~
mkdir go-demo
cd go-demo
go mod init go-demo
go get github.com/kataras/iris/v12 # 获取 iris web 开发框架
~~~

**注1：没有特殊 cd 命令 则默认在 Go 语言项目目录中 ，后续章节不再做提醒**

**注2：布局目录介绍所包含的 sample code 均放在 go-demo，读者可以进行下载**

## 创建 pkg 目录

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
pkg/dao **放置数据访问层相关代码**

~~~
mkdir pkg/dao
~~~

### pkg 目录下创建 entites 目录
pkg/entities **放置实体对象层相关代码**
~~~
mkdir pkg/entities
~~~

### pkg 目录下创建 biz 目录
pkg/biz **放置业务逻辑实现层相关代码** 
~~~
mkdir pkg/biz
~~~

### pkg 目录下创建 tools 目录
pkg/tools **放置工具类相关代码** 
~~~
mkdir pkg/tools
~~~

## 创建 web 目录

需要注意的是：web 目录主要放置与 web 访问视图相关的文件，js, images, css 都是传统静态资源目录，如果是完全的前后端分离方案，如 vue.js 做 MVVM 模式，则连 web 目录都可以不用建立，直接使用 services 就可以了。


~~~
mkdir web
~~~
### web 目录下创建 views 目录
web/views **放置视图HTML模板**
~~~
mkdir web/views
~~~

### web 目录下创建 传统 WEB 项目的静态文件目录
~~~
mkdir web/js
mkdir web/images
mkdir web/css
~~~

## 创建 cmd 目录
将最后打包的所有web相关的文件、配置相关文件及可执行文件 main 放入其中，在部署的时候，发布该目录即可
~~~
mkdir cmd
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

