package db

import (
	"app/config"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//G 全局操作实体
var G *gorm.DB

func init() {
	G = getDb("mytest")
}

func getDb(dbname string) *gorm.DB {
	var db *gorm.DB
	var err error
	user, pass, host, port := config.Info.Db.Mysql.User, config.Info.Db.Mysql.Pass, config.Info.Db.Mysql.Host, config.Info.Db.Mysql.Port
	db, err = gorm.Open("mysql", user+":"+pass+"@tcp("+host+":"+port+")/"+dbname+"?loc=Asia%2FShanghai&parseTime=true")
	if err != nil {
		log.Panicln("err:", err.Error())
	}
	db.SingularTable(true)
	db.LogMode(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(15 * time.Minute)
	return db
}
