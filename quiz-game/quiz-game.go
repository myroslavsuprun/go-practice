package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

var (
	fileName  string
	timeLimit int
)

func main() {
	const (
		defaultFileName  = "problems.csv"
		defaultTimeLimit = 10
	)
	flag.StringVar(&fileName, "fileName", defaultFileName, "A .csv file with questions and answers in {question},{answer} format.")
	flag.IntVar(&timeLimit, "timeLimit", defaultTimeLimit, "Time limit for the quiz in seconds.")

	flag.Parse()

	records := getRecords(&fileName)
	correctAnswersNum := askQuestions(records, timeLimit)
	fmt.Printf("You answered %v questions correctly out of %v.\n", correctAnswersNum, len(records))
}

func getRecords(fileName *string) [][]string {
	file, err := os.Open("./assets/" + *fileName)
	if err != nil {
		log.Fatal("Error while reading the file. Most probably it doesn't exist at all.", err)
	}

	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal("Error reading records. Please provide a valid .csv file.")
	}

	return records
}

func askQuestions(records [][]string, timeLimit int) int {
	answers := make(map[string]string)
	questions := make([]string, 0, len(records))

	for _, eachrecord := range records {
		question := fmt.Sprintf("How much will be: %v?", eachrecord[0])
		questions = append(questions, question)
		answers[question] = eachrecord[1]
	}

	var answeredCorrectly int
	timeUp := make(chan bool, 1)

	go func() {
		time.Sleep(time.Duration(timeLimit) * time.Second)
		timeUp <- true
	}()

	for _, question := range questions {
		answerReceived := make(chan int)
		fmt.Println(question)

		go func() {
			var answer int
			fmt.Scanln(&answer)
			answerReceived <- answer
		}()

		correctAnswer, _ := strconv.Atoi(answers[question])

		select {
		case <-timeUp:
			fmt.Println("\nTime's up! Quiz is over.")
			return answeredCorrectly
		case <-time.After(time.Duration(timeLimit) * time.Second):
			fmt.Println("\nTime's up! Quiz is over.")
			return answeredCorrectly
		case answer := <-answerReceived:
			if answer == correctAnswer {
				answeredCorrectly++
			}
		}
	}

	return answeredCorrectly
}
