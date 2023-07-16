# go-toolbox

Go 语言实现的工具包

## 示例用法

查看命令帮助信息：
> go run main.go help word

### 单词格式转换

首字母大写：
> go run main.go word -s=abcefg -m=3

### 时间工具

获取当前时间：
> go run main.go time now

计算加和时间：
> go run main.go time calc -c="2029-09-04 12:02:33" -d=-2h 

### SQL 语句到结构体的转换

获取指定数据表结构体：
> go run main.go sql struct --username 数据库的账号 --password 数据库的密码 --db=数据库名称 --table "需要转换的表名"

### Json 数据到结构体的转换

> go run main.go json struct -s={\"name\":\"Alice\",\"age\":30,\"address\":\"123\"}

输出后结果如下：
```
2023/07/17 00:08:55 输出结果: type Baiyang struct {
    Age float64
    Address string
    Name string
}
```


