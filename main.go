package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	log "github.com/sirupsen/logrus"
)

const (
	checkPeriod = time.Second * 5
)

var (
	token     string
	projectID int
	url       string
	branch    string
	virtual   bool
	verbose   bool
)

func init() {
	flag.IntVar(&projectID, "p", 0, "GitLab project ID")
	flag.StringVar(&url, "url", "https://gitlab.com", "GitLab server URL")
	flag.StringVar(&branch, "branch", "master", "GitLab project branch")
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

	light := NewFakeLight()
	err := light.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer light.Close()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	defer func() {
		signal.Stop(c)
	}()

	ticker := time.NewTicker(checkPeriod)
	defer ticker.Stop()

	branchURL := fmt.Sprintf("%s/api/v4/projects/%d/repository/commits/%s?private_token=%s", url, projectID, branch, token)

	for {
		select {
		case <-ticker.C:
			log.WithField("URL", branchURL).Debug("Fetching")
			resp, err := http.Get(branchURL)
			if err != nil {
				log.Warn(err)
				continue
			}
			defer resp.Body.Close()

			var build struct {
				Status string `json:"status"`
			}
			if err := json.NewDecoder(resp.Body).Decode(&build); err != nil {
				log.Warn(err)
				continue
			}
			log.WithField("object", fmt.Sprintf("%#v", build)).Debug("Decoded")
			log.WithField("status", build.Status).Info()
			switch build.Status {
			case "success":
				light.Toggle(Green)
			case "canceled":
				fallthrough
			default:
				light.Toggle(All)
			}
		case <-c:
			return
		}
	}
}
