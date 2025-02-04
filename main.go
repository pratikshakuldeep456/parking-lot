package main

import (
	parkinglot "pratikshakuldeep456/parking-lot/pkg/parkingLot"
)

func main() {

	car1 := &parkinglot.Vehicle{Vehicle: parkinglot.Car, LicenceNo: "ABC01"}
	//bike1 := &parkinglot.Vehicle{Vehicle: parkinglot.Motorcycle, LicenceNo: "ABC02"}
	pl := parkinglot.GetInstance()
	pl.Park(car1)
	pl.Unpark(car1)
}
