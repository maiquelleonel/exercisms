// Package weather helps to work with weather forecast.
package weather

// CurrentCondition represents the current weather condition.
var CurrentCondition string

// CurrentLocation represents the current geograpfic location.
var CurrentLocation string

// Forecast function return the weather condition for a passed city and current condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
