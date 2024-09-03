package main

import "fmt"

func main25() {
	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var name string
		fmt.Scan(&name)
		fmt.Println("Takk " + name)
	}

}
