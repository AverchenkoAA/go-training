package main

import (
	"4th_task/weather"
	"fmt"
)

func main() {
	city:="Mahilyow"
	var meteo weather.Meteorologist
	fmt.Println(meteo.WF.FormatWeather(meteo.GetWeather(city)))
}
