package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	c1 := 1
	c2 := 0

	f := func() int {
		tmp := c2
		c2 = c1 + c2
		c1 = tmp
		return c1
	}

	return f
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
