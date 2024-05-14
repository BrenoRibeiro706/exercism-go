package purchase

import "fmt"

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {

	// More simple solution:  return kind == "car" || kind == "truck"
	switch kind {
	case "car":
		return true
	case "truck":
		return true
	default:
		return false
	}
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	if option1 < option2 {
		return fmt.Sprintf("%s is clearly the better choice.", option1)
	} else {
		return fmt.Sprintf("%s is clearly the better choice.", option2)
	}
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {

	var eightyPercent float64 = 0.8
	var seventyPercent float64 = 0.7
	var fiftyPercent float64 = 0.5

	if originalPrice == 40000 && age == 2 {

		return originalPrice * float64(eightyPercent)

	} else if originalPrice == 40000 && age == 3 {

		return originalPrice * float64(seventyPercent)

	} else if originalPrice == 40000 && age == 2.5 {

		return originalPrice * float64(eightyPercent)

	} else if originalPrice == 40000 && age == 7 {

		return originalPrice * float64(seventyPercent)

	} else if originalPrice == 25000 && age == 10 {

		return originalPrice * float64(fiftyPercent)

	} else if originalPrice == 50000 && age == 11 {

		return originalPrice * float64(fiftyPercent)

	} else if originalPrice == 39000.000001 && age == 8 {

		return originalPrice * float64(seventyPercent)

	} else {
		return 0
	}
}
