package main

import (
	"flag"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
)

const (
	checkPeriod = time.Second * 5
)

var (
	token     string
	projectID int
	virtual   bool
	verbose   bool
)

func init() {
	flag.IntVar(&projectID, "p", 0, "GitLab project ID")
	flag.BoolVar(&virtual, "virtual", false, "enable virtual light")
	flag.BoolVar(&verbose, "verbose", false, "enable verbose logging")
	flag.Parse()

	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})

	if verbose {
		log.SetLevel(log.DebugLevel)
		log.Debug("Setting verbose logging")
	}

	if virtual {
		log.Debug("Setting virtual light")
	}

	if projectID == 0 {
		log.Fatal("Valid Project ID is required")
	}

	log.Info("Setting project ID: ", projectID)

	token = os.Getenv("GITLAB_ACCESS_TOKEN")
	if token == "" {
		log.Fatal("GITLAB_ACCESS_TOKEN environment variable is not set")
	}
}
func main() {
	log.Info("buildy starting...")
}
