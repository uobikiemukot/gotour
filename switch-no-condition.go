package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t, t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute())
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
