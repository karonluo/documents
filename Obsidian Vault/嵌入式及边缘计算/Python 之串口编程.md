# Python 之串口编程
## 插件安装
~~~Shell
pip3 install pyserial
~~~
## 简单示例
~~~python
import serial
ser = serial.Serial('com1', 9600, timeout=1)
# serial.Serial("串口号 linux 应该是 /dev/ttyCOM1 或者 USB 等等", 波特率, 超时设置 单位 毫秒)
~~~
## Serial 所有参数介绍
~~~Python
ser = serial.Serial(
	port=None,       # 设备端口号，Windows 一般是 COM* Linux 一般是 /dev/tty*
	baudrate=9600,     # 波特率
	bytesize=EIGHTBITS,   # 数据位
	parity=PARITY_NONE,   # 奇偶校验位
	stopbits=STOPBITS_ONE, # 停止位
	timeout=None,      # 超时设置
	xonxoff=0,       # 是否打开 Software 软件流控制
	rtscts=0,        # 是否打开 RTS/CTS (硬件) 流控制
	interCharTimeout=None  # 字符间隔超时时间 None 为关闭 默认是关闭
)
~~~
## 接入端口不同平台举例
~~~Python
ser = serial.Serial("/dev/ttyUSB0",9600,timeout=0.5) #使用USB连接串行口
ser = serial.Serial("/dev/ttyAMA0",9600,timeout=0.5) #使用树莓派的GPIO口连接串行口
ser = serial.Serial(1,9600,timeout=0.5)#winsows系统使用com1口连接串行口
ser = serial.Serial("com1",9600,timeout=0.5)#winsows系统使用com1口连接串行口
ser = serial.Serial("/dev/ttyS1",9600,timeout=0.5)#Linux系统使用com1口连接串行口
~~~
## Serial 对象属性
~~~text
name:设备名字
port：读或者写端口
baudrate：波特率
bytesize：字节大小
parity：校验位
stopbits：停止位
timeout：读超时设置
writeTimeout：写超时
xonxoff：软件流控
rtscts：硬件流控
dsrdtr：硬件流控
interCharTimeout:字符间隔超时
~~~
## Serial 对象常用方法
~~~text
ser.isOpen()：查看端口是否被打开。
ser.open() ：打开端口。
ser.close()：关闭端口。
ser.read()：从端口读字节数据。默认1个字节。
ser.read_all():从端口接收全部数据。
ser.write("hello")：向端口写数据。
ser.readline()：读一行数据。
ser.readlines()：读多行数据。
in_waiting()：返回接收缓存中的字节数。
flush()：等待所有数据写出。
flushInput()：丢弃接收缓存中的所有数据。
flushOutput()：终止当前写操作，并丢弃发送缓存中的数据。
~~~
## 完整的 Sample Code
~~~Python
#!/usr/bin/python3
# -*- coding: UTF-8 -*-
'''
本 sample code 用于 分贝传感器，获取分贝值，判断噪音是否超过阈值
'''
import serial
import time
class DecibleDetection:
	_bps = 9600 # 波特率
	_portx = "/dev/ttyUSB0" # USB0号端口
	_timeout = 0.5 # 通讯超时 0.5 毫秒
	def detection(self):
		decibel_value = 0
		# 初始化并打开端口
		ser = serial.Serial(self._portx, self._bps, timeout=self._timeout)
		if ser.isOpen():
			str_data = "01 03 00 00 00 01 84 0A" # 要写入端口的信息
			bytes_data = bytes.fromhex(str_data) # 转换成 bytes
			ser.write(bytes_data) # 写入信息到端口中
			n = ser.inWaiting() # 等待返回
			time.sleep(self._timeout) # 暂停进程
			result_buffer = ser.read_all().hex() # 读取传感器返回的信息（十六进制） 
			ser.close() # 关闭端口
			decibel_value = int(result_buffer[6:10],16) / 10 # 按照传感器公布的数据协议 转换成整型获取当前分贝数
		return decibel_value
~~~
