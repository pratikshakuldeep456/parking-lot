package parkinglot

// level represent a floor of parkinglot
type Level struct {
	ParkingSpots   []*ParkingSpot
	Floor          int
	AvailableSpots int
}

func (l *Level) Park(v *Vehicle) bool {
	for _, spot := range l.ParkingSpots {
		if spot.IsAvailable() {
			spot.ParkVehicle(v)
			l.AvailableSpots--
			return true
		}

	}
	return false
}

func (l *Level) Unpark(v *Vehicle) {
	for _, spot := range l.ParkingSpots {
		if spot.ISOccupied {
			spot.UnparkVehicle(v)
		}
	}
}
