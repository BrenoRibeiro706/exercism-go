package cars

import "fmt"

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {

	successRate = successRate / 100

	return float64(productionRate) * successRate
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {

	perHour := CalculateWorkingCarsPerHour(productionRate, successRate)

	return int(perHour) / 60
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {

	carsGroup := carsCount / 10 // dividing cars by group of 10

	resExist := carsCount % 10 //  resExist

	if carsGroup < 1 {

		carsCount := uint(carsCount) * 10000
		fmt.Println("carsCount: ", carsCount)
		return carsCount

	} else if carsGroup >= 1 {

		if resExist >= 1 {

			groupTotal := carsGroup * 95000

			resTotal := resExist * 10000

			return uint(resTotal) + uint(groupTotal)
		} else {

			return uint(carsGroup) * 95000
		}
	} else {
		return 0
	}
}
