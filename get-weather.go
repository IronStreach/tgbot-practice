package main

import (
	"encoding/json"
	"log"
	"net/http"
)


func GetWeather(url string, key string) Weather {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("X-Yandex-API-Key", key)
	client := &http.Client{}
	result, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	var resultInt Weather
	json.NewDecoder(result.Body).Decode(&resultInt)
	return resultInt
}