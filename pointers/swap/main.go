// ---------------------------------------------------------
// EXERCISE: Swap
//
//  Using funcs, swap values through pointers, and swap
//  the addresses of the pointers.
//

//

//

//

//
// ---------------------------------------------------------

package main

import "fmt"

func main() {
	//1- Swap the values through a func
	//a- Declare two float variables
	a, b := 3.14, 6.28

	//c- Pass the variables' addresses to the func
	swap(&a, &b)

	//d- Print the variables
	fmt.Println("a: ", a)
	fmt.Println("b: ", b)

	//2- Swap the addresses of the float pointers through a func
	//a- Declare two float pointer variables and,
	//assign them the addresses of float variables
	aa, bb := &a, &b

	//c- Pass the pointer variables to the func
	aa, bb = swapAdd(aa, bb)

	//d- Print the addresses, and values that are
	//pointed by the pointer variables
	fmt.Printf("aa: %p — bb: %p\n", aa, bb)
	fmt.Printf("aa: %g — bb: %g\n", *aa, *bb)

}

//     b- Declare a func that can swap the variables' values
//        through their memory addresses
func swap(a, b *float64) {
	*a, *b = *b, *a
}

//     b- Declare a func that can swap the addresses
//        of two float pointers
func swapAdd(a, b *float64) (*float64, *float64) {
	return b, a
}
