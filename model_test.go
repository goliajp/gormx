package gormx

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestModel(t *testing.T) {
	Convey("mysql: model and relatedModel", t, func() {
		type One struct {
			Model
			Name        string `json:"name"`
			Description string `json:"description"`
		}

		type Two struct {
			RelateModel
			OneId   int `json:"oneId" gorm:"index"`
			ThreeId int `json:"threeId" gorm:"index"`
		}

		type Three struct {
			Model
			Key   string `json:"key"`
			Count int    `json:"count"`
		}
		m := NewMysql(nil)
		db := m.Open()
		if err := CreateDatabase(db, "tdb"); err != nil {
			t.Error(err)
		}
		tdb := m.Open("tdb")
		if err := tdb.AutoMigrate(new(One), new(Two), new(Three)); err != nil {
			t.Errorf("mysql: automigrate failed: %v", err)
		}
		one1 := One{
			Name:        "one1-name",
			Description: "one1-description",
		}
		if err := tdb.Create(&one1).Error; err != nil {
			t.Errorf("mysql: create one1 failed: %v", err)
		}
		three1 := Three{
			Key:   "three1-key",
			Count: 1,
		}
		if err := tdb.Create(&three1).Error; err != nil {
			t.Errorf("mysql: create three1 failed: %v", err)
		}
		two1 := Two{
			OneId:   one1.Id,
			ThreeId: three1.Id,
		}
		if err := tdb.Create(&two1).Error; err != nil {
			t.Errorf("mysql: create two1 failed: %v", err)
		}
		So(one1.Name, ShouldEqual, "one1-name")
		So(one1.Description, ShouldEqual, "one1-description")
		So(three1.Key, ShouldEqual, "three1-key")
		So(three1.Count, ShouldEqual, 1)
		So(two1.OneId, ShouldEqual, one1.Id)
		So(two1.ThreeId, ShouldEqual, three1.Id)

		// clean up
		if err := DropDatabase(db, "tdb"); err != nil {
			t.Error(err)
		}
	})
	Convey("pg: model and relatedModel", t, func() {
		type One struct {
			Model
			Name        string `json:"name"`
			Description string `json:"description"`
		}

		type Two struct {
			RelateModel
			OneId   int `json:"oneId" gorm:"index"`
			ThreeId int `json:"threeId" gorm:"index"`
		}

		type Three struct {
			Model
			Key   string `json:"key"`
			Count int    `json:"count"`
		}
		p := NewPg(nil)
		db := p.Open()
		if err := CreateDatabase(db, "tdb"); err != nil {
			t.Error(err)
		}
		tdb := p.Open("tdb")
		if err := tdb.AutoMigrate(new(One), new(Two), new(Three)); err != nil {
			t.Errorf("pg: automigrate failed: %v", err)
		}
		one1 := One{
			Name:        "one1-name",
			Description: "one1-description",
		}
		if err := tdb.Create(&one1).Error; err != nil {
			t.Errorf("pg: create one1 failed: %v", err)
		}
		three1 := Three{
			Key:   "three1-key",
			Count: 1,
		}
		if err := tdb.Create(&three1).Error; err != nil {
			t.Errorf("pg: create three1 failed: %v", err)
		}
		two1 := Two{
			OneId:   one1.Id,
			ThreeId: three1.Id,
		}
		if err := tdb.Create(&two1).Error; err != nil {
			t.Errorf("pg: create two1 failed: %v", err)
		}
		So(one1.Name, ShouldEqual, "one1-name")
		So(one1.Description, ShouldEqual, "one1-description")
		So(three1.Key, ShouldEqual, "three1-key")
		So(three1.Count, ShouldEqual, 1)
		So(two1.OneId, ShouldEqual, one1.Id)
		So(two1.ThreeId, ShouldEqual, three1.Id)

		// clean up
		if err := DropDatabase(db, "tdb"); err != nil {
			t.Error(err)
		}
	})
}
