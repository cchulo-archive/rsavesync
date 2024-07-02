package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	logFile, err := os.OpenFile("~/.config/rsavesync/logs/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open log file: %v\n", err)
		os.Exit(1)
	}
	defer logFile.Close()

	multiWriter := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(multiWriter)

	versionFlag := flag.Bool("version", false, "Print the version number")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] [arguments]\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

	if len(os.Args) == 1 {
		flag.Usage()
		log.Println("Displayed help message")
		os.Exit(0)
	}

	if *versionFlag {
		log.Println("Displayed version number")
		os.Exit(0)
	}

	posArgs := flag.Args()
	if len(posArgs) > 0 {
		log.Printf("Positional arguments: %v\n", posArgs)
	} else {
		log.Println("No positional arguments provided")
	}
}
