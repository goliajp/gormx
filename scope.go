package gormx

import "gorm.io/gorm"

type Scope = func(db *gorm.DB) *gorm.DB

func ScopeOrderByCreatedAtDesc(db *gorm.DB) *gorm.DB {
	return db.Order("created_at DESC")
}

func ScopeOrderByUpdatedAtDesc(db *gorm.DB) *gorm.DB {
	return db.Order("updated_at DESC")
}

func ScopePagination(page, pageSize int) Scope {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset((page - 1) * pageSize).Limit(pageSize)
	}
}
