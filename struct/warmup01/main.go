package main

import (
	"fmt"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Warm Up
//
//  Starting with this exercise, you'll build a command-line
//  game store.
//
//  1. Declare the following structs:
//
//     + item: id (int), name (string), price (int)
//
//     + game: embed the item, genre (string)
//
//  2. Create a game slice using the following data:
//
//     id  name          price    genre
//
//     1   god of war    50       action adventure
//     2   x-com 2       30       strategy
//     3   minecraft     20       sandbox
//
//
//  3. Print all the games.
//
// EXPECTED OUTPUT
//  Please run the solution to see the output.
// ---------------------------------------------------------

type item struct {
	id    int
	name  string
	price int
}

type game struct {
	item
	genre string
}

func main() {

	games := []game{
		{
			item:  item{1, "god of war", 50},
			genre: "action adventure",
		},
		{
			item:  item{2, "x-com 2", 30},
			genre: "strategy",
		},
		{
			item:  item{3, "minecraft", 20},
			genre: "sandbox",
		},
	}

	fmt.Printf("Kutman's game store has %d games.\n\n", len(games))

	fmt.Printf("%s  %-15s %-20s %s\n%s\n",
		"ID", "Name", "Genre", "Price", strings.Repeat("=", 45))

	for _, g := range games {
		fmt.Printf("#%d: %-15q %-20s $%d\n",
			g.id, g.name, "("+g.genre+")", g.price)
	}

}
