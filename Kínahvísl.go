package main

import "fmt"

func main27() {
	var word1, word2 string
	fmt.Scan(&word1, &word2)
	var count int
	if word1 == word2 {
		fmt.Println(1)
	} else {
		for i := 0; i < len(word1); i++ {
			if word1[i] != word2[i] {
				count++
			}
		}
		fmt.Println(count + 1)
	}

}
