// ---------------------------------------------------------
// EXERCISE: Only Evens
//
//  1. Extend the "Sum up to N" exercise
//  2. Sum only the even numbers
//
// RESTRICTIONS
//  Skip odd numbers using the `continue` statement
//
// EXPECTED OUTPUT
//  Let's suppose that the user runs it like this:
//
//    go run main.go 1 10
//
//  Then it should print:
//
//    2 + 4 + 6 + 8 + 10 = 30
// ---------------------------------------------------------
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]

	if len(args) != 2 {
		fmt.Println("Give me a two numbers")
		return
	}

	var sum int
	min, err1 := strconv.Atoi(args[0])
	max, err2 := strconv.Atoi(args[1])

	if err1 != nil || err2 != nil || min >= max {
		fmt.Println("Wrong numbers")
		return
	}

	for i := min; i <= max; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Print(i)
		if i < max {
			fmt.Print(" + ")
		}
		sum += i
	}
	fmt.Printf(" = %d\n", sum)
}
