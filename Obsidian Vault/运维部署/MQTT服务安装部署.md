#边缘计算 #消息队列 #T200 #嵌入式 #分布式计算 #OPS #物联网 
# MQTT 服务安装部署
## 描述
**MQTT（Message Queuing Telemetry Transport，消息队列遥测传输协议）**

## 产品 EMQX

[EMQX: 大规模分布式物联网 MQTT 消息服务器](https://www.bing.com/ck/a?!&&p=786efa01c9f7833fJmltdHM9MTY2NzY5MjgwMCZpZ3VpZD0yMDNkZTZhYi0xOGM1LTY2Y2EtMTExZi1mNjIzMTk4ZjY3ZjgmaW5zaWQ9NTE3Mw&ptn=3&hsh=3&fclid=203de6ab-18c5-66ca-111f-f623198f67f8&psq=EMQX&u=a1aHR0cHM6Ly93d3cuZW1xeC5pby96aA&ntb=1)

### 简介
EMQX 是一款大规模可弹性伸缩的云原生分布式物联网 MQTT 消息服务器。
作为全球最具扩展性的 MQTT 消息服务器，EMQX 提供了高效可靠海量物联网设备连接，能够高性能实时移动与处理消息和事件流数据，帮助您快速构建关键业务的物联网平台与应用。

-   **开放源码**：基于 Apache 2.0 许可证完全开源，自 2013 年起 200+ 开源版本迭代。
-   **MQTT 5.0**：100% 支持 MQTT 5.0 和 3.x 协议标准，更好的伸缩性、安全性和可靠性。
-   **海量连接**：单节点支持 500 万 MQTT 设备连接，集群可扩展至 1 亿并发 MQTT 连接。
-   **高性能**：单节点支持每秒实时接收、移动、处理与分发数百万条的 MQTT 消息。
-   **低时延**：基于 Erlang/OTP 软实时的运行时系统设计，消息分发与投递时延低于 1 毫秒。
-   **高可用**：采用 Masterless 的大规模分布式集群架构，实现系统高可用和水平扩展。

### 连接
-   完整支持 MQTT v3.1、v3.1.1 和 v5.0 协议规范
    -   QoS 0、QoS 1、QoS 2 消息支持
    -   持久会话和离线消息支持
    -   保留消息（Retained Message）支持
    -   遗嘱消息（Will Message）支持
    -   共享订阅支持
    -   `$SYS/` 系统主题支持
-   MQTT 支持 4 种传输协议
    -   TCP
    -   TLS
    -   WebSocket
    -   QUIC（实验性）
-   HTTP 消息发布接口
-   网关
    -   CoAP
    -   LwM2M
    -   MQTT-SN
    -   Stomp
    -   GB/T 32960（企业版）
    -   JT/T 808（企业版）

更多 MQTT 扩展支持：
-   延迟发布
-   代理订阅
-   主题重写

### 安全
-   基于用户名/密码的身份认证，支持使用内置数据库、Redis、MySQL、PostgreSQL、MongoDB 作为数据源，也支持使用 HTTP Server 提供认证服务
-   基于 JWT 的身份认证与权限控制，支持 JWKs
-   MQTT 5.0 增强认证
-   PSK 身份验证
-   基于 Client ID、IP 地址，用户名的访问控制，支持使用内置数据库、Redis、MySQL、PostgreSQL、MongoDB 作为数据源，也支持使用 HTTP Server 提供授权服务
-   客户端黑名单支持

### 可伸缩性
-   多节点集群 (Cluster)
-   支持手动、dns、etcd、k8s 集群发现方式集群
-   多服务器节点桥接 (Bridge)

### 数据集成
-   SQL 语法数据集成，实时提取、过滤、丰富和转换 MQTT 消息或内部事件为用户所需格式，并将其发送到外部数据平台
-   通过 MQTT 与其他 Broker 或物联网平台进行双向数据桥接（如 EMQX Cloud，AWS IoT Core，Azure IoT Hub）
-   通过 Webhook 与其他应用集成

### 可靠性
-   过载保护
-   消息速率限制
-   连接速率限制

### 可观测性
-   客户端在线状态查询
-   集群状态与指标查询
-   Prometheus/StatsD 集成
-   自动网络分区恢复
-   在线日志追踪(Log Trace)
-   Erlang 运行时追踪工具

### 可扩展性
-   插件
-   钩子
-   gRPC 钩子扩展
-   gRPC 协议扩展

## EMQX 安装步骤

### 配置 EMQX Apt 源
~~~Shell
curl -s https://assets.emqx.com/scripts/install-emqx-deb.sh | bash
~~~

### 安装 EMQX 
~~~Shell
apt -y install emqx
~~~


### 启动 EMQX
~~~Shell
emqx start
~~~

## EMQX 文档

[EMQX 服务手册](https://www.emqx.io/docs/zh/v5.0/)



