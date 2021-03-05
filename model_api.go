package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func Init() {
	DB, _ = gorm.Open("mysql", "root:Jinjun123!@tcp(49.234.129.151:3306)/coupon?charset=utf8&parseTime=True&loc=Local")
}

func Close() {
	DB.Close()
}

func Renew() {
	Init()
	defer Close()
	DB.Set("gorm:table_options", "ENGINE=InnoDB")

	DB.DropTableIfExists(&User{}, &Refund{}, &Ship{}, &Cat{}, &Item{}, &Trade{}, &Order{})
	DB.CreateTable(&User{}, &Refund{}, &Ship{}, &Cat{}, &Item{}, &Trade{}, &Order{})
	DB.Create(&Refund{RefundID: "101502290600487397", Status: "NEWHERE"})
	DB.Create(&User{Openid: "ocpij4ivTmwb1o_IrWbp912Y4cRE", WeChat: "fg2060", BuyerNick: "印霜设计"})
}

func Migrate() {
	Init()
	defer Close()
	DB.Set("gorm:table_options", "ENGINE=InnoDB")
	DB.AutoMigrate(&User{}, &Refund{}, &Ship{}, &Cat{}, &Item{}, &Trade{}, &Order{})
}
