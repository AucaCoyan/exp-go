package main

import (
	"fmt"
	"math"
)

func main() {

	// operators: math, assignment, comparison
	fmt.Println("Math Operators")
	fmt.Println("10+5 =", 10+5)
	fmt.Println("10+5 =", 10+5)
	fmt.Println("10*5 =", 10*5)
	fmt.Println("10/5 =", 10/5)
	fmt.Println("10%5 =", 10%5)

	// power
	var base, exponent float64
	base = 2
	exponent = 3
	fmt.Println("2 to the 3rd power", math.Pow(base, exponent))

	num := 10
	num++
	fmt.Println(num)

	num2 := 10
	num2--
	fmt.Println(num2)

	// assignment operators, =, +=, -=, *=, /=
	fmt.Println("Assignment Operators")
	myNum := 10
	myNum *= 5 // multiplies the left side by the right side
	fmt.Println(myNum)

	// Comparison Operators ==, !=, >, <, >=, <=
	fmt.Println("Comparison Operators")
	fmt.Println("5==5", 5 == 5)
	fmt.Println("5!=6", 5 != 6)
	fmt.Println("5>6", 5 > 6)
	fmt.Println("5<6", 5 < 6)
	fmt.Println("5>=6", 5 >= 6)
	fmt.Println("5<=6", 5 <= 6)

}
