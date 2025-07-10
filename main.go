package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var array_of_int []int

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter a message: ")
	input, err := reader.ReadString('\n')

	if err != nil {
		fmt.Print("terdapat error di penggunaan")
		return
	}

	parts := strings.SplitAfter(input, " ")

	fmt.Println(parts)
	parts_length := len(parts)
	array_of_int = make([]int, parts_length)
	for i := 0; i < parts_length-1; i++ {

		array_of_int[i] = len(parts[i]) - 1

	}

	for i := 0; i < parts_length-1; i++ {
		fmt.Println(array_of_int[i])
	}
}
