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

    * protoc-gen-go
    ```
    $ go get github.com/golang/protobuf/protoc-gen-go
    ```

    * protoc-gen-micro
    ```
    $ go get github.com/micro/micro/v3/cmd/protoc-gen-micro
    ```

* Install
```
$ go get github.com/micro/micro/v3
```

* Running a services
```
# start the micro service
$ micro server

# interacting with the micro server (admin, micro)
$ micro login
```
