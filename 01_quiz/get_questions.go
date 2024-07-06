package main

import (
	"encoding/csv"
	"log"
	"os"
)

// GetQuestions returns the questions and answers.
func GetQuestions(file *os.File) [][]string {
	readerPtr := csv.NewReader(file)

	questions, err := readerPtr.ReadAll()

	if err != nil {
		log.Fatal(err)
	}

	return questions
}
