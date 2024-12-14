package main

import (
	"fmt"
	"groupie_tracker/database"
	"groupie_tracker/models"
	"groupie_tracker/routes"
	"os"
	"sync"
)

func main() {
	if len(os.Args) != 1 {
		os.Exit(1)
	}
	var err error
	channnels := make(chan error, 3)
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		wg.Done()
		database.Artists, err = models.GetArtists()
		if err != nil {
			channnels <- err
		}
	}()
	go func() {
		wg.Done()
		database.Locations, err = models.GetLocation()
		if err != nil {

			channnels <- err
		}

	}()
	go func() {
		wg.Done()
		database.Dates, err = models.GetDate()
		if err != nil {
			channnels <- err
		}
	}()

	go func() {
		wg.Wait()
		close(channnels)
	}()
	for err := range channnels {
		if err != nil {
			fmt.Println(err)
		}
	}
	routes.Router()
}
