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

	fmt.Println("%s", events)
}
