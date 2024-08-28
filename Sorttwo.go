package main

import "fmt"

func main4() {
	var i, j int
	fmt.Scan(&i, &j)
	if i > j {
		fmt.Println(j, i)
	} else {
		fmt.Println(i, j)
	}
}
