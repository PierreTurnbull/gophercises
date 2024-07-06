package main

import (
	"os"
	"strings"
)

// GetPath returns the path to the csv file. It can be overriden by passing the --path or -p flag.
func GetPath() string {
	path := "./quizzes/colors.csv"

	for i, v := range os.Args {
		isPathFlat := v == "--path" || v == "-p"

		hasNextArg := true
		if i+1 >= len(os.Args) {
			hasNextArg = false
		}

		nextArgIsFlag := false
		if hasNextArg {
			nextArgIsFlag = strings.HasPrefix(os.Args[i+1], "-")
		}

		if isPathFlat && hasNextArg && !nextArgIsFlag {
			path = os.Args[i+1]
		}
	}

	return path
}
