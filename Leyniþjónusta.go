package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main31() {
	reader := bufio.NewReader(os.Stdin)
	words, _ := reader.ReadString('\n')
	words = strings.TrimSpace(words)
	words = strings.Replace(words, " ", "", -1)
	fmt.Println(words)

}
