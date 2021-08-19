Golang Micro
============
```
1. Go高并发特性适合大型系统开发
2. Go编译迅速，无依赖环境，适合容器化
```

微服务整体架构
--------------
    * 前端应用 -
    * API 网关 - micro API
    * 接口层
        * 用户接口
        * 订单接口
        * 购物车接口
        * 支付接口
        * 分类接口
    * 应用层
    * 领域层
        * 领域服务
            * 用户服务
            * 订单服务
            * 购物车服务
            * 支付服务
            * 分类服务
        * 微服务技术
            * 注册中心 consul
            * 链路追踪 jaeger
            * 监控 prometheus
            * 熔断器 hystrix
            * 通信 gRPC
            * 限流 ratelimit
            * 负载均衡 selector
            * 配置中心 consul
            * 日志 zap + ELK
            * 协议 protobuf
        * 工具架构
            * kompose
            * docker-compose
            * micro cli
            * Docker
            * go-micro
        * 基础设施调用
    * 基础设施层
        * 基础设施调用
            * MySQL/Postgres
            * Redis

微服务基本介绍
--------------
```
# 微服务
    1. 是一种架构模式
    2. 相比较单体架构，微服务架构更加独立，能够单独更新和发布
    3. 微服务里面的服务仅仅用于某一个特定的业务功能

    -- 为什么需要
        1. 逻辑清晰
        2. 快速迭代
        3. 多语言灵活组合

    --- DDD Domain Driven Design 领域驱动设计
        微服务颗粒划分
        康威定律: Conway's Law
            组织 对应 微服务拆分
        1. 真正决定软件复杂性的是设计方法 -- 设计复杂性

        2. 有助于知道确定系统边界
        3. 能够聚焦在系统核心元素上
        4. 帮助我们拆分系统
    领域:
        领域: 领域是有范围界限，也可以说是边界
        核心领域: 核心域是业务系统的核心价值
        通用子域：所有子域的消费者，提供者通用服务
        支撑子域: 专注业务系统的某一重要业务
    界限上下文:
        不在于如何划分边界，而是在于如何控制边界
    领域模型:
        是解决问题的抽象表达
        领域: 反应我们业务上需要解决的问题
        模型：针对问题提出的解决方案

# DDD领域微服务四层架构:
    Interface 接口 --
    Application 应用层 --
    Domain 领域层 -- 业务规则
    Infrastructure 基础设施层 -- 中间件 云设施

# 设计原则
    1. 要使用领域驱动设计，而不是数据驱动设计，也不是界面驱动设计
    2. 要清晰边界的微服务 [服务该做什么；不该做什么]
    3. 职能清晰的分层
    4. 要做自己能hold住的微服务，而不是过度拆分微服务
```

Micro & Go-micro
----------------
> Go Micro is a framework for distributed systems development.
    * RPC
    ```
    RPC代指远程过程调用 Remote Procedure Call
    包含传输协议和编码协议(对象序列化)协议
    允许运行与一台计算机的程序调用另一台计算机的子程序
    ```
    * gRPC
    ```
    gRPC 是一个高性能、开源、的通用RPC框架
    基于HTTP2.0 协议标准设计开发
    支持多语言，默认采用Protocol Buffers 数据序列化协议
        gRPC Server: <-- proto Response
        gRPC Stub: --> proto request
    ```

    * Protocol Buffer
    ```
    是一种轻便高效的序列化结构化数据协议
    通常使用在存储数据和需要远程数据通信的程序
    跨语言、更小、更快、更简单
    加速站点内数据传输速度
    解决数据传输不规范问题

    -- Protocol Buffer 常用概念
        Note that the name of the message should be UpperCamelCase, and the name of the field should be lower_snake_case
        Message: 描述一个请求/响应的消息格式
            字段修饰符:
                singular: 表示成员有0个或者一个，一般省略不写
                repeated: 表示该字段可以包含0-N个元素
        字段标识: 消息的定义中，每个字段都有一个唯一的数值标签
        常用数据类型:
            double
            float
            int32/uint32/int64/uint64/sint32/sint64/fixed32/fixed64/sfixed32/sfixed64
            bool
            string
            bytes
            map
        Service服务定义:
            可以定义一个RPC服务接口
    ```
    * Micro
    > Micro is a cloud platform for API development
    - [micro-about](./misc/micro-3.0.png)
    ```
    1. Server
        API - HTTP Gateway maps HTTP/json -> RPC
        Auth - Authentication(验证) and Authenrization(授权) jwt token and rule based access control
        Broker - pubsub messaging for asynchronous communication and distributed notification
        Config - 动态配置
        Events - 事件流
        Network - inter-service networking, isolation and routing plane
        Proxy - remote access and external grpc request
        Runtime -
        Registry - service discovery and API endpoint explorer
        Store - Key-Value storage with TTL expiry and persistent crud to keep microservices stateless
    2. Framework: 框架
    3. Command Line:
    4. Environments:

    Micro service which do one thing well and communication over the network or via pubsub messaging where necessary
    ```
    * go-micro
    ```
    Micro:
        用来构建和管理分布式程序的系统
        Runtime(运行时) 用来管理配置，认证，网络
            工具集: micro
            docker pull micro/micro

            api: api网关
            broker: 允许异步消息的消息代理
            network: 通过微网络服务构建多云网络
            new: 服务模版生成器
            proxy: 建立在Go Micro上透明服务代理
            registry: 服务资源管理器
            store: 简单的状态存储
            web: Web仪表盘允许浏览服务

        Framework(程序开发框架): 用来方便编写微服务
            是对分布式系统的高度抽象
            提供分布式系统开发的核心库
            可插拔的架构，按需使用

            -- 组件:
                注册 Registry: 提供服务发现机制
                选择器 Selector: 能够实现负载均衡
                传输 Transport: 服务与服务之间通信接口
                broker: 提供异步通信的消息发布/订阅接口
                codec: 消息传递到两端时进行编码与解码

                Server, Client
    ```

GORM
----
```markdown
Go 语言实现数据库访问的ORM(对象关系映射)库

# 依赖库
    -- gorm
    go get github.com/jinzhu/gorm
    
    -- 数据库驱动
    go get github.com/go-sql-driver/mysql
    -- 
```

Docker 介绍
-----------
```
# 为什么需要Docker
    软件更新发布及部署低效， 过程繁琐需要人工介入
    环境一致性难以保证，不同环境之间迁移成本太高
    构建容易分发简单

# 应用场景
    构建运行环境
    微服务
    CI/CD 环境的一致性

# Docker 介绍
    客户端Client: docker command line
    服务器进程: Docker Daemon -- 管理镜像和容器
    镜像仓库: 存储镜像

    Images: 镜像模版
    Containers: 容器

    $ docker build
    $ docker push/pull
    $ docker run/start/stop/rm -p -v
    $ docker images rmi
    $ docker version
    $ docker ps

    # 启动Nginx服务
    $ docker run -p 80:80 nginx
```

Micro Getting Started
---------------------
* Dependencies
    * protobuf
    > protoc: The Google protobuf compiler
    ```
    # macOS
    $ brew install protobuf (brew upgrade protobuf)

    # Linux
    $ PROTOC_ZIP=protoc-3.14.0-linux-x86_64.zip
    $ curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v3.14.0/$PROTOC_ZIP
    $ sudo unzip -o $PROTOC_ZIP -d /usr/local bin/protoc
    $ sudo unzip -o $PROTOC_ZIP -d /usr/local 'include/*'
    $ rm -f $PROTOC_ZIP
    ```

    * protoc-gen-go [编译Golang插件]
    > is a protoc plugin to generate a Go protocol buffer package.
    ```
    $ go get github.com/golang/protobuf/protoc-gen-go
    ```

    * protoc-gen-micro
    > using protoc-gen-micro to reduce boilerplate code
    ```
    $ go get github.com/micro/micro/v3/cmd/protoc-gen-micro
    ```

* Install
```
# local install 
$ go get github.com/micro/micro/v3@v3.2.0

# Docker install
$ docker pull micro/micro
```

* Running a services
```
# start the micro service
$ micro server

# interacting with the micro server (admin, micro)
$ micro login

# list the services
$ micro services

# create the micro new
(using local install micro)
$ micro new project-name

(using docker micro image)
$ docker run --rm -v $(pwd):$(pwd) -v $(pwd) micro/micro new project-name

# 注意事项
    使用micro new 创建的go.mod 需要删除 重新go mod init github.com/
```

Go Module Settings
------------------
```markdown
# Go Module 基本设置
    开启GOMODULE
    export GO111MODULE=on
    1. Go module 使用GOPROXY环境变量解决无法使用go get 的问题
    export GOPROXY="https://goproxy.io"

# Go Module 私有仓库设置
    $ go env # 查看参数
    GOPRIVATE="*.chyidl.com"

# Go Module 设置Git转发
    go get 内部使用 git clone 命令，默认只支持共有仓库
    替换https为ssh请求
    $ git config --global url."ssh://git@git".insteadOf "https://"
```

gRPC
----
> gRPC is a high-performance, open-source, feature-rich RPC framework, originally developed by Google.
```markdown
RPC stands for Remote Procedure Calls.

# How gRPC works?
    Client has a generated stub that provides the same methods as the server
    The stub calls gRPC framework under the hood to exchange information over network.
    Client and server use stubs to interact with each other, so they only need to implement their core service logic.

From proto file, the server and client stub codes are generated by the protocol buffer compiler (or protoc).

```

- [Protobuf, FlatBuffers](https://capnproto.org/news/2014-06-17-capnproto-flatbuffers-sbe.html)

|Feature    |Protobuf   |Cap'n Proto    |SBE    |FlatBuffers    |
|:----------|:----------|:--------------|:------|:--------------|
|Schema evolution| yes|yes|caveats|yes|
| Zero-copy| no|yes|yes|yes|
|Random-access reads|no|yes|no|yes|
|other languages|yes|yes|no|no|

* Zero-copy
```markdown
data should be structured the same way in-memory and on the wire. thus avoiding costly encode/decode steps
```

- HTTP/2
> gRPC uses HTTP/2 as its transfer protocol.s
> gRPC use HTTP/2 as its transfer protocol
```markdown
# Binary framing 二进制传输流
    More performant and robust
    Lighter to transport, safer to decode
    Greate combination with Protocol Buffer

# compresses the headers using HPACK, 压缩头
    Reduce overhead and improve performance

# Multiplexing 
    Send multiple requests and responses in parallel over a single TCP connection
    Reduce latency and improve network utilization

# Server Push
    One client request, multiple response.
    Reduce round-trip latency
```

```markdown
# Under the HTTP/2 
    1. Single TCP connection carries multiple bidirectional streams
    2. Each stream has a unique ID and caries multiple bidirectional messages.
    3. Each message (request/response) is broken down into multiple binary frames.
    4. Frame is the smallest unit that carries different types of data: HEADERS, SETTINGS, PRIORITY, DATA, etc
    5. Frames from different streams are interleaved(交错) and then reassembled on the other side.
```

- HTTP/2 vs HTTP/1.1

| | HTTP/2| HTTP/1.1|
|:---|:----|:--------|
|TRANSFER PROTOCOL| Binary | Text|
| Headers| Compressed| Plain Text|
| Multiplexing| Yes| No|
| Requests per connection| Multiple| 1|
| Server Push| Yes| No|
|Release Year| 2015| 1997|

- Types of gRPC
    * Unary
    * Client Streaming
    * Server Streaming
    * Bidirectional Streaming
    ```markdown
    arbitrary order， flexible and no blocking
    ```

- gRPC vs. REST

| FEATURE       | GRPC              | REST          |
|:--------------|:------------------|:--------------|
| Protocol      |HTTP/2 fast        |HTTP/1.1 slow  |
| Payload       |Protobuf (binary, small)| JSON (text, large)|
| API contract  |Strict, required (.proto)|Loose, optional (OpenAPI)|
| Code generation| Built-in (protoc)| Third-party tools (Swagger)|
| Security      | TLS/SSL           | TLS/SSL |
| Streaming     | Bidirectional streaming | Client -> server request only |

- gRPC Suit
```markdown
# Microservices
    Low latency and high throughput communication
    Strong API contract

# Polyglot environments
    Code generation out of the box for many languages

# Point-to-point realtime communication
    Excellent support for bidirectional streaming

# Network constrained environments
    Lightweight message format
```

- ProtocolBuf
```markdown
1. Note that the name of the message should be UpperCamelCase, and the name of the field should be lower_snake_case
2. Each message field should be assigned a unique tag.
3. A tag is simply an arbitrary integer with the smallest value of 1, and the biggest value of 2^29-1. except for numbers from 19000 to 19999, as they're reserved for internal protocol buffers implementation.
4. Note that tags from 1 to 15 take only 1 byte to encode, while those from 16 to 2047 take 2 bytes.
```


Golang Interview
----------------
0. OS Scheduler
```markdown
CPU caches: 
NUMA:
PC: program counter / IP: Instruction pointer 记录将要执行下一行地址
```

2. Golang调度机制
```markdown
Concurrent: 并发
Parallel: 并行
Concurrency is about dealing with lots of things at once. Parallelism is about doing lots of things at once.
    
goroutine(协程) are a golang wrapper on top of threads(线程) and managed by Go runtime(运行时) rather than the operating system.
Goroutine does not have a one-to-one relationship with threads.

gorountine:
用户态调度器:

C、C++语言编译之后执行完全依赖操作系统内核控制执行
Golang 在编译时加入自己的调度器代码，执行上按照自己的调度器进行调度执行

Gorountine调度器组成部分:
    1. 数据结构
        全局Global Queue 存放等待运行的G
        G: Goroutine 存储goroutine的执行stack信息，goroutinue状态以及goroutinue的任务函数
        M: Thread 操作系统线程, 真正执行的计算资源，
            从P的本地队列中获取G，P队列为空时，M会从全局队列中那一批G放在P的本地队列，或从其他P的本地队列偷一半放到自己P的本地队列
            M的数量: go语言本身限制，go程序启动时，会设置M的最大数量，默认10000
            一个M阻塞会创建新的M
        P: Processor 逻辑CPU处理器(golang对CPU的抽象) runtime.GOMAXPROCS(numLogicalProcessors) 控制P = CPU核心数, P的数量决定系统内最大可并行的G的数量
            P中存放等待运行的G队列,新创建的G',优先加入到P的本地队列，如果队列满了，则会把本地队列中的一半G移动到全局队列
            P列表：所有的P都在程序启动时创建，并保存在数组中，最多有GOMAXPROCS可配置
            P的数量：由启动时环境变量$GOMAXPROCS或者runtime.GOMAXPROCS 决定, 意味着在程序执行的任意时刻都只有$GOMAXPROCS个goroutine同时运行
        Goroutine调度器和OS调度器时通过M结合起来的，每个M都代表1个内核线程，OS调度器负责把内核线程分配到CPU的核上执行
    调度器设计策略：
        复用线程: 避免频繁的创建、销毁线程，而是对线程的复用
            1.work stealing 机制
                > 当本线程无可执行的G时，尝试从其他线程绑定P偷取G，而不是销毁线程
            2. hand off机制
                当本地线程因为G进行系统调用阻塞时，线程释放绑定的P，把P转移给其他空闲的线程执行
        利用并行: GOMAXPROCS 设置P的数量，最多有GOMAXPROCS个线程分布在多个CPU上同时运行
        抢占: 在coroutine中要等待一个协程主动让出CPU才执行下一个协程， 在Go中，一个goroutine最多占用CPU 10ms, 防止其他goroutine被饿死，
        全局G队列: 新的调度器中依然有全局G队列，但功能已经被弱化，当M执行work stealing从其他P偷不到G时，可以从全局G队列中获取G
    
    特殊的M0和G0:
        M0: 是启动程序后编号为0的主线程,这个M对应的实例会在全局变量runtime.m0中，不需要在heap上分配，M0负责执行初始化操作和启动第一个G，在之后M0就和其他M一样
        G0: 是每次启动一个M都会第一个创建的goroutine, GO仅用于负责调度G，G0不知想任何可执行的函数，每个M都会有一个自己的G0，在调度或系统调用时会使用G0的栈栈空间，全局变量的G0是M0的G0
            G0负责协调写成切换schedule()函数
        全局队列到P本地队列的负载均衡:
            n = min(len(GQ)/GOMAXPROCS + 1, len(GQ/2))
        M自旋线程: 会寻找可运行的G(全局为空，则向其他的MP者偷取),系统中最多有GOMAXPROCS个自旋的线程 - 多余的线程会进入休眠

    2. 调度算法
        fifo: first in first out
        权重调度
        时间片调度
    3. 环境管理
        堆栈管理，Golang的栈管理使用连续栈实现方式，空间不足时，分配更大的栈并把旧栈拷贝到新栈

GPM之间关系:
    1. G需要绑定在M上才能运行
    2. M需要绑定P才能运行
    3. 程序中多个M并不会同时处于执行状态，最多只有GO MAXPROCS个M在执行
    4. 每个P维护一个G队列
    5. 当一个G被创建出来，或者变为可执行状态时，就把他放到P的可执行队列中
    6. 当一个G执行结束时，P会从队列中把G取出，如果此时P的队列为空，即没有其他G可以执行，就随机选择另外一个P，从其可执行的G中偷取一半(work-stealing调度算法)

线程分为: "内核态线程" -co-routine 和 "用户态线程"
线程是有CPU调度是抢占室 协程是由用户态调度是写协做时的
```