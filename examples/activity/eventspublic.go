package main

import (
	"fmt"
	"github.com/google/go-github/github"
)

func main() {
	client := github.NewClient(nil)

	opt := &github.ListOptions{Page: 2}
	events, _, err := client.Activity.ListEvents(opt)
	if err != nil {
		fmt.Println("Activities.ListEvents returned error: %v", err)
	}

	fmt.Printf("Number of events = %v\n", len(events))

	for i, v := range events {
		fmt.Printf("array value at [%d]=%s\n", i, *v.Type)
	}

	rate, _, err := client.RateLimit()
	if err != nil {
		fmt.Printf("Error fetching rate limit: %#v\n\n", err)
	} else {
		fmt.Printf("API Rate Limit: %d\n", rate.Remaining)
	}

}
