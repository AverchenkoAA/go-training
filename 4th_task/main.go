package main

import (
	"4th_task/weather"
	"fmt"
)

func main() {
	var meteo weather.Meteorologist
	var wfc weather.WeatherForecast

	//Mahilyow, Rome, Madrid, Paris, London, Brasilia, Calgary, Sydney, Tokio, Moscow, Oslo, Murmansk
	var cities = []string{
		"Mahilyow",
		"Rome",
		"Madrid",
		"Paris",
		"London",
		"Brasilia",
		"Calgary",
		"Sydney",
		"Tokio",
		"Moscow",
		"Oslo",
		"Murmansk",
	}

	for _, city := range cities {
		weath := meteo.GetWeather(city)
		fmt.Println("\n" + wfc.FormatWeather(weath))
	}
}
