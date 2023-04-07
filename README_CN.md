# GormX - Golang Gorm 实用工具库

[![Go Report Card](https://goreportcard.com/badge/github.com/goliajp/gormx)](https://goreportcard.com/report/github.com/goliajp/gormx)
[![GoDoc](https://pkg.go.dev/badge/github.com/goliajp/gormx)](https://pkg.go.dev/github.com/goliajp/gormx)
[ENGLISH](README.md)
[日本語](README_JP.md)

GormX 是一个 Golang Gorm 实用工具库，可以帮助您轻松地连接到 MySQL 和 PostgreSQL 数据库


## 特点

- 轻松创建和管理连接到 MySQL 和 PostgreSQL 数据库
- 创建、检查和删除数据库
- 创建、检查和删除表
- 使用 `gormx.Model` 和 `gormx.RelatedModel` 定义表结构
- 预设作用域进行排序和分页

## 安装

```sh
go get -u github.com/goliajp/gormx
```

## 使用
### 创建连接
您可以使用 gormx.NewMysql() 和 gormx.NewPg() 函数分别创建 MySQL 和 PostgreSQL 数据库的连接。这些函数返回一个 *gorm.DB 实例，您可以用于 Gorm 操作

#### MySQL 连接
```go
m := gormx.NewMysql(nil) // nil is the default config: localhost:3306, root, root, mysql
db := m.DB() // db is *gorm.DB, then you can use Gorm normally
```

#### PostgreSQL 连接
```go
p := gormx.NewPg(nil) // nil is the default config: localhost:5432, postgres, postgres, postgres, Asia/Shanghai
db := p.DB() // db is *gorm.DB, then you can use Gorm normally
```

#### 自定义连接配置
您还可以为 MySQL 和 PostgreSQL 连接提供自定义配置

##### MySQL
```go
cfg := gormx.MysqlConfig{
    User:     "your username",
    Password: "your password",
    Addr:     "host:port",
    Dbname:   "connect dbname",
}
m := gormx.NewMysql(&cfg)
db := m.Open()
```

##### PostgreSQL
```go
cfg := gormx.PgConfig{
    User:   "your username",
    Password: "your password",
    Host:   "your host",
    Port:   9999, // int, your host port
    Dbname: "connect dbname",
    Tz:     "Asia/Shanghai", // timezone
}
p := gormx.NewPg(&cfg)
db := p.Open()
```

### 数据库操作
使用 gormx，您可以执行各种数据库操作，如创建、检查和删除数据库

#### 创建新数据库
```go
if err := gormx.CreateDatabase(db, "testdb"); err != nil {
    // error handler
}
```

#### 检查数据库是否存在
```go
hasTdb, err := gormx.HasDatabase(db, "testdb")
if err != nil {
    // error handler
}
```

#### 打开数据库
```go
tdb := m.Open("testdb")
```

#### 删除数据库
```go
if err := gormx.DropDatabase(db, "testdb"); err != nil {
    // error handler
}
```

### 表操作
使用 gormx，您可以执行各种表操作，如创建、检查和删除表

#### 删除表
```go
if err := gormx.DropTables(db, "table1", "table2"); err != nil {
    // error handler
}
```

#### 检查表是否存在
```go
hasTable, err := gormx.HasTable(db, "table1")
if err != nil {
    // error handler
}
```

### 定义表结构
gormx 提供了 gormx.Model 和 gormx.RelatedModel 来帮助您定义表结构

```go
type One struct {
    gormx.Model
    // attrs...
}
type Two struct {
    gormx.Model
    // attrs...
}
type OneTwo struct { // define the relationship of "one" and "two"
    gormx.RelatedModel
    OneId int
    TwoId int
    One *One
    Two *Two
}
```

### 使用预设范围
gormx 提供了预设范围，当您需要获取列表结果时非常有用

#### 按 created_at 降序排序结果
```go
var rs []Foo
db.Scopes(gormx.ScopeOrderByCreatedAtDesc).Find(&rs)
```

#### 按 updated_at 降序排序结果
```go
var rs []Foo
db.Scopes(gormx.ScopeOrderByUpdatedAtDesc).Find(&rs)
```

#### 分页
```go
var rs []Foo
db.Scopes(gormx.ScopePagination(1, 10)).Find(&rs)
```

## 贡献
我们欢迎对 GormX 的贡献！请随时在 GitHub 上提交问题、功能请求或拉取请求

## 许可
本项目采用MIT许可证 - 有关详细信息，请参阅 [LICENSE](LICENSE) 文件