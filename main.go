package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Guess game")
	fmt.Println("A number has been drawn, you have 10 shots to figure it out. the number is between 1 and 100")

	chosenNumber := rand.Int64N(101)
	scanner := bufio.NewScanner(os.Stdin)
	shots := [10]int64{}

	for i := range shots {
		fmt.Println("What are your guess?")
		scanner.Scan()
		shot := scanner.Text()
		shot = strings.TrimSpace(shot)
		shotInteger, err := strconv.ParseInt(shot, 10, 64)

		if err != nil {
			fmt.Println("It has to be an integer")
			return
		}

		switch {
		case shotInteger < chosenNumber:
			fmt.Println("Wrong! The number is higher than", shotInteger)
		case shotInteger > chosenNumber:
			fmt.Println("Wrong! The number is lower than", shotInteger)
		case shotInteger == chosenNumber:
			fmt.Printf(
				"Congrats! You did it! The number was: %d\n"+
				"You win it on %d shot\n"+
				"These were your shots: %v\n",
				chosenNumber, i+1, shots[:i],
		)
			return
		}

		shots[i] = shotInteger
	}

	fmt.Printf(
		"Unfortunately, you didn't get it right, the number was: %d\n"+
		"You've had 10 shots.\n"+
		"These were your shots: %v\n",
		chosenNumber, shots,
	)
}
