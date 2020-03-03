package main

import "time"

type Weather struct {
	Now   int       `json:"now"`
	NowDt time.Time `json:"now_dt"`
	Info  struct {
		Lat        float64 `json:"lat"`
		Lon        float64 `json:"lon"`
		Tzinfo     TZinfo  `json:"tzinfo"`
		PressureMm float64 `json:"def_pressure_mm"`
		PressurePa float64 `json:"def_pressure_pa"`
		Url        string  `json:"url"`
	} `json:"info"`
	Fact struct {
		Temp         int64   `json:"temp"`
		Feels        float64 `json:"feels_like"`
		WaterTemp    float64 `json:"temp_water"`
		Icon         string  `json:"icon"`
		Condition    string  `json:"condition"`
		WindSpeed    float64 `json:"wind_speed"`
		WindGust     float64 `json:"wind_gust"`
		WindDir      string  `json:"wind_dir"`
		PressureMm   float64 `json:"pressure_mm"`
		PressurePa   float64 `json:"pressure_pa"`
		Humidity     float64 `json:"humidity"`
		DayTime      string  `json:"daytime"`
		Polar        bool    `json:"polar"`
		Season       string  `json:"season"`
		ObsTime      int64   `json:"obs_time"`
		PrecType     int     `json:"prec_type"`
		PrecStrength float64 `json:"prec_strength"`
		Cloudness    float64 `json:"cloudness"`
	} `json:"fact"`
	Forecasts interface{} `json:"forecasts"`
}

type TZinfo struct {
	Offset int    `json:"offset"`
	Name   string `json:"name"`
	Abbr   string `json:"abbr, omitempty"`
	Dst    bool   `json:"dst"`
}
