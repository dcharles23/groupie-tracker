// backend/handlers/combined_data_handler.go
package handlers

import (
	"groupie-tracker/backend/models"
)

// GetArtistsWithRelations returns a combined data struct with artists and relations
func GetArtistsWithRelations() (*models.CombinedData, error) {
	artists, err := GetArtists()
	if err != nil {
		return nil, err
	}

	// Create a map of relations with artist ID as key
	relationsMap := make(map[int]*models.Relations)
	for i := range artists {
		relations, err := GetRelations(artists[i].Relations)
		if err != nil {
			return nil, err
		}
		relationsMap[artists[i].ID] = relations
	}

	// Create a combined data struct
	combinedData := &models.CombinedData{
		Artists:       artists,
		RelationsData: relationsMap,
	}

	return combinedData, nil
}
