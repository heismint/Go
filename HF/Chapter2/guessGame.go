// Guess challenges players to guess a random number.
package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"			// Imported packages
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	seconds := time.Now().Unix()			// Get the current date and time as an integer
	rand.Seed(seconds)			// Seed the random number generator
	target := rand.Intn(100) + 1			// Generate number between 1 and 100
	fmt.Println("I've chosen a random number between 1 and 100.")
	fmt.Println("Can you guess it?")
	//fmt.Println(target) 				// Commented code that prints the target

	reader := bufio.NewReader(os.Stdin)			// Create a bufioReader that lets us read keyboard input

	success := false			// Set up to print a failure message by default
	for guesses := 0; guesses < 10; guesses++ { 			// Loop to tell user how many guesses they have left
		fmt.Println("You have", 10-guesses, "guesses left")
		fmt.Print("Make a guess: ") 			// Ask for a number
		input, err := reader.ReadString('\n') 		//Read what the user types, up until they press enter
		if err != nil {
			log.Fatal(err)			// If there's an error, print the message and exit
		}
		input = strings.TrimSpace(input)			// Newline when user enters an inputs is removed
		guess, err := strconv.Atoi(input)			// Input is converted to an Int using Atoi
		if err != nil {
			log.Fatal(err)			// If there's an error, print the message and exit
		}

		if guess < target {			// If the guess was too low, this tells the user
			fmt.Println("Oops. Your guess was LOW.")
		} else if guess > target {			// If the guess was tp high this tells the user
			fmt.Println("Oops. Your guess was HIGH.")
		} else {			// Otherwise the guess must be correct...
			success = true			// Prevent the failure message from displaying
			fmt.Println("Good job! You guessed it!")
			break			// Exit the loop
		}
	}
	if !success {			// The boolean negation (!) turns true values to false and false values to true, if "succes" is false, tell the player what the target was
		fmt.Println("Sorry, you didn't guess my number. It was:", target)
	}
}
