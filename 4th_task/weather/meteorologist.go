package weather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Meteorologist struct {
	WF WeatherForecast
}

func (m *Meteorologist) GetWeather(city string) Weather {
	url := "http://api.openweathermap.org/data/2.5/weather?q=" + city + "&lang=ru&units=metric&appid=2c19a8c670afc70f2ae7a81f229fce3d"
	var weath Weather
	err1 := json.Unmarshal(m.getBody(url), &weath)
	if err1 != nil {
		fmt.Println("error:", err1)
	}
	return weath
}

func (m *Meteorologist) getBody(url string) []byte {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	return data
}
