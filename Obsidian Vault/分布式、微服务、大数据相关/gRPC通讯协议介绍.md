# gRPC 通讯协议介绍
## 前言

现如今，微服务变得越来越流行，而服务间的通信也变得越来越重要，服务间通信本质是交换信息，而交换信息的中介/桥梁正是我们的API。

诚然，目前构建API最受欢迎的仍然是使用Restful（HTTP-JSON），因为它简单、快速、易懂。

但是在本文中，我们不妨尝试探索使用gRPC来构建我们的API，并在这个过程比较gRPC与Restful两者的异同。

最后，让我们来思考下，一个API应该是怎样的？如果从宏观的角度看，一个API应该是非常简单的，因为它就做了两件事：

-   1.客户端发送一个请求（Request）
-   2.服务端接收请求，并返回一个响应（Response）

这种思想在gRPC中也体现得非常明显，而从微观上看，构建一个API，我们可能需要考虑到：

-   1.我们使用什么数据模型？JSON、XML还是二进制流
-   2.我们使用什么协议传输？HTTP、TCP、还是HTTP/2
-   3.我们如何调用方法，以及处理错误
-   4.我们如何应对数据量大的情况
-   5.我们如何减少接口的延时等
-   ...

话不多说，让我们进入正题。

## 什么是gRPC

在聊聊什么是gRPC前，我们先来聊聊什么是RPC。

RPC，全称`Remote Procedure Call`，中文译为远程过程调用。通俗地讲，使用RPC进行通信，调用远程函数就像调用本地函数一样，RPC底层会做好数据的序列化与传输，从而能使我们更轻松地创建分布式应用和服务。

![](https://pic1.zhimg.com/80/v2-361866d30e1b42814b577280dd7afaf0_720w.webp)

而gRPC，则是RPC的一种，它是免费且开源的，由谷歌出品。使用gRPC，我们只需要定义好每个API的Request和Response，剩下的gRPC这个框架会帮我们自动搞定。

另外，gRPC的典型特征就是使用protobuf（全称protocol buffers）作为其接口定义语言（Interface Definition Language，缩写IDL），同时底层的消息交换格式也是使用protobuf。

## gRPC基本通信流程

![](https://pic3.zhimg.com/80/v2-13d685915ee28ac36b80b110d1deecca_720w.webp)

这是官方文档的一张图，通过这张图，我们可以大致了解下gRPC的通信流程：

1.gRPC通信的第一步是定义IDL，即我们的接口文档（后缀为.proto）

2.第二步是编译proto文件，得到存根（stub）文件，即上图深绿色部分。

3.第三步是服务端（gRPC Server）实现第一步定义的接口并启动，这些接口的定义在存根文件里面

4.最后一步是客户端借助存根文件调用服务端的函数，虽然客户端调用的函数是由服务端实现的，但是调用起来就像是本地函数一样。

以上就是gRPC的基本流程，从图中还可以看出，由于我们的proto文件的编译支持多种语言（Go、Java、Python等），所以gRPC也是跨语言的。

## gRPC VS Restful

gRPC和Restful之间的对比，历来是学习gRPC的必修课，我会从文档规范、消息编码、传输协议、传输性能、传输形式、浏览器的支持度以及数据的可读性、安全性等方面进行比较。

### 文档规范

文档规范这种东西有点见仁见智，在我看来，gRPC使用proto文件编写接口（API），文档规范比Restful更好，因为proto文件的语法和形式是定死的，所以更为严谨、风格统一清晰；而Restful由于可以使用多种工具进行编写（只要人看得懂就行），每家公司、每个人的攥写风格又各有差异，难免让人觉得比较混乱。

另外，Restful文档的过时相信很多人深有体会，因为维护一份不会过时的文档需要很大的人力和精力，而公司往往都是业务为先；而gRPC文档即代码，接口的更改也会体现到代码中，这也是我比较喜欢gRPC的一个原因，因为不用花很多精力去维护文档。

### 消息编码

消息编码这块，gRPC使用`protobuf`进行消息编码，而Restful一般使用`JSON`进行编码

### 传输协议

传输协议这块，gRPC使用`HTTP/2`作为底层传输协议，据说也可替换为其他协议，但目前还未考证；而RestFul则使用`HTTP`。

### 传输性能

由于gRPC使用protobuf进行消息编码（即序列化），而经protobuf序列化后的消息体积很小（传输内容少，传输相对就快）；再加上HTTP/2协议的加持（HTTP1.1的进一步优化），使得gRPC的传输性能要优于Restful。

### 传输形式

传输形式这块，gRPC最大的优势就是支持流式传输，传输形式具体可以分为四种（unary、client stream、server stream、bidirectional stream），这个后面我们会讲到；而Restful是不支持流式传输的。

### 浏览器的支持度

不知道是不是gRPC发展较晚的原因，目前浏览器对gRPC的支持度并不是很好，而对Restful的支持可谓是密不可分，这也是gRPC的一个劣势，如果后续浏览器对gRPC的支持度越来越高，不知道gRPC有没有干饭Restful的可能呢？

### 消息的可读性和安全性

由于gRPC序列化的数据是二进制，且如果你不知道定义的Request和Response是什么，你几乎是没办法解密的，所以gRPC的安全性也非常高，但随着带来的就是可读性的降低，调试会比较麻烦；而Restful则相反（现在有HTTPS，安全性其实也很高）

### 代码的编写

由于gRPC调用的函数，以及字段名，都是使用stub文件的，所以从某种角度看，代码更不容易出错，联调成本也会比较低，不会出现低级错误，比如字段名写错、写漏。

## gRPC的适用场景

从上面gRPC和Restful的比较中，我们其实也从侧面了解gRPC的优劣势，也能顺势推断出其应用场景。

总的来说，gRPC主要用于公司内部的服务调用，性能消耗低，传输效率高，服务治理方便。Restful主要用于对外，比如提供接口给前端调用，提供外部服务给其他人调用等，

## gRPC简单实践

一般来讲，实现一个gRPC服务端和客户端，主要分为这几步：

-   1.安装 protobuf 依赖
-   2.编写 proto 文件（IDL）
-   3.编译 proto 文件（生成stub文件）
-   4.编写server端，实现我们的接口
-   5.编写client端，测试我们的接口

### 1.安装 protobuf 依赖

```Shell
# 1.安装protoc
$ brew install protoc

# 2.检查安装是否成功
$ protoc --version
libprotoc 3.7.1

# 3.安装编译插件
$ export GO111MODULE=on  # 开启Go Module
$ go get google.golang.org/protobuf/cmd/protoc-gen-go \
         google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

### 2.编写 proto 文件

```Go
syntax = "proto3";

package greeter.srv;

option go_package = "proto/greeter";

service Greeter {
    rpc SayHello (HelloRequest) returns (HelloReply) {}
}

// The request message containing the user's name.
message HelloRequest {
    string name = 1;
}

// The response message containing the greetings
message HelloReply {
    string message = 1;
}
```

这里使用了官方文档的例子，算是比较简单的一个例子。首先第一行`syntax = "proto3"` 表明了我们使用的是proto3，而不是proto2，proto2是之前的版本。

而后面的`service Greeter`则表示定义一个Greeter的服务，这个服务下有SayHello接口，这个接口的请求是`HelloRequest`结构体，返回是`HelloReply`接口体（前面提到过API本质就是Request+Response！）。

同一份proto文件可以定义多个服务（不太建议，除非有关联），每个服务下面可以定义多个接口。

注：更多protobuf的语法，可以参考网上的其他教程

### 3.编译 proto 文件

编译命令是：

> protoc --go_out=. --go_opt=paths=source_relative  
> --go-grpc_out=. --go-grpc_opt=paths=source_relative  
> proto/greeter/greeter.proto  

注：proto文件的编译是我最想吐槽protobuf的一个点，首先是语法晦涩难懂，比如上述的`paths=source_relative`，其次是stub文件的存放位置，也需要多次尝试才能写入到我们想要的位置；最后是protobuf的各种插件，这些插件没办法输出版本，在使用的时候经常因为版本的不同遇到一些奇奇怪怪的问题。

### 4.server端的编写

server端的编写如下：

```Go
// greeter_server.go
type server struct {
}

// 实现我们的接口
func (s *server) SayHello(ctx context.Context, req *greeter.HelloRequest) (rsp *greeter.HelloReply, err error) {
 rsp = &greeter.HelloReply{Message: "Hello " + req.Name}
 return rsp, nil
}

func main() {
 listener, err := net.Listen("tcp", ":52001")
 if err != nil {
  log.Fatalf("failed to listen: %v", err)
 }
 // gRPC 服务器
 s := grpc.NewServer()
 // 将服务器与处理器绑定
 greeter.RegisterGreeterServer(s, &server{})

 //reflection.Register(s)
 fmt.Println("gRPC server listen in 52001...")
 err = s.Serve(listener)
 if err != nil {
  log.Fatalf("failed to serve: %v", err)
 }
}
```

前面在proto文件中，我们只是定义了SayHello接口，并没有实现，所以当我们要实现一个server端，第一步就需要实现我们的接口。最后就是一个服务器的基本流程（不管是HTTP，还是gRPC），即声明一个服务器实例、绑定处理器以及最后的运行。

### 5.client端的编写

```Go
// greeter_client.go
func main() {
 // 发起连接，WithInsecure表示使用不安全的连接，即不使用SSL
 conn, err := grpc.Dial("127.0.0.1:52001", grpc.WithInsecure())
 if err != nil {
  log.Fatalf("connect failed: %v", err)
 }

 defer conn.Close()

 c := greeter.NewGreeterClient(conn)

 ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
 defer cancel()

    // 虽然SayHello由远程服务端实现，但调用起来就像一个本地函数一样
 r, err := c.SayHello(ctx, &greeter.HelloRequest{Name: "World"})
 if err != nil {
  log.Fatalf("call service failed: %v", err)
 }
 fmt.Println("call service success: ", r.Message)
}
```

client端的实现也比较简单，就是发起连接、创建客户端实例、调用方法，以及得到结果。

### 源代码

本文所有源代码均在本人的Github上，需要的话[点击这里](https://link.zhihu.com/?target=https%3A//github.com/yangancode/go-business/tree/master/grpc)即可查阅

