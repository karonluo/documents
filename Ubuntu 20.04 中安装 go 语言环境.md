
# Ubuntu 20.04 中安装 Go 语言环境

- [Ubuntu 20.04 中安装 Go 语言环境](#ubuntu-2004-中安装-go-语言环境)
  - [下载及安装GO语言包](#下载及安装go语言包)
  - [修改系统配置](#修改系统配置)
  - [简易项目布局](#简易项目布局)
  - [参阅](#参阅)

---

## 下载及安装GO语言包
~~~
cd /
wget https://studygolang.com/dl/golang/go1.19.1.linux-amd64.tar.gz
tar xf go1.19.1.linux-amd64.tar.gz

~~~
## 修改系统配置

**注意： GOPOXY 是设置代理，目前中国处在局域网内部，访问互联网非常慢或无法访问，因此需要设置中国国内的代理 即：https://goproxy.cn** 

在 /etc/bash.bashrc 文件最后添加:
~~~
echo "export GOPATH=\"/go\"" >> /etc/bash.bashrc
echo "export PATH=\$PATH:\$GOPATH\"/bin\"" >> /etc/bash.bashrc
echo "export GO111MODULE=on" >> /etc/bash.bashrc
echo "export GOPROXY=\"https://goproxy.cn,direct\"" >> /etc/bash.bashrc
echo "export GOMODCACHE=\$GOPATH/pkg/mod" >> /etc/bash.bashrc


~~~

激活环境
~~~
source /etc/bash.bashrc
~~~

## 简易项目布局
**注意:等配置完成 Microsoft Visual Studio Code IDE 后再行布局**
~~~
cd ~
mkdir -p ./goproj/go-demo
cd ./goproj/go-demo
go mod init go-demo # go 语言 mod 模式 初始化 mod 目录，做为项目目录 vscode 也会识别到这个目录

~~~

## 参阅
>  <a target="_blank" href="https://github.com/karonluo/documents/blob/main/Ubuntu%2020.04%20%E4%B8%AD%E5%AE%89%E8%A3%85%20go%20%E8%AF%AD%E8%A8%80%E7%8E%AF%E5%A2%83.md">Ubuntu 20.04 Docker 容器环境配置</a>\
> <a target="_blank" href="https://github.com/karonluo/documents/blob/main/Ubuntu%2020.04%20%E4%B8%AD%E8%BF%9B%E8%A1%8C%20Go%20%E8%AF%AD%E8%A8%80%20Iris%20Web%20%E6%A1%86%E6%9E%B6%E9%A1%B9%E7%9B%AE%E5%B8%83%E5%B1%80.md">Ubuntu 20.04 中进行 Go 语言 Iris Web 框架项目布局</a>

