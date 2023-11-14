package handlers

import (
	"encoding/json"
	"net/http"

	"groupie-tracker/backend/models"
)

// GetArtists returns a list of artists
func GetArtists() ([]models.Artist, error) {
	url := "https://groupietrackers.herokuapp.com/api/artists"

	// Make a GET request to the API
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// Decode the response into a slice of artists
	var artists []models.Artist
	if err := json.NewDecoder(response.Body).Decode(&artists); err != nil {
		return nil, err
	}

	return artists, nil
}
