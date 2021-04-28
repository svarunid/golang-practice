//Create a variable of type string using a raw string literal. Print it.
package main

import "fmt"

func main() {
	rawString := `Hello there
I am Arun
I am code in golang
"I am enclosed double quotes"
'And I am in single quotes'`
	fmt.Println(rawString)
}
