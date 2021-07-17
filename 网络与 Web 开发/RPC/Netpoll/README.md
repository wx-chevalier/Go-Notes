# 1. Introduction

[Netpoll](https://github.com/cloudwego/netpoll) 底层使用 epoll 管理连接，连接读写均为 epoll 事件驱动。 提供了 zero-copy 读写操作的能力，提升效率并降低内存和 gc
开销。 同时 [Netpoll](https://github.com/cloudwego/netpoll) 附带 NIO Server 实现，支持快速创建高性能 NIO server。
