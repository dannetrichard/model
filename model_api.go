package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func Renew() {
	db, _ := gorm.Open("mysql", "root:Jinjun123!@tcp(49.234.129.151:3306)/coupon?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	db.Set("gorm:table_options", "ENGINE=InnoDB")

	db.DropTableIfExists(&User{}, &Refund{}, &Ship{}, &Cat{}, &Item{}, &Trade{}, &Order{})
	db.CreateTable(&User{}, &Refund{}, &Ship{}, &Cat{}, &Item{}, &Trade{}, &Order{})
	db.Create(&Refund{RefundID: "101502290600487397", Status: "NEWHERE"})
	db.Create(&User{Openid: "ocpij4ivTmwb1o_IrWbp912Y4cRE", WeChat: "fg2060", BuyerNick: "印霜设计"})
}

func Migrate() {
	db, _ := gorm.Open("mysql", "root:Jinjun123!@tcp(49.234.129.151:3306)/coupon?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	db.Set("gorm:table_options", "ENGINE=InnoDB")
	db.AutoMigrate(&User{}, &Refund{}, &Ship{}, &Cat{}, &Item{}, &Trade{}, &Order{})
	db.Create(&Refund{RefundID: "101502290600487397", Status: "NEWHERE"})
	db.Create(&User{Openid: "ocpij4ivTmwb1o_IrWbp912Y4cRE", WeChat: "fg2060", BuyerNick: "印霜设计"})
}
