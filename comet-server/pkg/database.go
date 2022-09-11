package pkg

import (
	"comet-server/model"
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var _db *gorm.DB

func init() {

	username := "root"  //账号
	password := "root"  //密码
	host := "127.0.0.1" //数据库地址，可以是ip或者域名
	port := 3306        //数据库端口
	Dbname := "comet"   //数据库名
	timeout := "10s"    //连接超时，10秒

	lnk := fmt.Sprintf("%s:%s@tcp(%s:%d)/", username, password, host, port)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)

	{ //对于初次运行，会创建一个初始数据库
		db, err := sql.Open("mysql", lnk)
		if err != nil {
			panic(err)
		}
		_, err = db.Exec("CREATE DATABASE" + " " + Dbname)
		if err == nil {
			fmt.Println("检测到尚未拥有数据库，已创建初始数据库")
		}
		db.Close()
	}

	var err error
	_db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}

	InitMigrate() //自动同步数据库与当前模型

	sqlDB, _ := _db.DB()
	sqlDB.SetMaxOpenConns(100) //设置数据库连接池最大连接数
	sqlDB.SetMaxIdleConns(20)  //连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭。
}

func GetDB() *gorm.DB { // 从连接池获取数据库连接的接口
	return _db
}

func InitMigrate() {
	_db.AutoMigrate(&model.User{})
	_db.AutoMigrate(&model.UserAddress{})
	_db.AutoMigrate(&model.Item{})
}
