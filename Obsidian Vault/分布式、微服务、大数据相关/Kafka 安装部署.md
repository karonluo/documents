# Kafka 安装部署
## 描述
Kafka是最初由Linkedin公司开发，是一个分布式、分区的、多副本的、多订阅者，基于 zookeeper 协调的分布式日志系统（也可以当做 MQ 系统），常见可以用于web/nginx日志、访问日志，消息服务等等，Linkedin 于 2010年贡献给了Apache 基金会并成为顶级开源项目。

主要应用场景是：日志收集系统和消息系统。

Kafka主要设计目标如下：

-   以时间复杂度为O(1)的方式提供消息持久化能力，即使对TB级以上数据也能保证常数时间的访问性能。
-   高吞吐率。即使在非常廉价的商用机器上也能做到单机支持每秒100K条消息的传输。
-   支持Kafka Server间的消息分区，及分布式消费，同时保证每个partition内的消息顺序传输。
-   同时支持离线数据处理和实时数据处理。
-   Scale out:支持在线水平扩展
## 安装部署步骤 (单机)

### 步骤一 下载代码
登录Apache kafka 官方下载。  
http://kafka.apache.org/downloads.html  
备注：`2.11-1.1.0`版本才与`JDK1.7`兼容，否则更高版本需要`JDK1.8`

### 步骤二 启动服务
运行kafka需要使用Zookeeper，所以你需要先启动Zookeeper，如果你没有Zookeeper，你可以使用kafka自带打包和配置好的Zookeeper（PS：**在kafka包里**）。参考：[[ZooKeeper 分布式应用程序协调服务]]

启动 Zookeeper
~~~Shell 
# 这是前台启动，启动以后，当前就无法进行其他操作（不推荐）
./zookeeper-server-start.sh ../config/zookeeper.properties

# 后台启动（推荐）
./zookeeper-server-start.sh  ../config/zookeeper.properties 1> /dev/null 2>&1 &
~~~

启动 Kafka
~~~Shell
# 修改配置文件
vim config/server1.properties
~~~

修改如下内容
~~~text
broker.id=0
listeners=PLAINTEXT://192.168.10.130:9092
log.dirs=kafka-logs
zookeeper.connect=localhost:2181
~~~

后台运行 kafka 服务
~~~shell
# 后台启动kafka
./kafka-server-start.sh ../config/server.properties 1>/dev/null 2>&1 &
~~~

### 步骤三 创建一个主题
创建一个名为“test”的Topic，只有一个分区和备份（2181是zookeeper的默认端口）

**输入以下命令以创建主题**
~~~shell
./kafka-topics.sh --create --zookeeper localhost:2181 --config max.message.bytes=12800000 --config flush.messages=1 --replication-factor 1 --partitions 1 --topic test
~~~

**命令参数说明**

+ --create： 指定创建topic动作
+ --topic：指定新建topic的名称
+ --zookeeper： 指定 kafka 连接 Zookeeper 的连接 Url 地址，该值和 server.properties 配置文件中的配置项 {zookeeper.connect} 一样
+ --config：指定当前topic上有效的参数值，[参数列表参考文档](http://kafka.apache.org/082/documentation.html#brokerconfigs)
+ --partitions：指定当前创建的 kafka 分区数量，默认为1个
+ --replication-factor：指定每个分区的复制因子个数，默认1个

创建好之后，可以通过运行以下命令，查看已创建的topic信息。
~~~Shell
./kafka-topics.sh --list --zookeeper localhost:2181 test
~~~
除了手工创建topic外，你也可以配置你的broker，当发布一个不存在的topic时自动创建topic。

**查看对应topic的描述信息**
~~~Shell
./kafka-topics.sh --describe --zookeeper localhost:2181  --topic test0
~~~

**命令参数说明**

+ --describe： 指定是展示详细信息命令
+ --zookeeper： 指定 kafka 连接 Zookeeper 的连接 Url 地址，该值和 server.properties 配置文件中的配置项 {zookeeper.connect} 一样
+ --topic：指定需要展示数据的topic名称

**Topic 信息修改**
~~~Shell
./kafka-topics.sh --zookeeper 192.168.187.146:2181 --alter --topic test0 --config max.message.bytes=128000
./kafka-topics.sh --zookeeper 192.168.187.146:2181 --alter --topic test0 --delete-config max.message.bytes
./kafka-topics.sh --zookeeper 192.168.187.146:2181 --alter --topic test0 --partitions 10 
./kafka-topics.sh --zookeeper 192.168.187.146:2181 --alter --topic test0 --partitions 3 ## Kafka分区数量只允许增加，不允许减少
~~~

**Topic 删除**
~~~Shell
./kafka-topics.sh --delete --topic test0 --zookeeper 192.168.187.146:2181
~~~


`Note: This will have no impact if delete.topic.enable is not set to true`.
默认情况下，删除是标记删除，没有实际删除这个Topic；

如果运行删除Topic，两种方式：  
+ 方式一：通过 delete 命令删除后，手动将本地磁盘以及 Zookeeper 上的相关 topic 的信息删除即可;
+ 方式二：配置 server.properties 文件，给定参数 delete.topic.enable=true，重启 kafka 服务，此时执行 delete 命令表示允许进行 Topic 的删除;

### 步骤四 发送消息

`Kafka`提供了一个命令行的工具，可以从输入文件或者命令行中读取消息并发送给`Kafka`集群。每一行是一条消息。

运行producer（生产者）,然后在控制台输入几条消息到服务器。

**备注：这里的 localhost:9092 不是固定的，需要根据 server.properties 中配置的地址来写这里的地址！**


~~~shell
./kafka-console-producer.sh --broker-list localhost:9092 --topic test
~~~

开始输入发送信息
~~~text
>this is a message
>this is another message
~~~
按 `Ctrl+C` 可以终止输入

### 步骤五 消费消息

Kafka也提供了一个消费消息的命令行工具，将存储的信息输出出来。

**备注：这里的localhost:9092不是固定的，需要根据server.properties中配置的地址来写这里的地址！**

~~~Shell
./kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic test --from-beginning
~~~

~~~text
>this is a message
>this is another message
~~~
按 `Ctrl+C` 可以终止读取消息

## 安装步骤（集群）
