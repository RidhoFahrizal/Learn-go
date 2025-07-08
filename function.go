package main

import (
	"fmt"
	"rsc.io/quote"	
)

func main() {
    //Menampilkan quote
    fmt.Println(quote.Go())
	helper()
	loop()
}

// Fungsi untuk membagi dua bilangan dan mengembalikan hasil bagi dan sisa
func divide(a int, b int) (int, int) {
	return a / b, a % b
}

func greet(name string) string {
	return name
}
