package service

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Database is access database service
func Database() *gorm.DB {
	db, err := gorm.Open("mysql", "root:WF4La223e7Ek9usG@tcp(45.76.191.21:32770)/list?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
