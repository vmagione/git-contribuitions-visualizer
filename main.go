package main

import (
	"flag"

	"git-contribuitions-visualizer/visualizer"
)

func main() {
	var folder string
	var email string

	flag.StringVar(&folder, "add", "", "add a new folder to scan for Git repositories")
	flag.StringVar(&email, "email", "your@email.com", "the email to scan")

	flag.Parse()

	if folder != "" {
		visualizer.Scan(folder)
		return
	}

	visualizer.Stats(email)

}
