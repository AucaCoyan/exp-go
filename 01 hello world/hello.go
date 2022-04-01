package main

// import other packages
import (
	"fmt"
	"reflect"
)

func main() {

	// Variables and Constants
	// strings, int, float64 and bool

	var firstName string = "John"

	// := shorthand (interprets the type)

	lastName := "Elder"
	age := 44

	fmt.Println(firstName, lastName, age)

	// Create but don't assign
	var fullName string
	var amount int
	var price float32
	var tf bool

	fmt.Println(fullName, amount, price, tf)

	// You can also print with builting println

	println(fullName, amount, price, tf)

	// and later asign variables with just name = value

	fullName = "AucaCoyan"
	amount = 22
	price = 20.99
	tf = true

	fmt.Println(fullName, amount, price, tf)

	// multiple variable declaration

	var name1, name2 string = "tim", "mary"

	fmt.Println(name1, name2)

	// Constants
	// you can declare them and it doesn't throw an error if you don't use them
	const pi = 3.14

	// multiple declaration of constants

	const (
		// Go linters adjust equals!
		PIZZA1             = "Pepperoni"
		PIZZA2             = "Cheese"
		MYNUM              = 77
		EXTREMELYLONGCONST = "empty"
	)
	fmt.Println(reflect.TypeOf(PIZZA1))
}
