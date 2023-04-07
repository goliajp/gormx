package gormx

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestDatabase(t *testing.T) {
	Convey("mysql: database", t, func() {
		m := NewMysql(nil)
		db := m.Open()
		if err := CreateDatabase(db, "tdb"); err != nil {
			t.Error(err)
		}
		hasTdb, err := HasDatabase(db, "tdb")
		if err != nil {
			t.Error(err)
		}
		So(hasTdb, ShouldBeTrue)
		if err := DropDatabase(db, "tdb"); err != nil {
			t.Error(err)
		}
		hasTdb, err = HasDatabase(db, "tdb")
		if err != nil {
			t.Error(err)
		}
		So(hasTdb, ShouldBeFalse)
	})
	Convey("pg: database", t, func() {
		p := NewPg(nil)
		db := p.Open()
		if err := CreateDatabase(db, "tdb"); err != nil {
			t.Error(err)
		}
		hasTdb, err := HasDatabase(db, "tdb")
		if err != nil {
			t.Error(err)
		}
		So(hasTdb, ShouldBeTrue)
		if err := DropDatabase(db, "tdb"); err != nil {
			t.Error(err)
		}
		hasTdb, err = HasDatabase(db, "tdb")
		if err != nil {
			t.Error(err)
		}
		So(hasTdb, ShouldBeFalse)
	})
}
