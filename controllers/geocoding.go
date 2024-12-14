package controllers

import (
	"encoding/json"
	"groupie_tracker/models"
	"net/http"
	"strconv"
)

func GeocodingController(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		ErrorController(w, r, http.StatusMethodNotAllowed)
		return
	}

	// Parse artist ID from query parameter
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ErrorController(w, r, http.StatusBadRequest)
		return
	}

	// Fetch artist data
	artist, err := models.ArtistById(id)
	if err != nil {
		ErrorController(w, r, http.StatusNotFound)
		return
	}

	// Fetch geocoding data
	var locations models.Location
	if err := fetchData(artist.Locations, &locations); err != nil {
		ErrorController(w, r, http.StatusInternalServerError)
		return
	}

	geoLocations := make(map[string][2]float64)
	for _, location := range locations.Location {
		coords, err := models.GeocodeLocation(location)
		if err != nil {
			ErrorController(w, r, http.StatusInternalServerError)
			return
		}
		geoLocations[location] = coords
	}
	// Send geolocation data as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(geoLocations)
}
