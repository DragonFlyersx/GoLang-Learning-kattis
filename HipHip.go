package main

import "fmt"

func main10() {
	var name string
	var age int
	fmt.Scan(&name, &age)
	for i := 0; i < age; i++ {
		fmt.Println("Hipp hipp hurra, " + name + "!")
	}
}
