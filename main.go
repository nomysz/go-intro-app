package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

const (
	apiUrl    = "https://www.7timer.info/bin/astro.php?lon=%.4f&lat=%.4f&ac=0&unit=metric&output=json&tzshift=0"
	logFormat = "CLOUDCOVER: %d, SEEING: %d, TRANSPARENCY: %d"
)

type WeatherData struct {
	Product    string       `json:"product"`
	Init       string       `json:"init"`
	Dataseries []DataSeries `json:"dataseries"`
}

type DataSeries struct {
	Timepoint    int     `json:"timepoint"`
	Cloudcover   int     `json:"cloudcover"`
	Seeing       int     `json:"seeing"`
	Transparency int     `json:"transparency"`
	LiftedIndex  int     `json:"lifted_index"`
	Rh2m         int     `json:"rh2m"`
	Wind10m      Wind10m `json:"wind10m"`
	Temp2m       int     `json:"temp2m"`
	PrecType     string  `json:"prec_type"`
}

type Wind10m struct {
	Direction string `json:"direction"`
	Speed     int    `json:"speed"`
}

func getAPIUrl(lon, lat float64) string {
	return fmt.Sprintf(apiUrl, lon, lat)
}

func logWeather(lon, lat float64) {
	resp, err := http.Get(getAPIUrl(lon, lat))
	if err != nil {
		log.Fatal("Error reponse from weather API", err.Error())
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error when reading API response", err.Error())
	}

	var data WeatherData

	if err := json.Unmarshal(body, &data); err != nil {
		log.Fatal("Error parsing JSON", err.Error())
	}

	for _, dataPoint := range data.Dataseries {
		log.Println(
			fmt.Sprintf(logFormat, dataPoint.Cloudcover, dataPoint.Seeing, dataPoint.Transparency),
		)
	}
}

func main() {
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime)) // disable date & time from log output

	lon := flag.Float64("lon", 0.0, "longitude for weather")
	lat := flag.Float64("lat", 0.0, "latitude for weather")
	flag.Parse()

	if *lon == 0.0 || *lat == 0.0 {
		flag.Usage()
		os.Exit(1)
	}

	logWeather(*lon, *lat)
}
