// Package weather provides tools to forecast and report current weather conditions
// for cities in Goblinocus. It tracks location and weather state, and generates
// human-readable weather forecasts.
package weather

// CurrentCondition stores the current weather condition (e.g. "rainy", "sunny")
// for the last forecast made.
var CurrentCondition string

// CurrentLocation stores the city name (e.g. "Goblinopolis") for which
// the last weather forecast was made.
var CurrentLocation string

// Forecast takes a city name and weather condition, stores them as the current
// values, and returns a formatted weather forecast string for display.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
