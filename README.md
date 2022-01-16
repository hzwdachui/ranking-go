# 快速开始

## 启动

```bash
go run main.go
```

## 调用 api
- 投票: http://127.0.0.1:8080/vote?uid=123&rankId=1&starId=1
    - uid: 用户id
    - rankId: 榜单 id
    - starId: 明星 id
- 获取排行榜: http://127.0.0.1:8080/getRankList?rankId=1
  - rankId: 榜单 id


## 配置文件
config 目录下  
app.json 配置服务器参数


# 设计
项目分层
- main.go 主文件
- ./controller 注册路由
- ./servicer 接口业务逻辑
- ./data 数据模型和方法

为了方便本地启动，本项目都在内存中执行，没有做持久化，数据储存在两个全局变量中  
数据结构参考 redis zset 的实现，使用 sortedset 来存储排行榜

防止并发
- 用户并发投票：原子性的方法
  - SetNx：用户存在则返回存在的，不存在则创建新的并返回
  - Incr：用户投票数量+1，且不得超过投票数量限制
- 排行榜列表：原子性方法
  - SetNX：排行榜存在则返回存在的，不存在则创建并返回新的
  - Vote：排行榜中明星id对应票数增加