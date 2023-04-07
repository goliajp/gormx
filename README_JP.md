# GormX - Golang Gorm ユーティリティ

[![Go Report Card](https://goreportcard.com/badge/github.com/goliajp/gormx)](https://goreportcard.com/report/github.com/goliajp/gormx)
[![GoDoc](https://pkg.go.dev/badge/github.com/goliajp/gormx)](https://pkg.go.dev/github.com/goliajp/gormx)

---
[ENGLISH](README.md)
[简体中文](README_CN.md)

GormX は、MySQL および PostgreSQL データベースに簡単に接続できる便利な Gorm ユーティリティです


## 機能

- MySQL および PostgreSQL データベースへの接続を簡単に作成および管理
- データベースの作成、確認、削除
- テーブルの作成、確認、削除
- `gormx.Model` と `gormx.RelatedModel` を使用してテーブル構造を定義
- ソートおよびページネーションのプリセットスコープ

## インストール

```sh
go get -u github.com/goliajp/gormx
```

## 使い方
### 接続の作成
gormx.NewMysql() および gormx.NewPg() 関数を使用して、MySQL および PostgreSQL データベースへの接続を作成できます。 これらの関数は、*gorm.DB インスタンスを返します。これを使用して、Gorm 操作を行うことができます

#### MySQL 接続
```go
m := gormx.NewMysql(nil) // nil is the default config: localhost:3306, root, root, mysql
db := m.DB() // db is *gorm.DB, then you can use Gorm normally
```

#### PostgreSQL 接続
```go
p := gormx.NewPg(nil) // nil is the default config: localhost:5432, postgres, postgres, postgres, Asia/Shanghai
db := p.DB() // db is *gorm.DB, then you can use Gorm normally
```

#### カスタム接続設
MySQL および PostgreSQL 接続の両方にカスタム設定を提供することもできます

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

### データベース操作
GormX ユーティリティを使用して、データベースの作成、確認、削除などのさまざまなデータベース操作を行うことができます。

#### 新しいデータベースの作成
```go
if err := gormx.CreateDatabase(db, "testdb"); err != nil {
    // error handler
}
```

#### データベースが存在するかどうかの確認
```go
hasTdb, err := gormx.HasDatabase(db, "testdb")
if err != nil {
    // error handler
}
```

#### データベースを開く
```go
tdb := m.Open("testdb")
```

#### データベースの削除
```go
if err := gormx.DropDatabase(db, "testdb"); err != nil {
    // error handler
}
```

### テーブル操作
GormX ユーティリティを使用して、テーブルの作成、確認、削除などのさまざまなテーブル操作を行うことができます

#### テーブルの削除
```go
if err := gormx.DropTables(db, "table1", "table2"); err != nil {
    // error handler
}
```

#### テーブルが存在するかどうかの確認
```go
hasTable, err := gormx.HasTable(db, "table1")
if err != nil {
    // error handler
}
```

### テーブル構造の定義
GormX は、gormx.Model および gormx.RelatedModel を提供して、テーブル構造を定義するのに役立ちます

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

### プリセットスコープの使用
GormX は、リストの結果を取得する必要がある場合に役立つプリセットスコープを提供しています

#### 按 created_at の降順で結果を並べ替え
```go
var rs []Foo
db.Scopes(gormx.ScopeOrderByCreatedAtDesc).Find(&rs)
```

#### 按 updated_at の降順で結果を並べ替える
```go
var rs []Foo
db.Scopes(gormx.ScopeOrderByUpdatedAtDesc).Find(&rs)
```

#### ページネーション
```go
var rs []Foo
db.Scopes(gormx.ScopePagination(1, 10)).Find(&rs)
```

## 貢献
GormX への貢献を歓迎します！GitHub で問題を提出したり、機能リクエストを提出したり、プルリクエストを提出したりしてください

## ライセンス
このプロジェクトはMITライセンスの下でライセンスされています。詳細については、 [LICENSE](LICENSE) ファイルを参照してください