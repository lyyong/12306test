package model

type Station struct {
	ID   int    `gorm:"primary_key;AUTO_INCREMENT"`
	Name string `gorm:"size:100"`
}

func GetAllStations() []*Station {
	var stations []*Station
	db.Find(&stations)
	return stations
}
