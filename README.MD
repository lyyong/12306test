# 第一步
修改fastiml/model/dbinit.go中的数据库连接信息

# 第二步
替换test/machine_test.go中的TODO下testMachine 替换成自己的实现
```go
machine = testMachine{}
```
# 第三步
运行BenchmarkMachine