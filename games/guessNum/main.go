package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	//random seed
	rand.Seed(time.Now().UnixNano())

	//guess number
	num := os.Args[1:]

	//check if user inputs a number
	if len(num) != 1 {
		fmt.Println("Pick a number")
		return
	}

	//return string to int
	guess, err := strconv.Atoi(num[0])

	if guess < 0 {
		fmt.Println("Enter positive number 0-10")
		return
	}

	//check if input is number
	if err != nil {
		fmt.Printf("%q is not a number", num[0])
		return
	}

	for turn := 1; turn <= 5; turn++ {

		//random number
		n := rand.Intn(11)

		//turn
		fmt.Printf("turn: %d, n: %d\n", turn, n)

		//check if win at the first try
		if turn == 1 && n == guess {
			fmt.Printf("Congrats!!!, You win at the first try\n")
			return
		}

		//check if guess is exact number
		if n == guess {
			fmt.Printf("YOU WIN , at turn(%d) number was: %d\n", turn, guess)
			return
		}
	}
	fmt.Println("YOU LOSE")

}
