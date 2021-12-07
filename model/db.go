package model

import (
	"GO-GIN-Vue-blog/utils"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

var db *gorm.DB
var err error

func InitDb() {
	db, err = gorm.Open(utils.Db, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		utils.DbUser,
		utils.DbPassWorld,
		utils.DbHost,
		utils.DbPort,
		utils.DbName,
	))
	if err != nil {
		fmt.Println("连接数据库失败，请检查参数：", err)
	}
	// 禁用表名复数，如果设置为true，
	db.SingularTable(true)
	// 自动迁移
	db.AutoMigrate(&User{}, &Category{}, &Article{})
	sqlDB := db.DB()
	if err != nil {
		fmt.Println("获取数据库链接失败：", err)
	}
	// SetMaxIdleConns 设置空闲连接池中的最大连接数。
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置到数据库的最大打开连接数。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置连接可以重用的最长时间。
	sqlDB.SetConnMaxLifetime(10 * time.Second)
}
