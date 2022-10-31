# Kafka 安装部署
## 描述
Kafka是最初由Linkedin公司开发，是一个分布式、分区的、多副本的、多订阅者，基于 zookeeper 协调的分布式日志系统（也可以当做 MQ 系统），常见可以用于web/nginx日志、访问日志，消息服务等等，Linkedin 于 2010年贡献给了Apache 基金会并成为顶级开源项目。
主要应用场景是：日志收集系统和消息系统。
Kafka 主要设计目标如下：

-   以时间复杂度为O(1)的方式提供消息持久化能力，即使对TB级以上数据也能保证常数时间的访问性能。
-   高吞吐率。即使在非常廉价的商用机器上也能做到单机支持每秒 100K 条消息的传输。
-   支持 Kafka Server 间的消息分区及分布式消费，同时保证每个 partition 内的消息顺序传输。
-   同时支持离线数据处理和实时数据处理。
-   Scale out: 支持在线水平扩展

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
vim config/server.properties
~~~

修改如下内容
~~~Text
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
`注意`：kafka2.2 以上版本，不再使用 zookeeper 而是内置 bootstrap-server, 地址和配置文件中 listeners 一致

**例如应该输入以下命令创建主题(topic)**
~~~Shell
./kafka-topics.sh --create --topic test --replication-factor 1 --partitions 1 --bootstrap-server 192.168.10.130:9092
~~~
**命令参数说明**

+ --create： 指定创建topic动作
+ --topic：指定新建topic的名称
+ --zookeeper： 指定 kafka 连接 Zookeeper 的连接 Url 地址，该值和 server.properties 配置文件中的配置项 {zookeeper.connect} 一致 (kafka 2.2 以下使用)
+ --config：指定当前topic上有效的参数值，[参数列表参考文档](http://kafka.apache.org/082/documentation.html#brokerconfigs)
+ --partitions：指定当前创建的 kafka 分区数量，默认为1个
+ --replication-factor：指定每个分区的复制因子个数，默认1个
+ --bootstrap-server: kafka 内置调度器的链接 Url 地址，该值和server.properties 配置文件中的配置项 {listeners} 一致 (kafka 2.2 以上使用)

创建好之后，可以通过运行以下命令，查看已创建的 `topic` 信息。
`kafka 2.2` 以下版本：
~~~Shell
./kafka-topics.sh --list --zookeeper localhost:2181
~~~
`kafka 2.2` 以上版本:
~~~Shell
./kafka-topics.sh --list --bootstrap-server 0.0.0.0:9092
~~~
除了手工创建topic外，你也可以配置你的broker，当发布一个不存在的topic时自动创建topic。

**查看对应topic的描述信息**

`kafka 2.2` 以下版本:
~~~Shell
./kafka-topics.sh --describe --zookeeper localhost:2181  --topic test
~~~
`kafka 2.2` 以上版本:
~~~Shell
./kafka-topics.sh --describe --bootstrap-server 0.0.0.0:9092  --topic test
~~~

**命令参数说明**

+ --describe： 指定是展示详细信息命令
+ --zookeeper： 指定 kafka 连接 Zookeeper 的连接 Url 地址，该值和 server.properties 配置文件中的配置项 {zookeeper.connect} 一样
+ --topic：指定需要展示数据的topic名称
+ --bootstrap-server: kafka 内置调度器的链接 Url 地址，该值和server.properties 配置文件中的配置项 {listeners} 一致 (kafka 2.2 以上使用)

**Topic 信息修改**
`注意`: kakfa 版本 2.2 以下用 `--zookeeper` 2.2 以上则使用 `--bookstrap-server` 后续所有相关命令不再提醒。
~~~Shell
./kafka-topics.sh --zookeeper 192.168.187.146:2181 --alter --topic test --config max.message.bytes=128000
./kafka-topics.sh --zookeeper 192.168.187.146:2181 --alter --topic test --delete-config max.message.bytes
./kafka-topics.sh --zookeeper 192.168.187.146:2181 --alter --topic test --partitions 10 
./kafka-topics.sh --zookeeper 192.168.187.146:2181 --alter --topic test --partitions 3 ## Kafka分区数量只允许增加，不允许减少
~~~

**Topic 删除**
~~~Shell
./kafka-topics.sh --delete --topic test0 --zookeeper 192.168.187.146:2181
~~~


`Note: This will have no impact if delete.topic.enable is not set to true`.
默认情况下，删除是标记删除，没有实际删除这个Topic；

如果运行删除Topic，两种方式：  
+ 方式一：通过 delete 命令删除后，手动将本地磁盘以及 Zookeeper 上的相关 topic 的信息删除即可;
+ 方式二：配置 server.properties 文件，给定参数 `delete.topic.enable=true`，重启 kafka 服务，此时执行 delete 命令表示允许进行 Topic 的删除;

### 步骤四 发送消息

`Kafka` 提供了一个命令行的工具，可以从输入文件或者命令行中读取消息并发送给`Kafka`集群。每一行是一条消息。

运行 producer（生产者）,然后在控制台输入几条消息到服务器。

**备注：这里的 0.0.0.0:9092 不是固定的，需要根据 server.properties 中配置的地址来写这里的地址！**


~~~shell
./kafka-console-producer.sh --broker-list 0.0.0.0:9092 --topic test
~~~

开始输入发送信息
~~~Text
>this is a message
>this is another message
~~~
按 `Ctrl+C` 可以终止输入

### 步骤五 消费消息

Kafka也提供了一个消费消息的命令行工具，将存储的信息输出出来。

**备注：这里的 localhost:9092 不是固定的，需要根据 server.properties 中配置的地址来写这里的地址！**

~~~Shell
./kafka-console-consumer.sh --bootstrap-server 0.0.0.0:9092 --topic test --from-beginning
~~~

~~~Text
>this is a message
>this is another message
~~~
按 `Ctrl+C` 可以终止读取消息

如果你有2台不同的终端上运行上述命令，那么当你在运行生产者时，消费者就能消费到生产者发送的消息。

### 步骤六  设置多个 broker 集群（单机伪集群的配置）

到目前，我们只是单一的运行一个 broker，没什么意思。
对于 Kafka，一个 broker 仅仅只是一个集群的大小，让我们多设几个 broker。

首先为每个 broker 创建一个配置文件
~~~shell
cp config/server.properties config/server1.properties 
cp config/server.properties config/server2.properties 
~~~

编辑这两个配置文件
~~~shell
vim config/server1.properties 
~~~

~~~Text
broker.id=1
listeners=PLAINTEXT://192.168.10.130:9092
log.dirs=kafka-logs-1
zookeeper.connect=localhost:2181
~~~

~~~Shell
vim config/server2.properties: 
~~~

~~~Text
broker.id=2
listeners=PLAINTEXT://192.168.10.130:9093
log.dirs=kafka-logs-2
zookeeper.connect=localhost:2181
~~~

**备注1**：`listeners` 一定要配置成为IP地址；如果配置为 `localhost` 或服务器的 `hostname`, 在使用`java`发送数据时就会抛出异常：`org.apache.kafka.common.errors.TimeoutException: Batch Expired 。`，因为在没有配置 `advertised.host.name` 的情况下，`Kafka` 并没有像官方文档宣称的那样改为广播我们配置的 `host.name`，而是广播了主机配置的 `hostname`。远端的客户端并没有配置  `hosts`，所以自然是连接不上这个 `hostname` 的。

**备注2**：当使用java客户端访问远程的kafka时，一定要把集群中所有的端口打开，否则会连接超时。

以下参考防火墙 iptables 打开相关端口允许访问。
~~~Shell
/sbin/iptables -I INPUT -p tcp --dport 9092 -j ACCEPT
/sbin/iptables -I INPUT -p tcp --dport 9093 -j ACCEPT
/sbin/iptables -I INPUT -p tcp --dport 9094 -j ACCEPT
/etc/rc.d/init.d/iptables save
~~~

`broker.id`是集群中每个节点的唯一且永久的名称，我们修改端口和日志目录是因为我们现在在同一台机器上运行，我们要防止`broker`在同一端口上注册和覆盖对方的数据。

我们已经运行了 `zookeeper` 和 刚才的一个 `kafka` 节点，现在我们只需要在启动 2 个新的 `kafka` 节点。

~~~Shell
./kafka-server-start.sh ../config/server1.properties 1>/dev/null 2>&1 &
./kafka-server-start.sh ../config/server2.properties 1>/dev/null 2>&1 &
~~~

现在，我们创建一个新topic，把备份设置为：3

~~~Shell
./kafka-topics.sh --create --zookeeper localhost:2181 --replication-factor 3 --partitions 1 --topic my-replicated-topic
~~~

好了，现在我们已经有了一个集群了，我们怎么知道每个集群在做什么呢？运行命令 `describe topics`

~~~Shell
./kafka-topics.sh --describe --zookeeper localhost:2181 --topic my-replicated-topic
# 所有分区的摘要
Topic:my-replicated-topic    PartitionCount:1    ReplicationFactor:3    Configs:
# 提供一个分区信息，因为我们只有一个分区，所以只有一行。
Topic: my-replicated-topic    Partition: 0    Leader: 1    Replicas: 1,2,0    Isr: 1,2,0
~~~

+ “leader”：该节点负责该分区的所有的读和写，每个节点的leader都是随机选择的。
+ “replicas”：备份的节点列表，无论该节点是否是leader或者目前是否还活着，只是显示。
+ “isr”：“同步备份” 的节点列表，也就是活着的节点并且正在同步leader

其中 `Replicas` 和 `Isr` 中的 `1,2,0` 就对应着3个 `broker` 他们的 `broker.id` 属性！

我们运行这个命令，看看一开始我们创建的那个节点：

~~~Shell
./kafka-topics.sh --describe --zookeeper localhost:2181 --topic test
Topic:test    PartitionCount:1    ReplicationFactor:1    Configs:
Topic: test    Partition: 0    Leader: 0    Replicas: 0    Isr: 0
~~~

这并不奇怪，刚才创建的主题没有Replicas，并且在服务器“0”上，我们创建它的时候，集群中只有一个服务器，所以是“0”。

### 步骤七 测试集群的容错能力

#### 7.1 发布消息到集群

~~~Shell
./kafka-console-producer.sh --broker-list 192.168.10.130:9092 --topic my-replicated-topic
~~~

~~~Text
>cluster message 1
>cluster message 2
~~~
按 `Ctrl+C` 终止产生消息

#### 7.2 消费消息

~~~Shell
./kafka-console-consumer.sh --bootstrap-server 192.168.10.130:9093 --from-beginning --topic my-replicated-topic
~~~

~~~Text
cluster message 1
cluster message 2
~~~
按 `Ctrl+C` 终止消费消息

#### 7.3 干掉 Leader，测试集群容错

首先查询谁是 leader

~~~Shell
./kafka-topics.sh --describe --zookeeper localhost:2181 --topic my-replicated-topic
# 所有分区的摘要
Topic:my-replicated-topic    PartitionCount:1    ReplicationFactor:3    Configs:
# 提供一个分区信息，因为我们只有一个分区，所以只有一行。
Topic: my-replicated-topic    Partition: 0    Leader: 1    Replicas: 1,2,0    Isr: 1,2,0
~~~
可以看到 `Leader` 的 `broker.id` 为`1`，找到对应的 `Broker`

~~~Shell
jps -m
~~~

~~~Text
1010 Kafka ../config/server.properties
1030 QuorumPeerMain ../config/zookeeper.properties
1231 Bootstrap start start
7420 Kafka ../config/server2.properties
7111 Kafka ../config/server1.properties
9139 Jps -m
~~~

通过以上查询到 `Leader` 的 `PID`（`Kafka ../config/server-1.properties`）为`7111`，杀掉该进程

~~~Shell
# 杀掉该进程
kill -9 7111
# 再查询一下，确认新的Leader已经产生，新的Leader为broker.id=0
./kafka-topics.sh --describe --zookeeper localhost:2181 --topic my-replicated-topic
Topic:my-replicated-topic       PartitionCount:1        ReplicationFactor:3    Configs:
# 备份节点之一成为新的leader，而broker1已经不在同步备份集合里了
Topic: my-replicated-topic      Partition: 0    Leader: 0       Replicas: 1,0,2 Isr: 0,2
~~~

#### 7.4 再次消费消息，确认消息没有丢失

~~~Shell
./kafka-console-consumer.sh --zookeeper localhost:2181 --from-beginning --topic my-replicated-topic
~~~

~~~Text
cluster message 1
cluster message 2
~~~

消息正常接收。

## 安装步骤（集群）
