package main

import (
	"gopkg.in/src-d/go-git.v4"
)

// stats generates a nice graph of your Git contributions

func stats(email string) {
	commits := processRepositories(email)
	printCommitsStats(commits)
}

// processRepositories given a user email, returns the
// commits made in the last 6 months

func processRepositories(email string) map[int]int {
	filePath := getDotFilePath()
	repos := parseFileLinesToSlice(filePath)
	daysInMap := daysInLastSixMonths
	commits := make(map[int]int, daysInMap)
	for i := daysInMap; i > 0; i-- {
		commits[i] = 0
	}
	for _, path := range repos {
		commits = fillCommits(email, path, commits)
	}
	return commits
}

// fillCommits given a repository found in `path`, gets the commits and
// puts them in the `commits` map, returning it when completed
func fillCommits(email string, path string, commits map[int]int) map[int]int {
	// instantiate a git repo object from path
	repo, err := git.PlainOpen(path)
	if err != nil {
		panic(err)
	}
}
