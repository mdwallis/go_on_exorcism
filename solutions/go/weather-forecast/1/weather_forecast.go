// Package weather provides a forecast for a city and condition.
package weather

// CurrentCondition is a string that represents the current weather condition for the given location.
var CurrentCondition string
// CurrentLocation is a string that represents the current location.
var CurrentLocation string

// Forecast is a function that accepts city and condition arguments and uses them as part of a formatted string to print the current weather condition. The result is returned to the caller.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
