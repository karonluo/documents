#PostgreSQL #数据库 #T200 #DEV #OPS 
# PostgreSQL 安装部署
## 描述

## 单机部署

### 1. 安装 PostgreSQL 12
~~~Shell
sudo apt -y install postgresql-12
~~~

### 2. 启动 PostgreSQL 服务
~~~Shell
sudo systemctl enable postgresql
sudo systemctl start postgresql
~~~

### 3. 配置用户密码

#### 3.1 切换 postgres 用户
~~~Shell
sudo su postgres
~~~

#### 3.2 启动 PSQL SHELL
~~~Shell
psql
~~~

#### 3.3 设置密码
~~~sql
ALTER USER postgres WITH PASSWORD 'pass@@word123'
~~~

### 4. 初始化数据存储

`Note` 以下步骤都是基于 postgres 用户

`Note` 先找到 PostgreSQL 安装目录，可以通过以下命令快速查找

~~~Shell
sudo find / -name initdb
/usr/lib/postgresql/12/bin/initdb
~~~

#### 4.1 初始化
~~~Shell
mkdir ~/data
/usr/lib/postgresql/12/bin/initdb -E UNICODE -D ~/data/
cd ~/data
~~~

#### 4.2 修改配置

修改默认端口和监听端口、最大链接数
~~~Shell
vi ~/data/postgresql.conf
~~~

~~~Text
Port = 5432 # 默认 5432
listen_address = "*" # 默认 localhost
max_connections = 100 # 默认 100
~~~

~~~Shell
vi /etc/postgresql/12/main/postgresql.conf
~~~

~~~Text
Port = 5432 # 默认 5432
listen_address = "*" # 默认 localhost
max_connections = 100 # 默认 100
~~~

修改服务器访问权限
~~~Shell
vi ~/data/pg_hba.conf
~~~

~~~Text
# IPv4 local connections:
host    all             all             127.0.0.1/32            trust
host    all             all             0.0.0.0/0               md5
~~~

~~~Shell
vi /etc/postgresql/12/main/pg_hba.conf
~~~

~~~Text
# IPv4 local connections:
host    all             all             127.0.0.1/32            md5
host    all             all             0.0.0.0/0               md5
~~~


### 5. 重启服务
退出 postgres 用户，至 root 用户并执行以下命令
~~~Shell
sudo systemctl restart postgresql
~~~

### 6. 推荐通过 Navicat Premium 访问 PostgreSQL
新增链接，参数参考：

用户: postgres

初始数据库: postgres
