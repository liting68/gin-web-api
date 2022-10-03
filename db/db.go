package db

import (
	"app/config"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DB 数据库实体
var DB *gorm.DB

func init() {
	DB = getDb(config.Info)
}

// CloseDB 关闭连接释放连接池
func CloseDB() {
	DB.Close()
}

func getDb(conf config.Config) *gorm.DB {
	var db *gorm.DB
	var err error
	mysql := conf.DB.Mysql
	user, pass, host, port, dbname := mysql.User, mysql.Pass, mysql.Host, mysql.Port, mysql.DBName
	db, err = gorm.Open("mysql", user+":"+pass+"@tcp("+host+":"+port+")/"+dbname+"?loc=Asia%2FShanghai&parseTime=true")
	if err != nil {
		log.Panicln("err:", err.Error())
	}
	db.SingularTable(true)
	db.LogMode(conf.DB.Mysql.Debug)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(30 * time.Minute)
	return db
}
