package main

import (
	"fmt"
	"strings"
)

func main23() {
	var word string
	fmt.Scan(&word)
	if strings.Contains(word, "ss") {
		fmt.Println("hiss")
	} else {
		fmt.Println("no hiss")
	}
}
