package parkinglot

import "fmt"

type IParkingSpot interface {
	IsAvailable()
	ParkVehicle()
	UnparkVehicle()
}
type ParkingSpot struct {
	Vehicle    *Vehicle
	SpotNo     int
	ISOccupied bool
}

func (ps *ParkingSpot) IsAvailable() bool {
	return !ps.ISOccupied
}

func (ps *ParkingSpot) ParkVehicle(v *Vehicle) {
	if ps.IsAvailable() {
		ps.Vehicle = v
		ps.ISOccupied = true

		fmt.Println("parking vehicle at the spot ", ps.SpotNo)
	}

}

func (ps *ParkingSpot) UnparkVehicle(v *Vehicle) {
	if ps.ISOccupied {
		ps.Vehicle = nil
		ps.ISOccupied = false
	}
}
