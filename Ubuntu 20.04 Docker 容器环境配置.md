# Ubuntu 20.04 Docker 容器环境配置
## 1：下载安装 Docker 拉取 Ubuntu 20.04 初始化环境
> 宿主机环境
~~~
curl -fsSL https://get.docker.com -o get-docker.sh
sudo sh get-docker.sh
docker pull ubuntu:20.04
~~~
## 2：进入容器进行软件和网络初始化配置
### 2.1： 进入容器进行软件配置

> 宿主机环境
~~~
# 进入容器
docker run -it --name "ubuntuinit" ubuntu:20.04 /bin/bash
~~~
> 容器环境
~~~
# 更改 apt 源为阿里源

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

# 更新 apt 
apt update

# 安装常用软件
apt -y install vim net-tools openssh-server inetutils-ping g++ git cmake

# 完成后退出 容器
exit
~~~

### 2.2：提交保存生成新的 Docker 镜像
> 宿主机环境
~~~
docker commit -m "save software init" -a"gomicro" ubuntuinit gomicro:1.0
# 停止 命名为 ubuntuinit 的容器
docker stop ubuntuinit 
# 删除 命名为 ubuntuinit 的容器
docker rm ubuntuinit
~~~

### 2.3 重新启用刚才保存的镜像作为容器并配置网络
> 宿主机环境
~~~
docker run -itd --privileged --name "gomicro" -p 1080:80 -p 2222:22 gomicro:1.0 /usr/sbin/init
docker exec -it gomicro /bin/bash
~~~
> 容器环境
~~~
# 增加 ssh 配置节点允许 root 账户远程登录
echo "PermitRootLogin yes" >> /etc/ssh/sshd_config
# 重启 sshd 服务
systemctl restart sshd
# 修改 root 密码
passwd
# 退出容器 
exit
~~~
> 宿主机环境
~~~
# 尝试使用 ssh 登录 docker 容器
ssh root@127.0.0.1 -p 2222
~~~
> 容器环境
~~~
# 退出容器
exit 
~~~
> 宿主机环境
~~~
# 提交更新镜像
docker commit -m "update network" -a "gomicro" gomicro gomicro:1.0
# 至此Ubuntu 20.04 容器环境配置完成，同时可以分发提交的镜像
# 将镜像可以保存成 tar 包，分发到其他服务器宿主机上进行展开
docker save -o ./gomicro1.0.tar gomicro:1.0
~~~

