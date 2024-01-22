package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type weatherData struct {
	Name    string `json:"name"`
	Current struct {
		Temp float64 `json:"temperature_2m"`
	} `json:"current"`
}

type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type ApiResponse struct {
	Results []Location `json:"results"`
}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/weather/",
		func(w http.ResponseWriter, r *http.Request) {
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
	w.Write([]byte("Hello from go !"))
}
func query(city string) (weatherData, error) {

	coordinates := getLocation(city)
	fmt.Printf("Coordinates %v\n", coordinates)

	lat := fmt.Sprintf("%f", coordinates.Latitude)
	long := fmt.Sprintf("%f", coordinates.Longitude)
	res, err := http.Get("https://api.open-meteo.com/v1/forecast?latitude=" + lat + "&longitude=" + long + "&current=temperature_2m")
	if err != nil {
		return weatherData{}, err
	}

	defer res.Body.Close()

	var d weatherData
	if err := json.NewDecoder(res.Body).Decode(&d); err != nil {
		return weatherData{}, err
	}

	d.Name = city
	return d, nil
}

func getLocation(city string) Location {
	res, err := http.Get("https://geocoding-api.open-meteo.com/v1/search?name=" + city + "&count=1&language=en&format=json")
	if err != nil {
		return Location{}
	}

	defer res.Body.Close()

	var apiResponse ApiResponse
	err = json.NewDecoder(res.Body).Decode(&apiResponse)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return Location{}
	}
	return apiResponse.Results[0]
}
