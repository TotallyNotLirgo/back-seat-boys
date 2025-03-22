package users

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func prepareConnection() (conn *gorm.DB) {
	conn, err := gorm.Open(
		sqlite.Open(":memory:"),
		&gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		},
	)
	if err != nil {
		panic("failed to connect database")
	}
	conn.AutoMigrate(&User{})
	return
}
