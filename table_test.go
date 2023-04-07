package gormx

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestTable(t *testing.T) {
	Convey("mysql: table", t, func() {
		m := NewMysql(nil)
		db := m.Open()
		if err := CreateDatabase(db, "tdb"); err != nil {
			t.Error(err)
		}
		tdb := m.Open("tdb")
		type Foo struct {
			Name        string
			Description string
		}
		if err := tdb.AutoMigrate(new(Foo)); err != nil {
			t.Errorf("create table failed: %v", err)
		}
		hasFoo, err := HasTable(tdb, "foo")
		if err != nil {
			t.Error(err)
		}
		So(hasFoo, ShouldBeTrue)
		if err := DropTables(tdb, "foo"); err != nil {
			t.Error(err)
		}
		hasFoo, err = HasTable(tdb, "foo")
		if err != nil {
			t.Error(err)
		}
		So(hasFoo, ShouldBeFalse)

		// clean up
		if err := DropDatabase(db, "tdb"); err != nil {
			t.Error(err)
		}
	})
	Convey("pg: table", t, func() {
		p := NewPg(nil)
		db := p.Open()
		if err := CreateDatabase(db, "tdb"); err != nil {
			t.Error(err)
		}
		tdb := p.Open("tdb")
		type Foo struct {
			Name        string
			Description string
		}
		if err := tdb.AutoMigrate(new(Foo)); err != nil {
			t.Errorf("create table failed: %v", err)
		}
		hasFoo, err := HasTable(tdb, "foo")
		if err != nil {
			t.Error(err)
		}
		So(hasFoo, ShouldBeTrue)
		if err := DropTables(tdb, "foo"); err != nil {
			t.Error(err)
		}
		hasFoo, err = HasTable(tdb, "foo")
		if err != nil {
			t.Error(err)
		}
		So(hasFoo, ShouldBeFalse)

		// clean up
		if err := DropDatabase(db, "tdb"); err != nil {
			t.Error(err)
		}
	})
}
