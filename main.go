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
	fmt.Println("Welcome to the Guessing Game")
	fmt.Println("A random number will be drawn. Try to guess correctly. The number will be an integer between 0 and 100.")

	x := rand.Int64N(101)
	scanner := bufio.NewScanner(os.Stdin)

	attempts := [10]int64{}

	for i := range attempts {
		fmt.Println("What's your guess?")
		scanner.Scan()
		attempted := scanner.Text()
		attempted = strings.TrimSpace(attempted)

		attemptedInt, err := strconv.ParseInt(attempted, 10, 64)
		if err != nil {
			fmt.Println("The guess must be a whole number!")
			return
		}

		switch {
		case attemptedInt < x:
			fmt.Println("You're wrong. The number drawn is greater than: ", attemptedInt)
		case attemptedInt > x:
			fmt.Println("You're wrong. The number drawn is less than: ", attemptedInt)
		case attemptedInt == x:
			fmt.Printf("Congratulations!! You got it right, the number was: %d\n"+
				"You got it right in %d attempts\n"+"These were his attempts: %v\n", x, i+1, attempts[:i])
			fmt.Println("\nPress the ENTER key to close the program...")
			bufio.NewReader(os.Stdin).ReadBytes('\n')
			return
		}

		attempts[i] = attemptedInt
	}

	fmt.Printf("Unfortunately, you got the number wrong, which was: %d\n"+
		"These were his attempts: %v\n", x, attempts)

	fmt.Println("\nPress the ENTER key to close the program...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
