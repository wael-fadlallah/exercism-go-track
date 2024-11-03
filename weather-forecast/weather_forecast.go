// Package weather provides a library for obtaining weather forecast information.
package weather

// CurrentCondition is a variable that holds the current weather condition.
var CurrentCondition string

// CurrentLocation is a variable that holds the current location.
var CurrentLocation string

// Forecast returns a weather forecast for a specific city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
