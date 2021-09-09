package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	args := os.Args[1:]

	if len(args) != 2 {
		fmt.Println("Enter Your name & [positive][negative]")
		return
	}

	moods := [...][5]string{
		{"good", "awesome", "great", "cool", "happy"},
		{"bad", "terrible", "unhappy", "sad", "tired"},
	}

	name, mood := args[0], args[1]

	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(len(moods[0]))

	var mi int

	if mood != "positive" {
		mi = 1
	}

	fmt.Printf("%s feels %s\n", name, moods[mi][n])

}
