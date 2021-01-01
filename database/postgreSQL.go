package database

import "gorm.io/gorm"

var sqlDB *gorm.DB

func initPostgreSQL() {
}

func connectDB() {}

func SqlDB() *gorm.DB {
	return sqlDB
}
