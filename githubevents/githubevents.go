package githubevents

import (
	"fmt"
)

func Tally(eventCount *map[string]map[string]int, events []interface{}) {	
	for _, e := range events {
		event, _ := e.(map[string]interface{})
		eventName := event["type"].(string)
    repo := event["repo"].(map[string]interface{})
    repoName := repo["name"].(string)
		if (*eventCount)[repoName] == nil {
			(*eventCount)[repoName] = make(map[string]int)
		}
		(*eventCount)[repoName][eventName] += 1
	}
}

func Print(eventCount map[string]map[string]int) {
	for repo,event := range eventCount {
		fmt.Println()
		fmt.Println("In " + repo + ":")
		for thisEvent, count := range event {
			fmt.Printf("%s %v times\n", thisEvent, count)
		}
	}
}