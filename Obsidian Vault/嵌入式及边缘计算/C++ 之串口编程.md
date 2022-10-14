# C++ 之串口编程
## 描述

随着嵌入式开发在物联网行业中站的比重日益增加，Linux 环境下的C++也不断变得更为大众化。习惯了Window平台开发的开发人员， 都被Visual Studio的强大宠坏了， 无论是什么样的开发需求， 总能有现成的轮子可以直接拿来用。就好比这里要介绍的串口通信， 在Windows开发中， 无论是C++, MFC,还是C#, 巨硬大大都给我们做好了封装。可是在Linux下就没那么简单了，虽然开源， 但是很多的开发都偏底层，连一个标准库级别的串口通信SDK都没有，很是无奈。
作为开发人员， 要么自己花时间熟悉底层接口然后造轮子，要么去开源社区找一些别人造好的轮子。而实际上，C++的开发大多对项目的耦合性比较大，很多别人造的轮子也都是在作者自己开发项目的过程中慢慢积累的，那些公开的接口如果别人拿来直接用，并不见得会很友好。所以大部分情况下都是开发者自己根据系统级别的api来自己造轮子

当然大多数情况下使用现有的串口库会节省很多时间。
如果要使用现有库的话, <font color=#FFFF00 style='font-weight:bold'>ros-kinetic-serial</font> 库应该是一个不错的选择

## 串口通讯简单过程
1. 使用串口名( Linux 下面通常为 /dev/tty*, windows 下通常为 COM* ) 打开串口 
2. 给打开的串口设定参数， 如 波特率、数据位、停止位、校验位等等 
3. 向串口发送数据 
4. 从串口中接收数据

## C++ Sample Code

需要使用到的接口定义
~~~C++
// 头文件
#include <unistd.h>
#include <fcntl.h>
#include <termios.h>
// 串口打开关闭与读写
int open(const char *name, int flag)
int write(int fd, const void *data, size_t size)
int read(int fd, void *data, size_t size)
int close(int fd)
// 串口配置相关
struct termios;
tcgetattr(int fd, struct termios* tios)
cfsetispeed(struct termios*tios, int baudrate)
cfsetospeed(struct termios*tios, int baudrate)
int tcsetattr (int __fd, int __optional_actions, const struct termios *__termios_p)
~~~

### 头文件 serialport.hpp

~~~C++
#ifndef SERIALPORT_H
#define SERIALPORT_H
#include <string>
#include <vector>

struct termios;
class SerialPort
{
public:
	// 波特率
    enum BaudRate {
        BR0 = 0000000,
        BR50 = 0000001,
        BR75 = 0000002,
        BR110 = 0000003,
        BR134 = 0000004,
        BR150 = 0000005,
        BR200 = 0000006,
        BR300 = 0000007,
        BR600 = 0000010,
        BR1200 = 0000011,
        BR1800 = 0000012,
        BR2400 = 0000013,
        BR4800 = 0000014,
        BR9600 = 0000015,
        BR19200 = 0000016,
        BR38400 = 0000017,
        BR57600 = 0010001,
        BR115200 = 0010002,
        BR230400 = 0010003,
        BR460800 = 0010004,
        BR500000 = 0010005,
        BR576000 = 0010006,
        BR921600 = 0010007,
        BR1000000 = 0010010,
        BR1152000 = 0010011,
        BR1500000 = 0010012,
        BR2000000 = 0010013,
        BR2500000 = 0010014,
        BR3000000 = 0010015,
        BR3500000 = 0010016,
        BR4000000 = 0010017
    };
	// 数据位
    enum DataBits {
        DataBits5,
        DataBits6,
        DataBits7,
        DataBits8,
    };
	// 停止位
    enum StopBits {
        StopBits1,
        StopBits2
    };
	// 奇偶校验位
    enum Parity {
        ParityNone,
        ParityEven,
        PariteMark,
        ParityOdd,
        ParitySpace
    };

	// 打开端口的参数结构体
    struct OpenOptions {
        bool autoOpen;
        BaudRate baudRate;
        DataBits dataBits;
        StopBits stopBits;
        Parity parity;
        bool xon;
        bool xoff;
        bool xany;
        int vmin;
        int vtime;
    };

    static BaudRate BaudRateMake(unsigned long baudrate);

    static const OpenOptions defaultOptions;

    explicit SerialPort(const std::string& path, const OpenOptions options = defaultOptions);

    bool open();
    bool open(const std::string& path, const OpenOptions& options);

    bool isOpen() const;

    int write(const void *data, int length);
    int read(void *data, int length);


    void close();

    static std::vector<std::string > list();

protected:

    void termiosOptions(termios& tios, const OpenOptions& options);


private:
    std::string _path;
    OpenOptions _open_options;
    int _tty_fd;
    bool _is_open;
};


bool operator==(const SerialPort::OpenOptions& lhs, const SerialPort::OpenOptions& rhs);
bool operator!=(const SerialPort::OpenOptions& lhs, const SerialPort::OpenOptions& rhs);

#endif // SERIALPORT_H

~~~

### 实现文件  serialport.cpp
~~~C++
#include "serialport.h"
#include <dirent.h>
#include <unistd.h>
#include <fcntl.h>
#include <termios.h>

const SerialPort::OpenOptions SerialPort::defaultOptions = {
    true, //        bool autoOpen;
    SerialPort::BR9600, //    BaudRate baudRate;
    SerialPort::DataBits8, //    DataBits dataBits;
    SerialPort::StopBits1, //    StopBits stopBits;
    SerialPort::ParityNone,//    Parity parity;
    false,                  // input xon
    false,                  // input xoff
    false,                  // input xany
    0,                      // c_cc vmin
    50,                     // c_cc vtime
};

SerialPort::SerialPort(const std::string &path, const OpenOptions options)
    : _path(path), _open_options(options) {
    if(options.autoOpen) {
        _is_open = open(_path, _open_options);
    }
}


bool SerialPort::open() {
    return _is_open = open(_path, _open_options), _is_open;
}

bool SerialPort::open(const std::string &path, const OpenOptions &options) {

    if(_path != path) _path = path;
    if(_open_options != options) _open_options = options;

    _tty_fd = ::open(path.c_str(), O_RDWR | O_NOCTTY | O_NONBLOCK);
    if(_tty_fd < 0) {
        return false;
    }

    struct termios tios;
    termiosOptions(tios, options);
    tcsetattr(_tty_fd, TCSANOW, &tios);
    tcflush(_tty_fd, TCIOFLUSH);
    return true;
}

void SerialPort::termiosOptions(termios &tios, const OpenOptions &options) {


    tcgetattr(_tty_fd, &tios);

    cfmakeraw(&tios);
    tios.c_cflag &= ~(CSIZE | CRTSCTS);
    tios.c_iflag &= ~(IXON | IXOFF | IXANY | IGNPAR);
    tios.c_lflag &= ~(ECHOK | ECHOCTL | ECHOKE);
    tios.c_oflag &= ~(OPOST | ONLCR);

    cfsetispeed(&tios, options.baudRate);
    cfsetospeed(&tios, options.baudRate);

    tios.c_iflag |= (options.xon ? IXON : 0)
            | (options.xoff ? IXOFF: 0)
            | (options.xany ? IXANY : 0);

    // 数据位

    int databits[] =  {CS5, CS6, CS7, CS8};
    tios.c_cflag &= ~0x30;
    tios.c_cflag |= databits[options.dataBits];

    // 停止位
    if(options.stopBits == StopBits2) {
        tios.c_cflag |= CSTOPB;
    } else {
        tios.c_cflag &= ~CSTOPB;
    }

    // 奇偶校验位
    if(options.parity == ParityNone) {
        tios.c_cflag &= ~PARENB;
    } else {
        tios.c_cflag |= PARENB;

        if(options.parity == PariteMark) {
            tios.c_cflag |= PARMRK;
        } else {
            tios.c_cflag &= ~PARMRK;
        }

        if(options.parity == ParityOdd) {
            tios.c_cflag |= PARODD;
        } else {
            tios.c_cflag &= ~PARODD;
        }
    }

    tios.c_cc[VMIN] = options.vmin;
    tios.c_cc[VTIME] = options.vtime;
}

bool SerialPort::isOpen() const {
    return _is_open;
}

int SerialPort::write(const void *data, int length) {
    return ::write(_tty_fd, data, length);
}

int SerialPort::read(void *data, int length) {
    return ::read(_tty_fd, data, length);
}

void SerialPort::close() {
    ::close(_tty_fd);
    _is_open = false;
}

SerialPort::BaudRate SerialPort::BaudRateMake(unsigned long baudrate) {
    switch (baudrate) {
    case 50:
        return BR50;
    case 75:
        return BR75;
    case 134:
        return BR134;
    case 150:
        return BR150;
    case 200:
        return BR200;
    case 300:
        return BR300;
    case 600:
        return BR600;
    case 1200:
        return BR1200;
    case 1800:
        return BR1800;
    case 2400:
        return BR2400;
    case 4800:
        return BR4800;
    case 9600:
        return BR9600;
    case 19200:
        return BR19200;
    case 38400:
        return BR38400;
    case 57600:
        return BR57600;
    case 115200:
        return BR115200;
    case 230400:
        return BR230400;
    case 460800:
        return BR460800;
    case 500000:
        return BR500000;
    case 576000:
        return BR576000;
    case 921600:
        return BR921600;
    case 1000000:
        return BR1000000;
    case 1152000:
        return BR1152000;
    case 1500000:
        return BR1500000;
    case 2000000:
        return BR2000000;
    case 2500000:
        return BR2500000;
    case 3000000:
        return BR3000000;
    case 3500000:
        return BR3500000;
    case 4000000:
        return BR4000000;
    default:
        break;
    }
    return BR0;
}


std::vector<std::string> SerialPort::list() {
    DIR *dir;
    struct dirent *ent;
    dir = opendir("/dev");
    std::vector<std::string> ttyList;

    while(ent = readdir(dir), ent != nullptr) {
        if("tty" == std::string(ent->d_name).substr(0, 3)) {
            ttyList.emplace_back(ent->d_name);
        }
    }

    return ttyList;
}
bool operator==(const jzhw::SerialPort::OpenOptions& lhs, const jzhw::SerialPort::OpenOptions& rhs)
{
    return lhs.autoOpen == rhs.autoOpen
            && lhs.baudRate == rhs.baudRate
            && lhs.dataBits == rhs.dataBits
            && lhs.parity == rhs.parity
            && lhs.stopBits == rhs.stopBits
            && lhs.vmin == rhs.vmin
            && lhs.vtime == rhs.vtime
            && lhs.xon == rhs.xon
            && lhs.xoff == rhs.xoff
            && lhs.xany == rhs.xany;
}

bool operator!=(const SerialPort::OpenOptions& lhs, const SerialPort::OpenOptions& rhs){
    return !(lhs == rhs);
}

~~~
