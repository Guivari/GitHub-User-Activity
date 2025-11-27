package main

import (
	"github.com/Guivari/GitHub-User-Activity/githubapi"
	"github.com/Guivari/GitHub-User-Activity/githubevents"
	"fmt"
	"os"
)

func main() {
	fmt.Println("This will look for github user activities")
	githubUser := os.Args[1]
	
	req := githubapi.MakeRequest(githubUser)
	resp := githubapi.SendRequest(req)
	defer resp.Body.Close()
	events := githubapi.HandleResponse(resp)

	eventCount := make(map[string]map[string]int)

	githubevents.Tally(&eventCount, events)
	githubevents.Print(eventCount)
}
