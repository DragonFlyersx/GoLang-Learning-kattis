package main

import (
	"fmt"
	"regexp"
)

func main32() {
	var word string
	fmt.Scan(&word)
	fmt.Println(len(RemoveSpecialChars(word)))
}

func RemoveSpecialChars(str string) string {
	reg, _ := regexp.Compile("[^a-zA-Z0-9\\s]+")
	return reg.ReplaceAllString(str, "")
}
