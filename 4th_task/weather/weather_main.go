package weather

//"main":{"temp":13,"pressure":1021,"humidity":66,"temp_min":13,"temp_max":13},
type WeatherMain struct {
	Temp     float64 `json:"temp"`
	Pressure float64 `json:"pressure"`
	Humidity int     `json:"humidity"`
	TempMin  float64 `json:"temp_min"`
	TempMax  float64 `json:"temp_max"`
}
