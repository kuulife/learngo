package main

import (
	"fmt"
	"strings"
)

func main() {
	//declare a nil slices
	var (
		names    []string
		distance []float64
		data     []uint8
		ratious  []float64
		aliver   []bool
	)

	fmt.Printf("names: %T %d %t\n", names, len(names), names == nil)
	fmt.Printf("distance %T: %d %t\n", distance, len(distance), distance == nil)
	fmt.Printf("data: %T %d %t\n", data, len(data), data == nil)
	fmt.Printf("ratious: %T %d %t\n", ratious, len(ratious), ratious == nil)
	fmt.Printf("aliver: %T %d %t\n", aliver, len(aliver), aliver == nil)

	// ---------------------------------------------------------
	// EXERCISE: Assign slice literals
	//
	//   1. Assign the following data using slice literals to the slices that
	//      you've declared in the first exercise.
	//
	//    1. The names of your three best friends (to the names slice)
	fmt.Println(strings.Repeat("=", 20))
	names = append(names, "Askat", "Chika", "Dastan")
	fmt.Printf("names: %T %d %t\n", names, len(names), names == nil)
	//
	//    2. The distances to five different locations (to the distances slice)
	distance = append(distance, 130.45, 5000.4, 180.0, 60, 45.67)
	fmt.Printf("distance: %T %d %t\n", distance, len(distance), distance == nil)
	//    3. Five bytes of data (to the data slice)
	data = append(data, 1, 2, 3, 4, 5)
	fmt.Printf("data: %T %d %t\n", data, len(data), data == nil)
	//    4. Two currency exchange ratios (to the ratios slice)
	ratious = append(ratious, 23.5, 3.7)
	fmt.Printf("ratious: %T %d %t\n", ratious, len(ratious), ratious == nil)
	//    5. Up/Down status of four different web servers (to the alives slice)
	aliver = append(aliver, true, false, true, true)
	fmt.Printf("alives: %T %d %t\n", aliver, len(aliver), aliver == nil)

	//  3. Compare the length of the distances and the data slices; print a message
	//     if they are equal (use an if statement).
	var equal string

	if len(distance) != len(data) {
		equal = "not "
	}

	fmt.Printf("The length of the distances and the data slices are %sequal.", equal)

}
