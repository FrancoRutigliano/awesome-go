package external

import "math/rand/v2"

func GetWeather(country string) (float64, error) {
	// simulate call an api
	// the api response me with the temperature in kelvin
	return 273 + rand.Float64()*50, nil
}
