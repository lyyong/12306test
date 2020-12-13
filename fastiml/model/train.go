package model

import "fmt"

type Train struct {
	ID       int    `gorm:"primary_key;AUTO_INCREMENT"`
	Number   string `gorm:"size:20"`
	SeatSum  int    // 座位总数
	SeatASum int    // A座位总数
	SeatBSum int
	SeatCSum int
}

func GetAllTrains() []*Train {
	var trains []*Train
	if err := db.Find(&trains).Error; err != nil {
		fmt.Println(err)
	}
	return trains
}
