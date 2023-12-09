package repo

import "gorm.io/gorm"

func paginateScope(page, size uint) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(int((page - 1) * size)).Limit(int(size))
	}
}