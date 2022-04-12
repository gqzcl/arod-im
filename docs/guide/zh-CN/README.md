<p align="center"><a><img width="200px" src="https://cdn.jsdelivr.net/gh/gqzcl/blog_image/blog/arod-im.png"/></a></p>
<p align="center">
<img src="https://img.shields.io/badge/arod--im-im-green">
<img src="https://img.shields.io/github/go-mod/go-version/gqzcl/arod-im">
<img src="https://img.shields.io/github/license/gqzcl/arod-im">
</p>

# 关于 arod-im

arod-im 是一个通用的 IM 框架，通过高度抽象IM基本功能，可以自由组装实现不同业务场景下的IM服务
- 消息可靠性
- 离线消息
- 水平扩展
- 性能优化
- 拉取消息-批量拉取
- 组件式设计
- 指标监控，链路追踪，日志存储

通用性的设计，可以适应不同场景的需求，高度抽象，提供接口，根据业务需要自由扩展。
底层的性能优化，为消息通信保驾护航。

# 架构设计

## Comet 连接层

Comet维持与用户的长连接，仅负责将消息推送到用户和接收用户发送的消息
Comet将接收到的消息发送至Logic处理，同时也接收来自Job的消息

连接user

```
bucket->room->channel
      -> channel
```

可靠性:
消息发送成功：
消息送达：私聊-群聊-广播不做保证

## Logic 业务处理层

负责处理

## Job 异步消息消费层

Job从MQ捞取到消息后发送给对应Comet，Comet再将消息发送给对应用户

## 数据存储层

内存存储（redis），存储会话信息，在线状态等

持久化存储（Mysql），存储离线消息

# 实体

## 群消息


group_members(gid, uid, last_ack_msgid);

group_msgs(msgid,gid,sender_uid,time,content);

## 数据传输格式

数据传输格式选择的是 protobuf，原因如下：

- 易于使用，多语言支持，方便扩展
- 灵活（方便接口更新），高效，一条消息数据，用protobuf序列化后的大小是json的10分之一，xml格式的20分之一，是二进制序列化的10分之一，

## License

arod-im is open-sourced software licensed under the [MIT license](./LICENSE).