
# 安装 Go 语言环境

- [安装 Go 语言环境](#ubuntu-2004-中安装-go-语言环境)
  - [下载及安装GO语言包](#下载及安装go语言包)
  - [修改系统配置](#修改系统配置)
  - [简易项目布局](#简易项目布局)


---

**注意：以下步骤均在Ubuntu 20.04 系统中完成。**

## 下载及安装GO语言包
~~~Shell
cd /
wget https://studygolang.com/dl/golang/go1.19.1.linux-amd64.tar.gz
tar xf go1.19.1.linux-amd64.tar.gz

~~~
## 修改系统配置

**注意： GOPOXY 是设置代理，目前中国处在局域网内部，访问互联网非常慢或无法访问，因此需要设置中国国内的代理 即：https://goproxy.cn** 

在 /etc/bash.bashrc 文件最后添加:
~~~Shell
echo "export GOPATH=\"/go\"" >> /etc/bash.bashrc
echo "export PATH=\$PATH:\$GOPATH\"/bin\"" >> /etc/bash.bashrc
echo "export GO111MODULE=on" >> /etc/bash.bashrc
echo "export GOPROXY=\"https://goproxy.cn,direct\"" >> /etc/bash.bashrc
echo "export GOMODCACHE=\$GOPATH/pkg/mod" >> /etc/bash.bashrc

~~~

激活环境
~~~bash
source /etc/bash.bashrc
~~~

## 简易项目布局
**注意:等配置完成 Microsoft Visual Studio Code IDE 后再行布局**
~~~shell
cd ~
mkdir -p ./goproj/go-demo
cd ./goproj/go-demo
go mod init go-demo # go 语言 mod 模式 初始化 mod 目录，做为项目目录 vscode 也会识别到这个目录
~~~

