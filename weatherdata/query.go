package weatherdata

import (
  "encoding/json"
  "net/http"
  "fmt"
)

func Query(city string) (weatherData, error) {
	apikey := LoadConfig("./weatherdata/config.json")

	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?APPID=%s&q=%s", apikey.Key, city)

	resp, err := http.Get(url)
	if err != nil {
		return weatherData{}, err
	}

	defer resp.Body.Close()

	var d weatherData

	if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
		return weatherData{}, err
	}

	return d, nil
}

type weatherData struct {
	Name string `json:"name"`
	Main struct {
		Kelvin float64 `json:"temp"`
	} `json:"main"`
}
