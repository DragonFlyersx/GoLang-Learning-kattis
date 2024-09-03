package main

import "fmt"

func main20() {
	var n int
	var name1, name2 string
	fmt.Scan(&n)
	fmt.Scan(&name1, &name2)
	if n == 2 {
		fmt.Println("Blandad best")
	} else {
		fmt.Println(name1)
		fmt.Println(name2)
	}
}
