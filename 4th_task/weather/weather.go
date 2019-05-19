package weather

type Weather struct {
	City    string      `json:"name"`
	General WeatherMain `json:"main"`
	Wind    WeatherWind `json:"wind"`
	Time WeatherTime    `json:"sys"`
	Clouds  []struct {
		Descript string `json:"description"`
	} `json:"weather"`
}

func (w *Weather) GetTemperature() (temp float64, tempMin float64, tempMax float64) {
	return w.General.Temp, w.General.TempMin, w.General.TempMax
}

func (w *Weather) GetCloudiness() (description string) {
	return w.Clouds[0].Descript
}
func (w *Weather) GetHumidity() (humidity int) {
	return w.General.Humidity
}
func (w *Weather) GetWind() (speed int, gust int, direction string) {
	return int(w.Wind.Speed),int(w.Wind.Gust), w.Wind.GetWindDirection()
}
func (w *Weather) GetTime() (sunrise int64, sunset int64) {
	return w.Time.Sunrise, w.Time.Sunset
}
func (w *Weather) GetCity() string {
	return w.City
}
