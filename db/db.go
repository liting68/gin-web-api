package db

import (
	"app/config"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//EPR 电子病历库
var EPR *gorm.DB

//PATDOC 医患平台库
var PATDOC *gorm.DB

func init() {
	conf := config.Info
	EPR = getDb(conf)
	// conf.DB.Mysql.DBName="doctor_patient"
	// PATDOC = getDb(conf)
}

//CloseDB 关闭连接释放连接池
func CloseDB() {
	EPR.Close()
	// db.PATDOC.Close()
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
	db.LogMode(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(30 * time.Minute)
	return db
}
