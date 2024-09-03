package main

import "fmt"

func main17() {
	var m map[int]string
	m = make(map[int]string)
	var higest int
	higest = 0
	var guest int
	fmt.Scan(&guest)
	for i := 0; i < guest; i++ {
		var name string
		var happy int
		fmt.Scan(&name, &happy)
		m[happy] = name
		if happy > higest {
			higest = happy
		}
	}
	fmt.Println(m[higest])
}
