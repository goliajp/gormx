package gormx

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestPg(t *testing.T) {
	var (
		cfg = &PgConfig{
			User:     "postgres",
			Password: "postgres",
			Host:     "127.0.0.1",
			Port:     5432,
			Dbname:   "postgres",
			Tz:       "Asia/Shanghai",
		}

		errCfg = &PgConfig{
			User:     "aaa",
			Password: "bbb",
			Host:     "127.0.0.1",
			Port:     5432,
			Dbname:   "ccc",
			Tz:       "Asia/Shanghai",
		}
	)
	Convey("with config", t, func() {
		p := NewPg(cfg)
		So(p, ShouldNotBeNil)
	})
	Convey("without config", t, func() {
		p := NewPg(nil)
		So(p, ShouldNotBeNil)
	})
	Convey("not initialized", t, func() {
		p := &Pg{}
		db := p.Open()
		So(db, ShouldBeNil)
	})
	Convey("connect failed", t, func() {
		p := NewPg(errCfg)
		db := p.Open()
		So(p, ShouldNotBeNil)
		So(db, ShouldBeNil)
	})
	Convey("normal", t, func() {
		p := NewPg(nil)
		db := p.Open()
		So(p, ShouldNotBeNil)
		So(db, ShouldNotBeNil)
	})
	Convey("open with params", t, func() {
		p := NewPg(nil)
		db := p.Open("postgres", "postgres", "postgres")
		So(p, ShouldNotBeNil)
		So(db, ShouldNotBeNil)
	})
	Convey("open with invalid params", t, func() {
		p := NewPg(nil)
		db := p.Open("postgres", "postgres", "postgres")
		db1 := p.Open("postgres", "postgres")
		db2 := p.Open("postgres", "postgres", "postgres", "localhost")
		db3 := p.Open("postgres", "postgres", "postgres", "localhost", "foo")
		db4 := p.Open("postgres", "postgres", "postgres", "localhost", "5432")
		db5 := p.Open("postgres", "postgres", "postgres", "localhost", "5432", "Asia/Shanghai")
		db6 := p.Open("postgres", "postgres", "postgres", "localhost", "5432", "Asia/Shanghai", "foo")
		So(p, ShouldNotBeNil)
		So(db, ShouldNotBeNil)
		So(db1, ShouldBeNil)
		So(db2, ShouldNotBeNil)
		So(db3, ShouldBeNil)
		So(db4, ShouldNotBeNil)
		So(db5, ShouldNotBeNil)
		So(db6, ShouldBeNil)
	})
	Convey("get gorm singleton", t, func() {
		p := NewPg(nil)
		db := p.DB()
		So(p, ShouldNotBeNil)
		So(db, ShouldNotBeNil)
	})
	Convey("get dsn string", t, func() {
		cfg := DefaultPgConfig
		dsn := pgDsn(cfg.Host, cfg.User, cfg.Password, cfg.Dbname, cfg.Tz, cfg.Port)
		expect := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
		So(dsn, ShouldEqual, expect)
	})
}
