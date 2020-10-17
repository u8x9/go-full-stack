# NSQ

是go语言编写的一个开源的实时分布式内存消息队列：

+ 提倡分布式和分散的拓扑，没有单点故障，支持容错和高可用性，并提供可靠的消息交付保证

+ 支持横向扩展，没有任何集中式代理

+ 易于配置和部署，并且内置了管理界面

## 应用场景

+ 异步处理

+ 应用解耦

+ 流量削峰

## nsq 组件

### nsqd

守护进程，接收、排队并向客户端发送消息

```bash
nsqd -broadcast-address=127.0.0.1
```

+ `broadcast-address`：配置广播地址

如果搭配 `nsqlookupd` 使用，还需要指定 `nsqlookupd` 地址：

```bash
nsqd -broadcast-address=127.0.0.1 -lookupd-tcp-address=127.0.0.1:4160
```

### nsqlookupd

维护所有 nsqd 状态，提供服务发现的守护进程。它能为消费者查找特定`topic`下的nsqd提供了运行时的自动发现服务

### nsqadmin

一个实时监控集群状态、执行各种管理任务的web管理平台

```bash
nsqadmin -lookupd-http-address=127.0.0.1:4161
```
