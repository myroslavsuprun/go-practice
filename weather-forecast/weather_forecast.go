// Package weather provides forecast by cite and conditions.
package weather

// CurrentCondition describes weather condition in the country.
var CurrentCondition string

// CurrentLocation describes the weather forecast.
var CurrentLocation string

/*
Forecast describes forecast now by city and conditions.
*/
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
