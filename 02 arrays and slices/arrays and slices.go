package main

import (
	"fmt"
)

func main() {

	// you establish the length when you create the array
	// you cant add or remove items from an array
	var firstNames = [3]string{"1", "234", "April"}
	fmt.Println(firstNames)

	// also, you can use the shorthand for definition
	// and infer the length too
	lastNames := []string{"Smith", "Pattinson"}
	fmt.Println(lastNames)

	// change items
	lastNames[1] = "Robertson"
	fmt.Println(lastNames)

	// You can default assignments
	var numb = [4]int{1, 2}            //defaults to 0
	var strings = [4]string{"a", "b"}  // defaults to ""
	var floats = [4]float32{0.0, 19.0} //defaults to 0

	fmt.Println(numb, strings, floats)

	// you can set the initial position of the array this way:
	var numbers = [10]int{0: 41, 4: 99}
	fmt.Println(numbers)

	// Slices
	var toppings = [5]string{
		"Pepperoni", "Onion", "Cheese", "Supreme", "Bacon"}
	fmt.Println(toppings)

	// make a slice fomr an array
	toppingsSlice := toppings[0:3]
	fmt.Println(toppingsSlice)
	// modify it
	toppingsSlice[1] = "Peppers"
	fmt.Println(toppingsSlice)
	// append items
	toppingsSlice = append(toppingsSlice, "Anana")
	fmt.Println(toppingsSlice)
	// you can add slices together
	otherSlice := toppings[3:4]
	fmt.Println(otherSlice)

	// you have to add "..." at the end of the slice
	otherSlice = append(otherSlice, toppingsSlice...)
	fmt.Println(otherSlice)
}
