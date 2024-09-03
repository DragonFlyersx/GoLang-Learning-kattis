package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main16() {
	reader := bufio.NewReader(os.Stdin)
	words, _ := reader.ReadString('\n')
	words = strings.TrimSpace(words)
	fmt.Println(words)
}
