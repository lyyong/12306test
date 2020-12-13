package model

import "fmt"

type StopInfo struct {
	ID        int `gorm:"primary_key;AUTO_INCREMENT"`
	Train     Train
	TrainID   int
	Station   Station
	StationID int
	Seq       int // 到达的顺序
}

func GetStopInfos(maps interface{}) []*StopInfo {
	var stopInfos []*StopInfo
	if err := db.Preload("Station").Where(maps).Find(&stopInfos).Error; err != nil {
		fmt.Println(err)
		return nil
	}
	return stopInfos
}
