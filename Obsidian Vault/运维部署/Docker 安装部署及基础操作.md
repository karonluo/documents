# Docker 安装部署及基本操作
## 安装部署
### 在 Ubuntu 20.04 中安装 Docker 
~~~Shell
curl -fsSL https://get.docker.com -o get-docker.sh
sudo sh get-docker.sh
~~~
### 基础操作
#### 1.拉取镜像
docker pull <镜像名称>:<镜像版本>
例如：
~~~Shell
docker pull ubuntu:20.04
~~~
#### 2.查询本机已有镜像
docker images
#### 3.删除镜像
docker image rm <镜像名称>:<镜像版本>
例如:
~~~Shell
docker image rm ubuntu:20.04
~~~
#### 4.启动镜像并进入容器
docker run -it --name <自定义容器命名> --privileged \[yes\] <镜像名称>:<镜像版本> -p <宿主机端口号>:<映射的容器端口> -p <宿主机端口号>:<映射的容器端口> </bin/bash>
例如:
~~~Shell
docker run -it --name "test" --privileged yes ubuntu:20.04 -p 2222:22 -p 8080:80 /bin/bash
~~~
#### 5.查询正在运行的容器
docker ps 
#### 6.停止容器
docker stop <容器名称>|<容器编号>
例如：
~~~Shell
docker stop test
~~~
或
~~~Shell
docker stop 36011a8ba27d
~~~
#### 7.删除容器
docker rm <容器名称>|<容器编号>
注意：删除容器之前必须先调用停止容器
~~~Shell
docker rm test
~~~
或
~~~Shell
docker rm 36011a8ba27d
~~~
#### 8.提交容器更新镜像
docker commit -m <备注> -a <镜像作者名> <容器名称>|<容器编号> <镜像名称>:<镜像版本号>
~~~Shell
docker commit -m "update network" -a "gomicro" gomicro gomicro:1.0
~~~
#### 9.导出镜像包

docker save -o <导出的镜像包路径> <镜像名称>:<镜像版本号>
~~~Shell
docker save -o ./gomicro1.0.tar gomicro:1.0
~~~

#### 10.导入镜像包
docker load -i <导入的镜像包路径> 
~~~Shell
docker load -i ./gomicro1.0.tar
~~~
