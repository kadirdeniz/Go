package carInt

import (
	"interface/Struct/car"
)

func GetCarById(cars car.Cars, carId int) car.Cars {
	for index, value := range cars {
		if index == carId {
			return value
		}
	}
	panic("Car Couldnt Found")
}
