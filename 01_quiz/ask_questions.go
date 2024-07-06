package main

import (
	"fmt"
	"time"
)

func AskQuestions(questions [][]string) {
	rightAnswers := 0
	wrongAnswers := 0

	timer := time.NewTimer(time.Duration(2 * time.Second))
	doneCh := make(chan bool)
	breakCh := make(chan bool)

	go func() {
	loop:
		for _, question := range questions {
			var answer string
			fmt.Println(question[0])
			fmt.Scanln(&answer)

			select {
			case <-breakCh:
				break loop
			default:
				if answer == question[1] {
					rightAnswers++
				} else {
					wrongAnswers++
				}
			}
		}
		doneCh <- true
	}()

	select {
	case <-doneCh:
	case <-timer.C:
		fmt.Println("Time's up! Please press enter to show the results.")
		breakCh <- true
	}

	fmt.Println("Right answers:", rightAnswers)
	fmt.Println("Wrong answers:", wrongAnswers)
	fmt.Println("Total questions:", len(questions))
}
