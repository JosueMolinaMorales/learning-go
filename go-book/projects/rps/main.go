package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	s "strings"
)

func main() {
	playerWins := 0
	compWins := 0
	keepPlaying := ""
	for keepPlaying != "n" {
		// Rock paper scissors
		choices := map[int]string{1: "Rock", 2: "Paper", 3: "Scissors"}
		compChoice := choices[rand.Intn(len(choices))+1]
		// Get the choice of the user
		fmt.Println("Please enter one of the following choices:")
		fmt.Println("1. Rock")
		fmt.Println("2. Paper")
		fmt.Println("3. Scissor")
		fmt.Print(">>>>>>>>> ")
		reader := bufio.NewReader(os.Stdin)
		pChoice, err := reader.ReadString('\n')

		if err != nil {
			log.Fatalf("Could not read user's input, %v", err)
			return
		}

		// Check the user's input
		pChoice = s.Trim(pChoice, " \n\t\r")
		piChoice, err := strconv.Atoi(pChoice)

		if err != nil {
			log.Fatalf("[ERROR] %s is not a number", pChoice)
			return
		}
		if piChoice < 1 || piChoice > 3 {
			log.Fatalf("[ERROR] %d is not a valid number", piChoice)
			return
		}

		pChoice = choices[piChoice]
		fmt.Printf("Player's choice is: %s\n", pChoice)
		fmt.Printf("Computer's choice is: %s\n", compChoice)
		// Check for winner
		if pChoice == compChoice {
			// Draw
			fmt.Println("It's a draw!")
		} else if pChoice == "Rock" {
			if compChoice == "Paper" {
				fmt.Println("Computer wins!")
				compWins++
			}
			if compChoice == "Scissors" {
				fmt.Println("Player wins!")
				playerWins++
			}
		} else if pChoice == "Paper" {
			if compChoice == "Rock" {
				fmt.Println("Player wins!")
				playerWins++
			}
			if compChoice == "Scissors" {
				fmt.Println("Computer wins!")
				compWins++
			}
		} else if pChoice == "Scissors" {
			if compChoice == "Rock" {
				fmt.Println("Computer wins!")
				compWins++
			}
			if compChoice == "Paper" {
				fmt.Println("Player wins!")
				playerWins++
			}
		}
		fmt.Print("Play again? (y/n) >>>>> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("Could not read the user's input, %v", err)
		}
		keepPlaying = s.Trim(input, " \n\t\r")
	}
	fmt.Println("Thanks for playing!")
	fmt.Println("============RESULTS==============")
	fmt.Println("Player wins: ", playerWins)
	fmt.Println("Computer wins: ", compWins)
}
