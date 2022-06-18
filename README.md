<p align="center"><a><img width="200px" src="https://raw.githubusercontent.com/gqzcl/blog_image/master/20220607232316.png"/></a></p>
<p align="center">
<img src="https://img.shields.io/badge/arod--im-im-green">
<img src="https://img.shields.io/github/go-mod/go-version/gqzcl/arod-im">
<img src="https://img.shields.io/github/license/gqzcl/arod-im">
</p>

# 关于 arod-im

arod-im 是一个使用golang实现的微服务架构（MSA）的分布式即时通信server，通过HTTP接口接收请求并通过Websocket推送消息，各服务间使用GRPC协议通信。

本项目使用[Kratos](https://github.com/go-kratos/kratos)作为微服务框架，使用[gnet](https://github.com/panjf2000/gnet)作为网络框架，构建了一个高性能，高可靠，高可用的IM应用。

arod-im 支持用户多端登录，支持私聊，群聊，聊天室等应用场景，也可以用作直播间的弹幕服务。

## 特性

* 高可靠、高可用、实时性、有序性
* 水平扩展
* 性能优化
* 拉取消息-批量拉取
* 组件式设计
* 指标监控，链路追踪
* 分布式事务
* 超时控制
* 注册中心、配置中心

通用性的设计，可以适应不同场景的需求，高度抽象，提供接口，根据业务需要自由扩展。底层的性能优化，为消息通信保驾护航。

## 快速开始

### 项目依赖

在开始前需要先安装好以下开源组件：

* [Nacos](https://github.com/alibaba/nacos) 用作服务注册中心以及配置中心
* [Redis](https://redis.io/) 用作会话信息存储
* [Kafka](https://kafka.apache.org/) 用作消息异步分发
* [protobuf](https://github.com/protocolbuffers/protobuf/releases)

### 快速部署

#### 通过Docker部署

* [ ] TODO

#### 手动部署

在开始前需要保证以及安装好上述组件并能正常使用。

首先将本项目clone到您的服务器

```bash
git clone https://github.com/gqzcl/arod-im.git
```

进入项目目录下，初始化项目

```bash
cd ./arod-im
export GOPROXY=https://goproxy.cn,direct
make init
make generate
```

然后修改配置文件，配置redis、nacos、kafka地址。

启动Logic服务

```bash
make run.logic
```

启动Connector服务

```bash
make run.connector
```

启动Job服务

```bash
make run.job
```

### 接口测试

[接口测试分享](https://www.eolink.com/share/index?shareCode=lSCXTf)

## 架构设计

"事件驱动架构设计"[^1]

## 实体

### 群消息

group_members(gid, uid, last_ack_msgid);

group_msgs(msgid,gid,sender_uid,time,content);

### 数据传输格式

数据传输格式选择的是 protobuf，原因如下：

* 易于使用，多语言支持，方便扩展
* 灵活（方便接口更新），高效，一条消息数据，用protobuf序列化后的大小是json的10分之一，xml格式的20分之一，是二进制序列化的10分之一，

## TODO

* [ ] 添加refresh token机制，添加refresh-token接口
* [ ] 添加获取连接服务地址的接口
* [ ] 为部分接口添加幂等逻辑
* [ ] 添加Promentheus的指标监控
* [ ] 添加服务链路追踪
* [ ] 添加服务间的超时控制
* [ ] 增加消息的两阶段确认机制

## 如何贡献


## 项目文档


## License

arod-im is open-sourced software licensed under the [MIT license](./LICENSE).

[^1]: # 事件驱动架构设计
