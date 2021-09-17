package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Warm-up
//
//  Create and print the following maps.
//
//  1. Phone numbers by last name
//  2. Product availability by Product ID
//  3. Multiple phone numbers by last name
//  4. Shopping basket by Customer ID

//     Each item in the shopping basket has a Product ID and
//     quantity. Through the map, you can tell:
//     "Mr. X has bought Y bananas"
// ---------------------------------------------------------

func main() {

	var (
		phones      map[string]string
		products    map[int]bool
		multiphones map[string][]string
		basket      map[int]map[int]int
	)

	fmt.Printf("phones       : %#v\n", phones)
	fmt.Printf("products     : %#v\n", products)
	fmt.Printf("multiphones  : %#v\n", multiphones)
	fmt.Printf("basket       : %#v\n", basket)
}
