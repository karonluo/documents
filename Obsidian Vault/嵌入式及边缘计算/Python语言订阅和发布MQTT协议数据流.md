# Python语言订阅和发布MQTT协议数据流
## 描述

^df4d12

本教程文档仅介绍 MQTT 相关消息通过 GO 语言进行订阅和发布
关于 MQTT 服务安装请参考：
[MQTT服务安装部署](MQTT服务安装部署)


## 组件安装
~~~Shell
python3 -m pip install -i https://pypi.doubanio.com/simple paho-mqtt
~~~

## 代码实例

### 创建 MQTT 客户端
~~~Python
from paho.mqtt import client as mqtt_client
import time
broker = '127.0.0.1'
port = 1883
topic = "topic/test"
client_id = "python_mqtt_client_publisher"
def connect_mqtt():
    def on_connect(client, userdata, flags, rc):
        if rc == 0:
			print("Connected to MQTT Broker!")
        else:
			print("Failed to connect, return code %d\n", rc)
	client = mqtt_client.Client(client_id)
    client.on_connect = on_connect
	client.username_pw_set("admin", "pass@@word123")
	client.connect(broker, port)
	return client
~~~
### 发布者
~~~Python
def publish(client):
     msg_count = 0
	while True:
		time.sleep(1)
		msg = f"messages: {msg_count}"
		result = client.publish(topic, msg)
		status = result[0]
		if status == 0:
			print(f"Send `{msg}` to topic `{topic}`")
		else:
			print(f"Failed to send message to topic {topic}")
		msg_count += 1
def __main__():
	mqtt_client = connect_mqtt()
	publish(mqtt_client)
if __name__ == "__main__":
	 __main__()
~~~

### 订阅者
~~~Python
def subscribe(client):
    def on_message(client, userdata, msg):
        print(f"Received `{msg.payload.decode()}` from `{msg.topic}` topic")
    client.subscribe(topic)
    client.on_message = on_message

def __main__():
    client = connect_mqtt()
    subscribe(client)
    client.loop_forever()

if __name__ == '__main__':
    __main__()
    
~~~

## 完整实例
#### 发布消息
文件：publisher.py
~~~Python
from paho.mqtt import client as mqtt_client
import time
broker = '127.0.0.1'
port = 1883
topic = "topic/test"
client_id = "python_mqtt_client_publisher"

def connect_mqtt():
	def on_connect(client, userdata, flags, rc):
		if rc == 0:
			print("Connected to MQTT Broker!")
		else:
			print("Failed to connect, return code %d\n", rc)
	# Set Connecting Client ID
	client = mqtt_client.Client(client_id)
	client.on_connect = on_connect
	client.username_pw_set("admin", "pass@@word123")
	client.connect(broker, port)
	
	return client

def publish(client):
	msg_count = 0
	while True:
		time.sleep(1)
		msg = f"messages: {msg_count}"
		result = client.publish(topic, msg)
		status = result[0]
		if status == 0:
			print(f"Send `{msg}` to topic `{topic}`")
		else:
			print(f"Failed to send message to topic {topic}")
		msg_count += 1

def __main__():
	mqtt_client = connect_mqtt()
	publish(mqtt_client)

if __name__ == "__main__":
	__main__()

	
~~~

运行 publisher.py
~~~Shell
python3 publisher.py
~~~

#### 订阅消息
文件: subscriber.py
~~~Python
from paho.mqtt import client as mqtt_client
import time
broker = '127.0.0.1'
port = 1883
topic = "topic/test"
client_id = "python_mqtt_client_subscriber"
def connect_mqtt():
	def on_connect(client, userdata, flags, rc):
		if rc == 0:
			print("Connected to MQTT Broker!")
		else:
			print("Failed to connect, return code %d\n", rc)
	# Set Connecting Client ID
	client = mqtt_client.Client(client_id)
	client.on_connect = on_connect
	client.username_pw_set("admin", "pass@@word123")
	client.connect(broker, port)
	return client
def subscribe(client):
	def on_message(client, userdata, msg):
		print(f"Received `{msg.payload.decode()}` from `{msg.topic}` topic")
	client.subscribe(topic)
	client.on_message = on_message
def __main__():
	client = connect_mqtt()
	subscribe(client)
	client.loop_forever()
if __name__ == '__main__':
	__main__()

~~~
运行 subscriber.go
~~~Shell
go run subscriber.py
~~~

运行结果

![[Pasted image 20221109185418.png]]

