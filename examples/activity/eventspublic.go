// Copyright 2014 The go-github AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

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

	fmt.Println("%s",events)
}
