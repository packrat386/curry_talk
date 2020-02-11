package main

import (
	"fmt"
	"math/rand"
)

// START OMIT
func formatter(i int) {
	fmt.Printf("<number: %d>\n", i)
}

func printRand(fmter func(int)) {
	fmter(rand.Int())
}

func main() {
	f := formatter
	printRand(f)
}

/*
go run first_class_function.go
<number: 5577006791947779410>
*/
// END OMIT
