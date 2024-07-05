package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"rsavesync/exec"
	"rsavesync/logger"
	"strings"
)

func main() {
	
	// define CLI flags
	versionFlag := flag.Bool("version", false, "Print the version number")
	// debugFlag := flag.Bool("debug", false, "Whether or not a log file is written")
	// aliasFlag := flag.String("alias", "", "Specify an alias for a non-steam game")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] [arguments]\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

	if len(os.Args) == 1 {
		flag.Usage()
		os.Exit(0)
	}

	if *versionFlag {
		fmt.Println("Version 0.5.0")
		os.Exit(0)
	}

	logger, logFile, err := logger.InitLogger()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open log file: %v\n", err)
		os.Exit(1)
	}
	defer logFile.Close()

	multiWriter := io.MultiWriter(os.Stdout, logFile)
	logger.SetOutput(multiWriter)

	posArgs := flag.Args()
	combinedArgs := strings.Join(posArgs, " ")

	if len(posArgs) > 0 {
		exec.RunCommandWithEnv(combinedArgs, logger)
	} else {
		logger.Println("No positional arguments provided")
	}
}
