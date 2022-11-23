#GO语言 #边缘计算 #消息队列 #T100 #DEV
# Go 语言订阅和发布 MQTT 协议数据流

## MQTT 简介
![[MQTT协议介绍#^2f7a7b]]
## 描述

本教程文档仅介绍 MQTT 相关消息通过 GO 语言进行订阅和发布
关于 MQTT 服务安装请参考：
[MQTT服务安装部署](MQTT服务安装部署.md)
关于 GO 语言环境安装请参考：
[安装Go语言环境](安装Go语言环境.md)

## 组件安装

~~~Shell
go get github.com/eclipse/paho.mqtt.golang
~~~

## 代码实例
### 创建 MQTT 客户端
~~~Go
import (
	emqx "github.com/eclipse/paho.mqtt.golang"
)
var MQTTClient emqx.Client
func ConnectionMQTTServer(){
	opts := emqx.NewClientOptions()
	dsn := fmt.Sprintf("tcp://%s:%d", "127.0.0.1", 1883)
    opts.AddBroker(dsn)
    opts.SetClientID("mqtt_client")
    opts.SetUsername("username")
    opts.SetPassword("password")
    MQTTClient = emqx.NewClient(opts)
    MQTTClient.Connect()
}
~~~
### 发布消息
**注**：接创建客户端代码
~~~Go
func PublishMessage(msg string){
	// 向 主题 topic/test 发送消息
	topic:="topic/test"
	QoS:=byte(0)
	token := mqtt.MQTTClient.Publish(topic, QoS, false, msg)
	token.Wait()
}
~~~

### 订阅消息
**注**：接创建客户端代码
~~~Go
func ReceiveMessage(client emqx.Client, msg emqx.Message){
	// 打印输出接收到的数据
	fmt.Printf("Received message: %s from topic: %s\r\n", string(msg.Payload()), msg.Topic())
}
// 订阅函数
func Subscribe(){
	topic := "topic/test" // 设置订阅 topic/test 主题
	QoS:=byte(0) // QoS 设置服务质量等级， 请参阅 QoS 相关介绍
	// 订阅 topic/test 主题、QoS、获取消息的函数
    token := MQTTClient.Subscribe(topic, QoS, ReceiveMessage)
    token.Wait()
    fmt.Printf("Subscribed to topic: %s\r\n", topic)
    for {
        time.Sleep(time.Second)
    }
}
~~~

### 完整实例
#### 发布消息
文件：publisher.go
~~~Go
package main

import (
	"fmt"
	"time"
	emqx "github.com/eclipse/paho.mqtt.golang"
)
var MQTTClient emqx.Client
func ConnectionMQTTServer() {
	opts := emqx.NewClientOptions()
	dsn := fmt.Sprintf("tcp://%s:%d", "127.0.0.1", 1883)
	opts.AddBroker(dsn)
	opts.SetClientID("mqtt_client_publish")
	opts.SetUsername("admin")
	opts.SetPassword("pass@@word123")
	MQTTClient = emqx.NewClient(opts)
	if token := MQTTClient.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
}
func PublishMessage(msg string) {
	// 向 主题 topic/test 发送消息
	topic := "topic/test"
	QoS := byte(0)
	fmt.Println(msg)
	token := MQTTClient.Publish(topic, QoS, false, msg)
	token.Wait()
}
func main() {
	ConnectionMQTTServer()
	var i int
	for {
		i = i + 1
		msg := fmt.Sprintf("Message %d", i)
		PublishMessage(msg)
		time.Sleep(time.Second)
	}
}
~~~

运行 publisher.go
~~~Shell
go run publisher.go
~~~

#### 订阅消息
文件: subscriber.go
~~~Go
package main
import (
    "fmt"
    "time"
    emqx "github.com/eclipse/paho.mqtt.golang"
)

var MQTTClient emqx.Client

func ConnectionMQTTServer() {
    opts := emqx.NewClientOptions()
    dsn := fmt.Sprintf("tcp://%s:%d", "127.0.0.1", 1883)
    opts.AddBroker(dsn)
    opts.SetClientID("mqtt_client_subscribe")
    opts.SetUsername("admin")
    opts.SetPassword("pass@@word123")
    MQTTClient = emqx.NewClient(opts)
    if token := MQTTClient.Connect(); token.Wait() && token.Error() != nil {
        panic(token.Error())
    }
}

func ReceiveMessage(client emqx.Client, msg emqx.Message) {
    // 打印输出接收到的数据
    fmt.Printf("Received message: %s from topic: %s\r\n", string(msg.Payload()), msg.Topic())
}
// 订阅函数
func Subscribe() {
    topic := "topic/test" // 设置订阅 topic/test 主题
    QoS := byte(0)        // QoS 设置服务质量等级， 请参阅 QoS 相关介绍
    // 订阅 topic/test 主题、QoS、获取消息的函数
    token := MQTTClient.Subscribe(topic, QoS, ReceiveMessage)
    token.Wait()
    fmt.Printf("Subscribed to topic: %s\r\n", topic)
    for {
        time.Sleep(time.Second)
    }
}
func main() {
    ConnectionMQTTServer()
    Subscribe()
}
~~~
运行 subscriber.go
~~~Shell
go run subscriber.go
~~~

运行结果

![[Pasted image 20221107173658.png]]
