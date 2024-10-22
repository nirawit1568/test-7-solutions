package main

import (
	"fmt"
	"strings"
)

func main() {

	var input string
	fmt.Print("Enter yout input: ")
	fmt.Scanln(&input)

	output := FindMinLeftRight(strings.TrimSpace(input))
	fmt.Println(output)
}
