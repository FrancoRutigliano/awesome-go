package main

import (
	"acl-pattern/internal/acl"
	"acl-pattern/internal/external"
	"log"
)

func main() {
	temp, _ := external.GetWeather("Argentina")
	log.Println("Temperature in Kelvin:", temp)

	tempC, _ := acl.KelvinToCelsius("Argentina")
	log.Println("Temperature in Celsius:", tempC)

	tempF, _ := acl.KelvinToFahrenheit("Argentina")
	log.Println("Temperature in Fahrenheit:", tempF)
}
