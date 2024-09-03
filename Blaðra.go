package main

import "fmt"

func main18() {
	var v, a, t float64
	fmt.Scan(&v, &a, &t)
	var d float64
	d = v*t + 0.5*a*t*t
	fmt.Println(d)
}
