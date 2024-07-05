import (
	"encoding/json"
	"fmt"
	"log"
)

type SaveLocation struct {
	Name            string   `json:"name"`
	SourceDirectory string   `json:"sourceDirectory"`
	Include         []string `json:"include,omitempty"` // Use omitempty to exclude empty slices
}

type Game struct {
	SteamAppID    int            `json:"steamAppId,omitempty"`
	DirectoryName string         `json:"directoryName,omitempty"` // Use omitempty to exclude empty fields
	Alias         string         `json:"alias,omitempty"`         // Use omitempty to exclude empty fields
	SaveLocations []SaveLocation `json:"saveLocations"`
}

type GameLibrary struct {
	Games []Game `json:"games"`
}

func LoadGameLibrary(filename string) (GameLibrary, error) {
	var library GameLibrary

	// Read the file
	data, err := os.ReadFile(filename)
	if err != nil {
		return library, fmt.Errorf("failed to read file: %v", err)
	}

	// Unmarshal JSON data into the struct
	err = json.Unmarshal(data, &library)
	if err != nil {
		return library, fmt.Errorf("error unmarshaling JSON: %v", err)
	}

	return library, nil
}
