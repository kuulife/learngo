package main

import (
	"fmt"
	"limitBackingArrSharing/api"
)

func main() {
	//get the first three elements from api.temps
	received := api.Read(1, 4)
	fmt.Println("api.read : ",received)

	//append changes the api package's temps slices's
	//backing array as well
	received = append(received, []int{1, 3}...)

	fmt.Println("api.temps     :", api.All())
	fmt.Println("main.received :", received)

}
