package gormx

import (
	"fmt"
	"gorm.io/gorm"
)

func DropTables(db *gorm.DB, tables ...string) error {
	for _, table := range tables {
		cmd := "DROP TABLE IF EXISTS " + table
		if err := db.Exec(cmd).Error; err != nil {
			return fmt.Errorf("drop table failed: %v", err)
		}
	}
	return nil
}

func HasTable(db *gorm.DB, table string) (bool, error) {
	var count int64
	switch db.Dialector.Name() {
	case "mysql":
		err := db.Raw("SELECT COUNT(*) FROM information_schema.tables WHERE table_schema = DATABASE() AND table_name = ?", table).
			Scan(&count).Error
		return count > 0, err
	case "postgres":
		err := db.Raw("SELECT COUNT(*) FROM information_schema.tables WHERE table_schema = 'public' AND table_name = ?", table).
			Scan(&count).Error
		return count > 0, err
	default:
		return false, fmt.Errorf("unsupported database %s", db.Dialector.Name())
	}
}
