package main

import (
	"os"

	"github.com/MNThomson/securitybot/events"
	"github.com/MNThomson/securitybot/ingestor"
	log "github.com/sirupsen/logrus"
)

func printUsage() {
	log.Errorf("Usage: %s [events|ingestor]\n", os.Args[0])
	os.Exit(1)
}

func main() {
	if len(os.Args) < 2 {
		printUsage()
	}
	switch os.Args[1] {
	case "events":
		events.InitEvents()
	case "ingestor":
		ingestor.Ingestor()
	default:
		printUsage()
	}
}
