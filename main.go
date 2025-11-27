package main

import (
	"github.com/Guivari/GitHub-User-Activity/githubapi"
	"github.com/Guivari/GitHub-User-Activity/githubevents"
	"os"
)

func main() {
	if (len(os.Args) < 2) {
		panic("empty name")
	}
	githubUser := os.Args[1]

	req := githubapi.MakeRequest(githubUser)
	resp := githubapi.SendRequest(req)
	defer resp.Body.Close()
	events := githubapi.HandleResponse(resp)

	eventCount := make(map[string]map[string]int)
	githubevents.Tally(&eventCount, events)
	githubevents.Print(eventCount)
}
