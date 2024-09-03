package main

import "fmt"

func main14() {
	var word string
	var n int
	fmt.Scan(&word, &n)

	for i := 0; i < n; i++ {
		fmt.Print(word)
	}
}
