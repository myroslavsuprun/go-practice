package main

import (
	"cyoa/server"
	"cyoa/stories"
	"cyoa/templates"
	"fmt"
)

func main() {
	stories, err := stories.Get()
	if err != nil {
		fmt.Printf("Error getting stories: %s\n", err)
		panic(err)
	}

	tmpl, err := templates.Get()
	if err != nil {
		fmt.Printf("Error getting templates: %s\n", err)
		panic(err)
	}

	err = server.Init(tmpl, stories)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
		panic(err)
	}
}
