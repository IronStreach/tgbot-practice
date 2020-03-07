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

func GetCityName() []City {
	connStr := "user=postgres password=timofey dbname=postgres sslmode=disable"
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

