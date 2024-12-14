package models

import (
	"encoding/json"
	"groupie_tracker/config"
	"net/http"
	"strconv"
	"strings"
)

func GeocodeLocation(location string) ([2]float64, error) {
	var result [2]float64

	loca := strings.Split(location, "-")[0]

	response, err := http.Get(config.GeoAPIURL + loca)
	if err != nil {
		return result, err
	}
	defer response.Body.Close()

	var data []struct {
		Lat string `json:"lat"`
		Lon string `json:"lon"`
	}
	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		return result, err
	}
	// If data is available, extract the coordinates
	if len(data) > 0 {
		lat, errLat := strconv.ParseFloat(data[0].Lat, 64)
		lon, errLon := strconv.ParseFloat(data[0].Lon, 64)
		if errLat != nil || errLon != nil {
			return result, err
		}
		result = [2]float64{lat, lon}
	}

	return result, nil
}
