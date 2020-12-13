package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db       *gorm.DB
	username = ""
	password = ""
	dbname   = ""
	host     = ""
)

func InitDB() {
	var err error
	db, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		username,
		password,
		host,
		dbname))
	if err != nil {
		fmt.Println(err)
		return
	}
}

func Close() {
	db.Close()
}
