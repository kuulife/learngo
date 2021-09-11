package main

import (
	"bytes"
	"fmt"
)

// ---------------------------------------------------------
// EXERCISE: Append
//  Please follow the instructions within the code below.
// EXPECTED OUTPUT
//  They are equal.
// HINTS
//  bytes.Equal function allows you to compare two byte
//  slices easily. Check its documentation: go doc bytes.Equal
// ---------------------------------------------------------

func main() {
	png, header := []byte{'P', 'N', 'G'}, []byte{}

	// 2. append elements to header to make it equal with the png slice
	header = append(header, png...)
	// 3. compare the slices using the bytes.Equal function
	hello := bytes.Equal(png, header)
	// 4. print whether they're equal or not
	if hello {
		fmt.Println("They are equal")
	}
}
