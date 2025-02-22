package acl

import "acl-pattern/internal/external"

func KelvinToFahrenheit(country string) (float64, error) {
	tempF, err := external.GetWeather(country)
	if err != nil {
		return 0, err
	}
	return tempF*9/5 - 459.67, nil
}

// KelvinToCelsius convierte Kelvin a Celsius
func KelvinToCelsius(country string) (float64, error) {
	tempK, err := external.GetWeather(country)
	if err != nil {
		return 0, err
	}
	return tempK - 273.15, nil
}
