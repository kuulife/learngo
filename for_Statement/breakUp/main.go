// ---------------------------------------------------------
// EXERCISE: Break Up
//
//  1. Extend the "Only Evens" exercise
//  2. This time, use an infinite loop.
//  3. Break the loop when it reaches to the `max`.
//
// RESTRICTIONS
//  You should use the `break` statement once.
//
// HINT
//  Do not forget incrementing the `i` before the `continue`
//  statement and at the end of the loop.
//
// EXPECTED OUTPUT
//    go run main.go 1 10
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

	min, err1 := strconv.Atoi(args[0])
	max, err2 := strconv.Atoi(args[1])

	if err1 != nil || err2 != nil || min >= max {
		fmt.Println("Wrong numbers")
		return
	}

	var (
		i   = min
		sum int
	)

	for {
		if i > max {
			break
		} else if i%2 != 0 {
			i++
			continue
		}
		fmt.Print(i)
		if i < max {
			fmt.Print(" + ")
		}
		sum += i
		i++
	}
	fmt.Printf(" = %d\n", sum)
}
