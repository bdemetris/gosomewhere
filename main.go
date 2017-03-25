package main

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
	"strings"
	"github.com/bdemetris/gosomewhere/weatherdata"
	"github.com/bdemetris/gosomewhere/campgrounddata"
)

func main() {
	http.HandleFunc("/hello", hello)

	http.HandleFunc("/weather/", func(w http.ResponseWriter, r *http.Request) {
		city := strings.SplitN(r.URL.Path, "/", 3)[2]

		data, err := weatherdata.QueryWeather(city)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(w).Encode(data)
	})

	http.HandleFunc("/amenity/", func(w http.ResponseWriter, r *http.Request) {
			amenityID := strings.SplitN(r.URL.Path, "/", 3)[2]

			data, err := campgrounddata.QueryCampground(amenityID)
			if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
			}

			w.Header().Set("Content-Type", "application/xml; charset=utf-8")
			xml.NewEncoder(w).Encode(data)
	})

	http.ListenAndServe(":8080", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello!"))
}
