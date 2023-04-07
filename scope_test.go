package gormx

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestScope(t *testing.T) {
	Convey("mysql: scope", t, func() {
		m := NewMysql(nil)
		db := m.Open()
		if err := CreateDatabase(db, "tdb"); err != nil {
			t.Error(err)
		}
		tdb := m.Open("tdb")
		type Foo struct {
			Model
			Name string
		}
		if err := tdb.AutoMigrate(new(Foo)); err != nil {
			t.Error(err)
		}
		for _, f := range []Foo{
			{Name: "foo1"},
			{Name: "foo2"},
			{Name: "foo3"},
		} {
			if err := tdb.Create(&f).Error; err != nil {
				t.Error(err)
			}
		}
		var fs []Foo
		if err := tdb.Scopes(ScopeOrderByCreatedAtDesc).
			Find(&fs).Error; err != nil {
			t.Error(err)
		}
		So(len(fs), ShouldEqual, 3)
		So(fs[0].Name, ShouldEqual, "foo3")
		So(fs[1].Name, ShouldEqual, "foo2")
		So(fs[2].Name, ShouldEqual, "foo1")
		if err := tdb.Model(&Foo{}).
			Where("name = ?", "foo2").
			Update("name", "foo2updated").Error; err != nil {
			t.Error(err)
		}
		if err := tdb.Scopes(ScopeOrderByUpdatedAtDesc, ScopePagination(1, 2)).
			Find(&fs).Error; err != nil {
			t.Error(err)
		}
		So(len(fs), ShouldEqual, 2)
		So(fs[0].Name, ShouldEqual, "foo2updated")
		So(fs[1].Name, ShouldEqual, "foo3")

		// clean up
		if err := DropDatabase(db, "tdb"); err != nil {
			t.Error(err)
		}
	})
	Convey("pg: scope", t, func() {
		p := NewPg(nil)
		db := p.Open()
		if err := CreateDatabase(db, "tdb"); err != nil {
			t.Error(err)
		}
		tdb := p.Open("tdb")
		type Foo struct {
			Model
			Name string
		}
		if err := tdb.AutoMigrate(new(Foo)); err != nil {
			t.Error(err)
		}
		for _, f := range []Foo{
			{Name: "foo1"},
			{Name: "foo2"},
			{Name: "foo3"},
		} {
			if err := tdb.Create(&f).Error; err != nil {
				t.Error(err)
			}
		}
		var fs []Foo
		if err := tdb.Scopes(ScopeOrderByCreatedAtDesc).
			Find(&fs).Error; err != nil {
			t.Error(err)
		}
		So(len(fs), ShouldEqual, 3)
		So(fs[0].Name, ShouldEqual, "foo3")
		So(fs[1].Name, ShouldEqual, "foo2")
		So(fs[2].Name, ShouldEqual, "foo1")
		if err := tdb.Model(&Foo{}).
			Where("name = ?", "foo2").
			Update("name", "foo2updated").Error; err != nil {
			t.Error(err)
		}
		if err := tdb.Scopes(ScopeOrderByUpdatedAtDesc, ScopePagination(1, 2)).
			Find(&fs).Error; err != nil {
			t.Error(err)
		}
		So(len(fs), ShouldEqual, 2)
		So(fs[0].Name, ShouldEqual, "foo2updated")
		So(fs[1].Name, ShouldEqual, "foo3")

		// clean up
		if err := DropDatabase(db, "tdb"); err != nil {
			t.Error(err)
		}
	})
}
