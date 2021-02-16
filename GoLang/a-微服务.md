# 微服务(微服务概览与治理)

> 化繁为简 分而治之

[ppt文档](./resource/Week1-01-微服务.pdf)

## 微服务定义

> 微服务 是 SOA(面向服务)的一种实践

关注单一业务
分布式系统

1. 原子服务
2. 独立进程
3. 隔离部署
4. 去中心化

* 劣势
  1. 基础设施建设 复杂度高(监控报警系统)
  2. 请求放大 单次请求 下游请求放大百倍
  3. 分区数据架构
  4. 测试基于微服务架构的应用复杂
  5. 连锁故障 服务模块间的依赖

### 组件服务化

* kit: 一个微服务的基础框架
* service： 业务代码+kit依赖+第三方依赖
* rpc+message queue(kafka):轻量级通讯

本质上等同于 多个微服务组合(compose)完成了一个完整的用户场景(usercase)

### 按业务组织服务

40分钟
大前端(移动/Web) -> 网关接入 -> 业务服务 -> 平台服务 -> 基础设施(PasS/SaaS)

### 去中心化

* 数据去中心化（每个微服务独占DB）
* 治理去中心化
* 技术去中心化

### 基础设施自动化

* CICD
* Testing 单元测试（go test） api自动化测试（yapi）
* 在线运行时

### 可用性 兼容性设计

* Design For Failure思想
* 粗粒度进程间通讯

第三张图

### API Gateway(Envoy)

强耦合
服务端组装（前轻后重）

app-interface(BFF backend for frontend)

### BFF 数据组装

移动端->API Gateway->BFF->Microservice

### Microservice 划分

如何划分业务边界
cannal->kafka

### gRPC（HTTP2）

多语言
轻量级通讯 PB
可插拔
IDL 文件定义服务
设计理念
移动端 基于HTTP2设计 单连接 多路复用 http2相对于http1.1的优势
healthCheck(接口) 平滑发布

服务非对象 消息而非引用 粗粒度消息交互设计

### 服务发现

### 多集群

google SRE

### 多租户

GC 吊打 CMS
