package routes

import (
	"groupie_tracker/controllers"
	"log"
	"net/http"
)

func Router() {
	http.HandleFunc("/", controllers.ArtistController)
	http.HandleFunc("/static/", controllers.StaticController)
	http.HandleFunc("/artist/{id}/", controllers.ArtistDetailController)
	http.HandleFunc("/api/geocoding", controllers.GeocodingController) // New endpoint
	log.Println("Server running on http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
