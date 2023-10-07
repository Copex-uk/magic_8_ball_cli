package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	// Create 3 arrays for the table
	affirmativeAnswers := []string{"It is certain", "It is decidedly so", "Without a doubt", "Yes definitely", "You may rely on it", "As I see it, yes", "Most likely", "Outlook good", "Yes", "Signs point to yes"}
	nonCommittalAnswers := []string{"Reply hazy, try again", "Ask again later", "Better not tell you now", "Cannot predict now", "Concentrate and ask again", "Very doubtful"}
	negativeAnswers := []string{"Don’t count on it", "My reply is no", "My sources say no", "Outlook not so good"}

	// Create a map to store the arrays
	answers := map[string][]string{
		"Affirmative Answers":   affirmativeAnswers,
		"Non-Committal Answers": nonCommittalAnswers,
		"Negative Answers":      negativeAnswers,
	}

	// Seed the random number generator with the current time
	rand.Seed(time.Now().UnixNano())

	// Loop until the user types "quit"
	for {
		// Get a question from the user
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Ask the Magic 8 Ball a question (or type 'quit' to exit): ")
		question, _ := reader.ReadString('\n')

		// Check if the user wants to quit
		if strings.TrimSpace(strings.ToLower(question)) == "quit" {
			break
		}

		// Generate a random number between 0 and the length of the map - 1
		randomIndex := rand.Intn(len(answers))

		// Get the key (category) at the random index
		i := 0
		var category string
		for k := range answers {
			if i == randomIndex {
				category = k
				break
			}
			i++
		}

		// Get the answer from the array at the random index
		answer := answers[category][rand.Intn(len(answers[category]))]

		// Print the answer
		fmt.Printf("Question: %sAnswer: %s\n", question, answer)
	}

	// Exit the program
	fmt.Println("Goodbye!")
	return
}
