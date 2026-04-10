//Package weather is a package that forecasts the current weather condition of a location.
package weather

var (
    // CurrentCondition represents the situation of the weather at a particular time.
	CurrentCondition string
    // CurrentLocation represents the place or position that experiences a particular weather condition.
	CurrentLocation  string
)
// Forecast recieves the location and weather condition and outputs the refined string of the current location and condition of the location at a particular time.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
