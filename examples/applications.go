package main

import (
	"log"
)

func applicationsExample() {
	git := gitlab.NewClient(nil, "yourtokengoeshere")
	git.SetBaseURL("https://gitlab.com/api/v4")

	applications, _, err := git.Applications.ListApplications()

	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Found %d applications :", len(applications))
	for _, app := range applications {
		log.Printf("Found app : %v", app)
	}
}
