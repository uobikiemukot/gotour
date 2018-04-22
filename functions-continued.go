package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func main() {
	var x, y int = 42, 13
	//x = 42
	//y = 13
	fmt.Println(add(x, y))
}
