// Package weather provides tools to forecast the weather for a given city.
package weather

// CurrentCondition stores the current weather condition.
var CurrentCondition string

// CurrentLocation stores the current location.
var CurrentLocation string

// Forecast returns a string with the weather forecast for the given city and condition.
func Forecast(city, condition string) string {
    CurrentLocation = city
    CurrentCondition = condition
    return CurrentLocation + " - current weather condition: " + CurrentCondition
}
