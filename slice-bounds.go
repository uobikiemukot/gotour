package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	a := [3]int{2, 3, 4}

	p := &s
	pp := &p
	fmt.Println(p)
	fmt.Println(pp)
	fmt.Println(s)
	fmt.Println(a)

	a = a[1:2]

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)
}
