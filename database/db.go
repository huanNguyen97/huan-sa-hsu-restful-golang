package database

import (
	// Go importing module
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Connect db
func Db_Connection() *gorm.DB {
	dsn := "b5ccecc68435b4:715fc6a8@tcp(us-cdbr-east-04.cleardb.com)/heroku_ce6f8f81c233b36"
	
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("?")
	}
	
	return db
}