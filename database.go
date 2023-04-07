package gormx

import (
	"fmt"
	"gorm.io/gorm"
)

func CreateDatabase(db *gorm.DB, dbname string) error {
	switch db.Dialector.Name() {
	case "mysql":
		if err := db.Exec("CREATE DATABASE IF NOT EXISTS " + dbname + " DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci").Error; err != nil {
			return fmt.Errorf("create database failed: %v", err)
		}
		return nil
	case "postgres":
		has, err := HasDatabase(db, dbname)
		if err != nil {
			return err
		}
		if has {
			return nil
		}
		if err := db.Exec("CREATE DATABASE " + dbname).Error; err != nil {
			return fmt.Errorf("create database failed: %v", err)
		}
		return nil
	default:
		return fmt.Errorf("unsupported database %s", db.Dialector.Name())
	}
}

func DropDatabase(db *gorm.DB, dbname string) error {
	switch db.Dialector.Name() {
	case "mysql":
		if err := db.Exec("DROP DATABASE IF EXISTS " + dbname).Error; err != nil {
			return fmt.Errorf("drop database failed: %v", err)
		}
		return nil
	case "postgres":
		if err := db.Exec("DROP DATABASE IF EXISTS " + dbname + " WITH (FORCE)").Error; err != nil {
			return fmt.Errorf("drop database failed: %v", err)
		}
		return nil
	default:
		return fmt.Errorf("unsupported database %s", db.Dialector.Name())
	}
}

func HasDatabase(db *gorm.DB, dbname string) (bool, error) {
	var count int64
	switch db.Dialector.Name() {
	case "mysql":
		if err := db.Raw("SELECT COUNT(*) FROM information_schema.schemata WHERE schema_name = ?", dbname).
			Scan(&count).Error; err != nil {
			return false, fmt.Errorf("check database exist failed: %v", err)
		}
		return count > 0, nil
	case "postgres":
		if err := db.Raw("SELECT COUNT(*) FROM pg_database WHERE datname = ?", dbname).
			Scan(&count).Error; err != nil {
			return false, fmt.Errorf("check database exist failed: %v", err)
		}
		return count > 0, nil
	default:
		return false, fmt.Errorf("unsupported database %s", db.Dialector.Name())
	}
}
