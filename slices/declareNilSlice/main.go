package main

import "fmt"

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

}
