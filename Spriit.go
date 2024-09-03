package main

import "fmt"

func main10() {
	var classrooms, sprit int
	fmt.Scan(&classrooms, &sprit)
	for i := 0; i < classrooms; i++ {
		var needed int
		fmt.Scan(&needed)
		sprit -= needed
	}
	if sprit < 0 {
		fmt.Println("Neibb")
	} else {
		fmt.Println("Jebb")
	}

}
