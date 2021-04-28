//Create TYPED and UNTYPED constants. Print the values of the constants.
package main

import "fmt"

const unTyped = 52
const typed float64 = 55.2

func main() {
	fmt.Println(unTyped)
	fmt.Println(typed)
}
