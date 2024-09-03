package main

import "fmt"

func main24() {
	var m map[int]int
	m = make(map[int]int)

	var n, c, q, t, number, higest int
	higest = 0
	for i := 0; i < 5; i++ {
		fmt.Scan(&n, &c, &q, &t)
		number = (n + c + q + t)
		m[number] = i
		if number > higest {
			higest = number
		}
	}
	fmt.Println(m[higest]+1, " ", higest)
}
