package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func GetWeather( lat float64, lon float64) (string, Weather) {
	latStr := strconv.FormatFloat(lat, 'f', 2, 64)
	lonStr := strconv.FormatFloat(lon, 'f', 2, 64)
	req, err := http.NewRequest("GET", "https://api.weather.yandex.ru/v1/forecast?"+"lat="+latStr+"&"+"lon="+lonStr+"&extra=true", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("X-Yandex-API-Key", "6a653901-d939-47c7-8868-db449fd6a7df")
	client := &http.Client{}
	result, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	var resultInt Weather
	json.NewDecoder(result.Body).Decode(&resultInt)
	var temper string
	if resultInt.Fact.Temp > 0 {
		temper = "+" + strconv.Itoa(int(resultInt.Fact.Temp))
	} else {
		temper = strconv.Itoa(int(resultInt.Fact.Temp))
	}

	return temper, resultInt
}

func GetTemperature(lat, lon float64) string {
	var weather Weather
	_, weather = GetWeather(lat, lon)
	var temperature string
	if weather.Fact.Temp > 0 {
		temperature = "+" + strconv.Itoa(int(weather.Fact.Temp))
	} else {
		temperature = strconv.Itoa(int(weather.Fact.Temp))
	}
	return temperature
}