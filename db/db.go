package db

import (
	"app/config"
	"log"
	"time"

	"github.com/go-redis/redis/v7"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//DB 数据库实体
var DB *gorm.DB

//Redis 缓存DB
var Redis *redis.Client

func init() {
	DB = getDb(config.Info)
	Redis = redis.NewClient(&redis.Options{
		Addr: config.Info.DB.Redis.Host + ":" + config.Info.DB.Redis.Port,
	})
}

//CloseDB 关闭连接释放连接池
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
	db.LogMode(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(30 * time.Minute)
	return db
}
