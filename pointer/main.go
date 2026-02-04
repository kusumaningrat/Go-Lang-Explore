/*
Pointer is a refrence or memory address. Variables that use pointers mean where the address of the variable is located.
For example, we have variable a = 4, for &b = <memory_address> and for *b = 4 cause it will access the value from that address.

There are two important things to note about pointer:
 1. We can get pointer value of normal variable using ampersand (&). For example we have variable a,
    then we can get the pointer value using &a (it will return memory address). This method is called refrencing
 2. Otherwise, we can get pointer variable value with artrisk (*). This method is called derefrencing
*/
package main

import "fmt"

type User struct {
	ID   string
	Name string
}

func main() {
	a := 10

	// Access memory address of variable a (eg, 0xc000012110)
	b := &a

	fmt.Println(a)
	fmt.Println(b)

	// Access address and get the value from that address
	c := *b

	// Assign new value to variable
	c = 11

	// Change a value with pointer
	*b = 15
	fmt.Println(c)

	// You will see 15 now
	fmt.Println(a)

	// Another example
	d := 15
	var e *int = &d
	fmt.Println(e)

	fmt.Println()

	// Pointer as parameter
	number := 5

	// Before changed
	fmt.Println(number)

	// After changed
	ChangeNumber(&number, 10)
	fmt.Println(number)
}

// Pointer as parameter
func ChangeNumber(original *int, value int) {
	*original = value
}
