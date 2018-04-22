package main

import "fmt"

func a() {
	i := 0
	defer fmt.Println(i)
	i++
	return
}
func c() (i int) {
	defer func() { i++ }()
	return 1
}

func main() {
	/*
		fmt.Println("counting")

		for i := 0; i < 10; i++ {
			defer fmt.Println(i)
		}

		fmt.Println("done")
	*/
	fmt.Println(c())
}
