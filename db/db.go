package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func init() {
	initDB()
}

// DB 数据库链接
var db *gorm.DB

func GetDB() *gorm.DB {
	return db
}

func Close() {
	db.Close()
}

func initDB() {
	database, err := gorm.Open("mysql", "zhangpeng:Aa123456.@(140.143.143.77)/service?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "tbl_" + defaultTableName
	}
	//禁用表名复数
	database.SingularTable(true)
	db = database
}
