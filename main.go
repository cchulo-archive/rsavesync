package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"rsavesync/exec"
	"rsavesync/logger"
	"rsavesync/parse"
	"strings"
)

func printVersion() {
	fmt.Println("rsavesync version 0.5.0")
}

func main() {

	versionFlag := flag.Bool("version", false, "Print the version number")
	aliasFlag := flag.String("alias", "", "Specify an alias for a non-steam game")
	gameSettings := flag.String("settings", "default-game-settings.json", "Specify a game settings file")

	flag.Usage = func() {
		_, err := fmt.Fprintf(os.Stderr, "Usage: %s [options] [arguments]\n", os.Args[0])
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Failed to print usage information: %v\n", err)
			os.Exit(1)
		}
		printVersion()
		flag.PrintDefaults()
	}

	flag.Parse()

	if len(os.Args) == 1 {
		flag.Usage()
		os.Exit(0)
	}

	if *versionFlag {
		printVersion()
		os.Exit(0)
	}

	rssLogger, logFile, err := logger.InitLogger()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Failed to open log file: %v\n", err)
		os.Exit(1)
	}

	defer func(logFile *os.File) {
		err := logFile.Close()
		if err != nil {
			rssLogger.Fatalf("Failed to close log file: %v\n", err)
		}
	}(logFile)

	multiWriter := io.MultiWriter(os.Stdout, logFile)
	rssLogger.SetOutput(multiWriter)
	posArgs := flag.Args()
	combinedArgs := strings.Join(posArgs, " ")

	if len(posArgs) > 0 {
		settings, err := parse.LoadGameSettings(*gameSettings)

		if err != nil {
			rssLogger.Fatalf("Failed to load settings json: %v", err)
		}

		steamAppId := exec.GetEnvVarOrDefault("SteamAppId", 0)

		if steamAppId == 0 && *aliasFlag == "" {
			rssLogger.Fatalf("An alias must be specified if loading a non-steam game")
		}

		game, err := settings.FindGameByAliasOrID(*aliasFlag, steamAppId)
		if err != nil {
			rssLogger.Fatalf("Failed to find game entry: %+v", err)
		}

		prettyJSON, err := json.MarshalIndent(game, "", "  ")
		if err != nil {
			log.Fatalf("Failed to marshal game library to JSON: %v", err)
		}

		rssLogger.Printf("Found game entry:\n%s", string(prettyJSON))

		runErr := exec.RunCommandWithEnv(combinedArgs, rssLogger)
		if runErr != nil {
			rssLogger.Fatalf("Error executing %s\nError: %v", combinedArgs, runErr)
		}

	} else {
		rssLogger.Println("No positional arguments provided")
	}
}
