// Package weather indicates the weather package.
package weather

// CurrentCondition is a string.
var CurrentCondition string

// CurrentLocation is a string.
var CurrentLocation string

// Forecast function receive two params, city and condition, CurrentLocation is equal city and CurrentCondition is equal condition.
// And return the follow setence: "Cityj - current wather condition: Condition".
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
