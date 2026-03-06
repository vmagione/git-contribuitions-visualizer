package main

import (
	"flag"

	"git-contribuitions-visualizer/visualizer"
)

func main() {
	var folder string
	var email string

	flag.StringVar(&folder, "add", "", "add a new folder to scan for Git repositories")
	flag.StringVar(&email, "email", "", "the email to scan (leave empty to include all authors)")

	flag.Parse()

	if folder != "" {
		visualizer.Scan(folder)
	}

	visualizer.Stats(email)

}
