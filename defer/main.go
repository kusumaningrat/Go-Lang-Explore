/*

`Defer` can be use to delay the execution of a statement before the function block ends.
While `Exit` can be use to stop the program forcely

*/

package main

import "fmt"

func Order(username string, menu string) {
	// Assume that first customer coming earlier but he orders a pizza, while the second customer only other mineral water
	if menu == "pizza" {
		defer fmt.Printf("%s is ordering %s\n", username, menu)
		return
	}
	fmt.Printf("%s is ordering %s\n", username, menu)

}

func main() {

	// Mineral water should be serve first cause it fast serving
	Order("Alex", "mineral water")
	Order("John", "pizza")
}
