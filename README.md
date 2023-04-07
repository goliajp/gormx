# GormX - Golang Gorm Utility

[![Go Report Card](https://goreportcard.com/badge/github.com/goliajp/gormx)](https://goreportcard.com/report/github.com/goliajp/gormx)
[![GoDoc](https://pkg.go.dev/badge/github.com/goliajp/gormx)](https://pkg.go.dev/github.com/goliajp/gormx)

---
[简体中文](README_CN.md)
[日本語](README_JP.md)

GormX is a useful Gorm utility for Golang, which makes it easy to connect to MySQL and PostgreSQL databases.

## Features

- Easily create and manage connections to MySQL and PostgreSQL databases
- Create, check, and drop databases
- Create, check, and drop tables
- Define table structures using `gormx.Model` and `gormx.RelatedModel`
- Preset scopes for sorting and pagination

## Installation

```sh
go get -u github.com/goliajp/gormx
```

## Usage
### Creating a Connection
You can create connections to MySQL and PostgreSQL databases using the gormx.NewMysql() and gormx.NewPg() functions respectively. These functions return a *gorm.DB instance, which you can use for Gorm operations.

#### MySQL Connection
```go
m := gormx.NewMysql(nil) // nil is the default config: localhost:3306, root, root, mysql
db := m.DB() // db is *gorm.DB, then you can use Gorm normally
```

#### PostgreSQL Connection
```go
p := gormx.NewPg(nil) // nil is the default config: localhost:5432, postgres, postgres, postgres, Asia/Shanghai
db := p.DB() // db is *gorm.DB, then you can use Gorm normally
```

#### Custom Connection Configuration
You can also provide a custom configuration for both MySQL and PostgreSQL connections.

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

### Database Operations
Using the gormx utility, you can perform various database operations such as creating, checking, and dropping databases.

#### Creating a New Database
```go
if err := gormx.CreateDatabase(db, "testdb"); err != nil {
    // error handler
}
```

#### Checking if a Database Exists
```go
hasTdb, err := gormx.HasDatabase(db, "testdb")
if err != nil {
    // error handler
}
```

#### Opening a Database
```go
tdb := m.Open("testdb")
```

#### Dropping a Database
```go
if err := gormx.DropDatabase(db, "testdb"); err != nil {
    // error handler
}
```

### Table Operations
Using the gormx utility, you can perform various table operations such as creating, checking, and dropping tables.

#### Dropping Tables
```go
if err := gormx.DropTables(db, "table1", "table2"); err != nil {
    // error handler
}
```

#### Checking if a Table Exists
```go
hasTable, err := gormx.HasTable(db, "table1")
if err != nil {
    // error handler
}
```

### Defining Table Structures
gormx provides gormx.Model and gormx.RelatedModel to help you define table structures.

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

### Using Preset Scopes
gormx provides preset scopes that can be helpful when you need to retrieve list results.

#### Sorting Results by created_at in Descending Order
```go
var rs []Foo
db.Scopes(gormx.ScopeOrderByCreatedAtDesc).Find(&rs)
```

#### Sorting Results by updated_at in Descending Order
```go
var rs []Foo
db.Scopes(gormx.ScopeOrderByUpdatedAtDesc).Find(&rs)
```

#### Pagination
```go
var rs []Foo
db.Scopes(gormx.ScopePagination(1, 10)).Find(&rs)
```

## Contributing
We welcome contributions to GormX! Feel free to submit issues, feature requests, or pull requests on GitHub.

## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.