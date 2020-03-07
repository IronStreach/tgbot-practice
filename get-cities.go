package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	)

type City struct {
	id int64
	city_name string
	lat float64
	lon float64
}

func GetCities() []City {
	connStr := "postgres://umqpaljvbdoskz:0b12c7a86db9cbee81034340adae2e6a2643a8506c2a61e39f15abb35d488fe1@ec2-54-217-204-34.eu-west-1.compute.amazonaws.com:5432/dejmrl4g0h1816"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	rows, err := db.Query("select * from city")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	Cities := []City{}

	for rows.Next(){
		c := City{}
		err := rows.Scan(&c.id, &c.city_name, &c.lat, &c.lon)
		if err != nil{
			fmt.Println(err)
			continue
		}
		Cities = append(Cities, c)
	}
	return Cities
}

func GetCitiesName() []string {
	var cities = GetCities()
	var names = make([]string, len(cities))
	for i := 0; i < len(cities); i++ {
		names[i] = cities[i].city_name
	}
	return names
}