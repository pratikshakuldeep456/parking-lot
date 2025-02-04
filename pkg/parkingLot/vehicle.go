package parkinglot

type VehicleType int

const (
	Car VehicleType = 0
	Motorcycle
	Truck
)

type Vehicle struct {
	Vehicle   VehicleType
	LicenceNo string
}
