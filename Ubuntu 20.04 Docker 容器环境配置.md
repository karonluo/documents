

---
# Ubuntu 20.04 Docker 容器环境配置

- [Ubuntu 20.04 Docker 容器环境配置](#ubuntu-2004-docker-容器环境配置)
  - [1 下载安装 Docker 拉取 Ubuntu 20.04 初始化环境](#1-下载安装-docker-拉取-ubuntu-2004-初始化环境)
  - [2 进入容器进行软件和网络初始化配置](#2-进入容器进行软件和网络初始化配置)
    - [2.1 进入容器进行软件配置](#21-进入容器进行软件配置)
    - [2.2 提交保存生成新的 Docker 镜像](#22-提交保存生成新的-docker-镜像)
    - [2.3 重新启用刚才保存的镜像作为容器并配置网络](#23-重新启用刚才保存的镜像作为容器并配置网络)
  - [3 参阅](#3-参阅)

---
## 1 下载安装 Docker 拉取 Ubuntu 20.04 初始化环境
> 宿主机环境
~~~Shell
curl -fsSL https://get.docker.com -o get-docker.sh
sudo sh get-docker.sh
docker pull ubuntu:20.04
~~~
## 2 进入容器进行软件和网络初始化配置
### 2.1 进入容器进行软件配置

> 宿主机环境
~~~Shell
docker run -it --name "ubuntuinit" ubuntu:20.04 /bin/bash # 进入容器
~~~
> 容器环境
~~~Shell
echo "更改 apt 源为阿里源，因为国内访问互联网比较慢或无法访问，因此需要改到第三方源"

echo "deb http://mirrors.aliyun.com/ubuntu/ focal main restricted universe multiverse
deb-src http://mirrors.aliyun.com/ubuntu/ focal main restricted universe multiverse

deb http://mirrors.aliyun.com/ubuntu/ focal-security main restricted universe multiverse
deb-src http://mirrors.aliyun.com/ubuntu/ focal-security main restricted universe multiverse

deb http://mirrors.aliyun.com/ubuntu/ focal-updates main restricted universe multiverse
deb-src http://mirrors.aliyun.com/ubuntu/ focal-updates main restricted universe multiverse

deb http://mirrors.aliyun.com/ubuntu/ focal-proposed main restricted universe multiverse
deb-src http://mirrors.aliyun.com/ubuntu/ focal-proposed main restricted universe multiverse

deb http://mirrors.aliyun.com/ubuntu/ focal-backports main restricted universe multiverse
deb-src http://mirrors.aliyun.com/ubuntu/ focal-backports main restricted universe multiverse">/etc/apt/sources.list

apt update

apt -y install vim net-tools openssh-server inetutils-ping g++ git cmake # 安装常用软件

exit # 完成后退出容器
~~~

### 2.2 提交保存生成新的 Docker 镜像
> 宿主机环境
~~~Shell
docker commit -m "save software init" -a"gomicro" ubuntuinit gomicro:1.0
docker stop ubuntuinit  # 停止 命名为 ubuntuinit 的容器
docker rm ubuntuinit # 删除 命名为 ubuntuinit 的容器

~~~

### 2.3 重新启用刚才保存的镜像作为容器并配置网络
> 宿主机环境
~~~Shell
docker run -itd --privileged --name "gomicro" -p 1080:80 -p 2222:22 gomicro:1.0 /usr/sbin/init
docker exec -it gomicro /bin/bash
~~~
> 容器环境

**增加 ssh 配置节点允许 root 账户远程登录**
~~~Shell
echo "PermitRootLogin yes" >> /etc/ssh/sshd_config
systemctl restart sshd # 重启 sshd 服务
passwd # 修改 root 密码
exit # 退出容器回到宿主机环境
~~~

> 宿主机环境

**尝试使用 ssh 登录 docker 容器**
~~~Shell
ssh root@127.0.0.1 -p 2222
~~~
> 容器环境

**退出容器**
~~~Shell
exit 
~~~
> 宿主机环境


**提交更新镜像**
~~~Shell
docker commit -m "update network" -a "gomicro" gomicro gomicro:1.0
~~~

**将镜像可以保存成 tar 包，分发到其他服务器宿主机上进行展开**
~~~Shell
docker save -o ./gomicro1.0.tar gomicro:1.0
~~~

## 3 参阅 
> <a>Ubuntu 20.04 中 C++ 开发环境配置</a>\
> <a target="_blank" href="https://github.com/karonluo/documents/blob/main/Ubuntu%2020.04%20%E4%B8%AD%E5%AE%89%E8%A3%85%20go%20%E8%AF%AD%E8%A8%80%E7%8E%AF%E5%A2%83.md">Ubuntu 20.04 中安装 GO 语言环境</a>\
> <a>Ubuntu 20.04 中 Python3.0+ 开发环境配置</a>