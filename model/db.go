package model

import (
	"fmt"
	"ginVueBlog/utils"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var err error

func InitDb() {
	db, err := gorm.Open(utils.Db, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		utils.DbUser,
		utils.DbPassword,
		utils.DbHost,
		utils.DbPort,
		utils.DbName,
	))
	if err != nil {
		fmt.Println("connect mysql false,please check parameter:", err)
	}

	db.SingularTable(true)
	db.AutoMigrate(&User{}, &Category{}, &Article{})
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	db.DB().SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	db.DB().SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	db.DB().SetConnMaxLifetime(10 * time.Second)

	//db.Close()
}
