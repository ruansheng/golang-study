# golang-study

### 安装go-redis
```
go get github.com/go-redis/redis
```

### 往redis添加测试数据
```
go run demo_insert.go
```

### 测试并发查询
```
go run demo_lock.go
go run demo_channel.go
```
