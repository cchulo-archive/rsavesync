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

type GameSettings struct {
	Games []Game `json:"games"`
}

func LoadGameSettings(filename string) (GameSettings, error) {
	var library GameSettings

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

func (library *GameSettings) FindGameByAliasOrID(alias string, steamAppID int) (*Game, error) {
	for _, game := range library.Games {
		if (steamAppID == 0 && game.Alias == alias) || (alias == "" && game.SteamAppID == steamAppID) {
			return &game, nil
		}
	}
	if steamAppID == 0 {
		return nil, fmt.Errorf("game not found with alias %s", alias)
	}
	return nil, fmt.Errorf("game not found with steamAppID %d", steamAppID)
}
