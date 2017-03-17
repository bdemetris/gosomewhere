package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/BurntSushi/toml"
)

func main() {
	http.HandleFunc("/hello", hello)

	http.HandleFunc("/weather/", func(w http.ResponseWriter, r *http.Request) {
		city := strings.SplitN(r.URL.Path, "/", 3)[2]

		data, err := query(city)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(w).Encode(data)
	})

	http.ListenAndServe(":8080", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello!"))
}

func query(city string) (weatherData, error) {

	var config keys

	if _, err := toml.DecodeFile("config.toml", &config); err != nil {
		return
	}

	resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?APPID=ebb6c3a404230a81e1ccc19a3095a6cc&q=" + city)
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

type keys struct {
	Key string
}
