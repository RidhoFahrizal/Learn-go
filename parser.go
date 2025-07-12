package main

import (
	"fmt"
	"strings"
)

type Parser struct {
	Parts        []string
	Array_of_int []int
}

func (p *Parser) Parse(input string) []int {
	p.Parts = strings.SplitAfter(input, " ")
	p.Array_of_int = make([]int, len(p.Parts))
	for i, part := range p.Parts {
		p.Array_of_int[i] = len(part) - 1
	}

	return p.Array_of_int
}

func (p *Parser) PrintLengths() {
	for i, length := range p.Array_of_int {
		fmt.Printf("Wrod %d: %d\n", i+1, length)
	}
}
