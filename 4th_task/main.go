package main

import (
	"4th_task/weather"
	"fmt"
)

func main() {
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
	var meteo weather.Meteorologist
	for _,city:=range cities{
		fmt.Println("\n"+meteo.WF.FormatWeather(meteo.GetWeather(city)))
	}
}
