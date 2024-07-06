package parse

import (
	"encoding/json"
	"fmt"
	"os"
)

type SaveLocation struct {
	Name            string   `json:"name"`
	SourceDirectory string   `json:"sourceDirectory"`
	Include         []string `json:"include,omitempty"`
}

type Game struct {
	SteamAppID    int            `json:"steamAppId,omitempty"`
	DirectoryName string         `json:"directoryName,omitempty"`
	Alias         string         `json:"alias,omitempty"`
	SaveLocations []SaveLocation `json:"saveLocations"`
}

type Settings struct {
	Games []Game `json:"games"`
}

func LoadSettings(filename string) (Settings, error) {
	var library Settings

	data, err := os.ReadFile(filename)
	if err != nil {
		return library, fmt.Errorf("failed to read file: %v", err)
	}

	err = json.Unmarshal(data, &library)
	if err != nil {
		return library, fmt.Errorf("error unmarshaling JSON: %v", err)
	}

	return library, nil
}

func (library *Settings) FindGameByAliasOrID(alias string, steamAppID int) (*Game, error) {
	for _, game := range library.Games {
		if game.Alias == alias || game.SteamAppID == steamAppID {
			return &game, nil
		}
	}
	return nil, fmt.Errorf("game not found with alias: %s or steamAppID: %d", alias, steamAppID)
}
