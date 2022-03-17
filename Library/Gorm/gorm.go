package Gorm

import (
	"fmt"
	"go-study/Config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"reflect"
	"time"
)

// Mysql 本地数据库链接
var Mysql *gorm.DB

func init() {
	Mysql = connectMysql()
}

//
// connectMysql
// @Description: 默认mysql数据库连接
// @return *gorm.DB
//
func connectMysql() *gorm.DB {
	/*dsn := "root:root@(127.0.0.1:3306)/go_gin?charset=utf8mb4&parseTime=True&loc=Local"
	db, dbErr := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if dbErr != nil {
		panic(dbErr)
	}*/

	//var db *gorm.DB
	dbConf := reflect.ValueOf(Config.Configs.Web).FieldByName("DB")

	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s%s?charset=%s&parseTime=true&loc=Local",
		dbConf.FieldByName("User"),
		dbConf.FieldByName("Pwd"),
		dbConf.FieldByName("Host"),
		dbConf.FieldByName("Port"),
		dbConf.FieldByName("Prefix"),
		dbConf.FieldByName("DbName"),
		dbConf.FieldByName("Charset"),
	)

	db, dbErr := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if dbErr != nil {
		panic(dbErr)
	}

	////连接池
	// 获取通用数据库对象 sql.DB ，然后使用其提供的功能
	sqlDB, _ := db.DB()

	// SetMaxOpenConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(30 * time.Minute)

	return db

}

//todo AutoMigrate自动建表 https://blog.csdn.net/qq_39787367/article/details/112567822
