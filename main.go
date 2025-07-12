package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	//"text/template/parse"

	"rsc.io/quote"
)

var array_of_int []int

func main() {
	fmt.Println(quote.Go())

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a message: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	input = strings.TrimSpace(input)

	//fmt.Println(input)
	// Gunakan struct
	parser := Parser{}
	test := parser.Parse(input)
	fmt.Println(test)
	parser.PrintLengths()

}
