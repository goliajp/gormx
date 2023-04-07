package gormx

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestMysql(t *testing.T) {
	var (
		cfg = &MysqlConfig{
			User:     "root",
			Password: "root",
			Addr:     "127.0.0.1",
			Dbname:   "mysql",
		}

		errCfg = &MysqlConfig{
			User:     "aaa",
			Password: "bbb",
			Addr:     "127.0.0.1",
			Dbname:   "mysql",
		}
	)
	Convey("with config", t, func() {
		m := NewMysql(cfg)
		So(m, ShouldNotBeNil)
	})
	Convey("without config", t, func() {
		m := NewMysql(nil)
		So(m, ShouldNotBeNil)
	})
	Convey("not initialized", t, func() {
		m := &Mysql{}
		db := m.Open()
		So(db, ShouldBeNil)
	})
	Convey("connect failed", t, func() {
		m := NewMysql(errCfg)
		db := m.Open()
		So(m, ShouldNotBeNil)
		So(db, ShouldBeNil)
	})
	Convey("normal", t, func() {
		m := NewMysql(nil)
		db := m.Open()
		So(m, ShouldNotBeNil)
		So(db, ShouldNotBeNil)
	})
	Convey("open with params", t, func() {
		m := NewMysql(nil)
		db := m.Open("mysql", "root", "root")
		So(m, ShouldNotBeNil)
		So(db, ShouldNotBeNil)
	})
	Convey("open with invalid params", t, func() {
		m := NewMysql(nil)
		db := m.Open("mysql", "root", "root")
		db1 := m.Open("mysql", "root")
		db2 := m.Open("mysql", "root", "root", "localhost")
		db3 := m.Open("mysql", "root", "root", "localhost", "foo")
		So(m, ShouldNotBeNil)
		So(db, ShouldNotBeNil)
		So(db1, ShouldBeNil)
		So(db2, ShouldNotBeNil)
		So(db3, ShouldBeNil)
	})
	Convey("get gorm singleton", t, func() {
		m := NewMysql(nil)
		db := m.DB()
		So(m, ShouldNotBeNil)
		So(db, ShouldNotBeNil)
	})
	Convey("get dsn string", t, func() {
		cfg := DefaultMysqlConfig
		dsn := mysqlDsn(cfg.User, cfg.Password, "tcp", cfg.Addr, cfg.Dbname)
		expect := "root:root@tcp(127.0.0.1)/mysql?charset=utf8mb4&parseTime=True&loc=Local"
		So(dsn, ShouldEqual, expect)
	})
}
