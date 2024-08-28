package main

import "fmt"

func main2() { //always named main
	var word string
	fmt.Scan(&word)
	for i := 0; 3 > i; i++ {
		fmt.Println(word)
	}
}
