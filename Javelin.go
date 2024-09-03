package main

import "fmt"

func main22() {
	var n int
	var total int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var rod int
		fmt.Scan(&rod)
		total += rod
	}
	fmt.Println(total - (n - 1))
}
