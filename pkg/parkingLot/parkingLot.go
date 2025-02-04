package parkinglot

import (
	"fmt"
	"sync"
)

//	type IParkingLot interface {
//		// Park(v *Vehicle)
//		// Unpark(v *Vehicle)
//		//GetInstance() *ParkingLot
//	}
type ParkingLot struct {
	Levels []*Level
}

func NewParkingLot(level, spotsPerLevel int) *ParkingLot {

	levels := make([]*Level, level)
	for i := range levels {
		spots := make([]*ParkingSpot, spotsPerLevel)
		for j := range spots {
			spots[j] = &ParkingSpot{SpotNo: j + 1}
		}
		levels[i] = &Level{Floor: i, ParkingSpots: spots, AvailableSpots: spotsPerLevel}
	}
	return &ParkingLot{Levels: levels}
}

func (pl *ParkingLot) Park(vehicle *Vehicle) bool {
	for _, j := range pl.Levels {
		j.Park(vehicle)
		return true
	}
	return false

}

func (pl *ParkingLot) Unpark(vehicle *Vehicle) {
	for _, level := range pl.Levels {
		level.Unpark(vehicle)
	}
}

var parkingInstance *ParkingLot
var lock = &sync.Mutex{}

func GetInstance() *ParkingLot {
	if parkingInstance == nil {
		lock.Lock()
		defer lock.Unlock()

		if parkingInstance == nil {
			fmt.Println(" cREATING Instance for ParkingLot")
			parkingInstance = NewParkingLot(4, 4)

		} else {
			fmt.Println(" Instance already created for ParkingLot")
		}
	} else {
		fmt.Println(" Instance already created for ParkingLot")
	}

	return parkingInstance
}
