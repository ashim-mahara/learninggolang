package main

import (
	"fmt"
)

// && is and, || is or
func main() {
	fmt.Printf("true && true \t%v\n", true && true)
	fmt.Printf("true && false \t%v\n", true && false)
	fmt.Printf("true || true \t%v\n", true || true)
	fmt.Printf("true || false \t%v\n", true || false)
	fmt.Printf("!true \t\t%v\n", !true)
	fmt.Printf("!false \t\t%v\n", !false)

}
