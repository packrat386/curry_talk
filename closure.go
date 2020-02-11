package main

import (
	"fmt"
)

// START OMIT
func newCounter() func() int {
	i := 0
	return func() int {
		ret := i
		i++
		return ret
	}
}

func main() {
	ctr := newCounter()
	fmt.Println(ctr())
	fmt.Println(ctr())
	fmt.Println(ctr())
}

/*
go run closure.go
0
1
2
*/
// END OMIT
