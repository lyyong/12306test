package main

import (
	"flag"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"strconv"
)

type Train struct {
	ID       uint   `gorm:"primary_key;AUTO_INCREMENT"`
	Number   string `gorm:"size:20"`
	SeatSum  uint   // 座位总数
	SeatASum uint   // A座位总数
	SeatBSum uint
	SeatCSum uint
}

type Station struct {
	ID   uint   `gorm:"primary_key;AUTO_INCREMENT"`
	Name string `gorm:"size:100"`
}

type StopInfo struct {
	ID        uint  `gorm:"primary_key;AUTO_INCREMENT"`
	Train     Train `gorm:"ForeignKey:TrainID"`
	TrainID   uint
	Station   Station `gorm:"ForeignKey:StationID"`
	StationID uint
	Seq       uint // 到达的顺序
}

func createData(db *gorm.DB) {
	train := &Train{
		Number:   "A01",
		SeatSum:  1500,
		SeatASum: 500,
		SeatBSum: 500,
		SeatCSum: 500,
	}
	//db.NewRecord(train)
	if err := db.Create(train).Error; err != nil {
		fmt.Println(err)
		return
	}

	station := &Station{
		Name: "00",
	}

	if err := db.Create(station).Error; err != nil {
		fmt.Println(err)
		return
	}
	db.Create(&StopInfo{
		Train:     *train,
		TrainID:   train.ID,
		Station:   *station,
		StationID: station.ID,
		Seq:       0,
	})

	for i, j := 1, 1; i < 60; {
		name := strconv.FormatInt(int64(i), 10)
		if i < 10 {
			name = "0" + name
		}
		station = &Station{
			Name: name,
		}
		db.Create(station)
		db.Create(&StopInfo{
			Train:     *train,
			TrainID:   train.ID,
			Station:   *station,
			StationID: station.ID,
			Seq:       uint(j),
		})

		i++
		j++
	}
}

var (
	username = flag.String("username", "GoStudy", "mysql username")
	password = flag.String("password", "996282116", "mysql password")
	dbname   = flag.String("dbname", "12306_test", "mysql dbname")
	host     = flag.String("host", "localhost:3306", "mysql host")
)

func main() {
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		*username,
		*password,
		*host,
		*dbname))
	if err != nil {
		fmt.Println(err)
		return
	}
	createData(db)
	defer db.Close()
}
