package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Random number generator

	source := rand.NewSource(time.Now().UnixNano())
	randomizer := rand.New(source)
	secretNumber := randomizer.Intn(100) // generates numbers between 0 and 100
	

	var guess int

	for {
		fmt.Println("Guess the number!")
		fmt.Println("Please input your guess: ")
		fmt.Scan(&guess)
		if guess > secretNumber {
			fmt.Println("Too big")
		} else if guess < secretNumber{
			fmt.Println("Too small")
		} else {
			fmt.Println("You win!")
			break
		}
	}
}