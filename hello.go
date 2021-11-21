package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var language string = "Go"
	out := fmt.Sprintf("Number: %07d is cool", 45)
	// language := "GoRust"
	// var language = "Golang"
	fmt.Printf("hello %s, I'm Here\n", language)
	fmt.Println(out)
	fmt.Printf("Type the year you were born: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	fmt.Printf("You will be %d year old at the end of 2021", 2021-input)

}

// Integers
// 		Unsigned Integers: uint8, uint16, uint32, uint64 (no negatives)
// 		Signed Integers: int8, int16, int32, int64
//		Machine Dependent Types: uint(32 or 64), int(same as uint), uintptr(unsigned int to store the uninterpreted bits of a pointer value)
// Floating
//		Floats: flaot32, float64
//		Complex(Imaginary Parts): complex64(float32 real and other half imaginary), complext128(float64 real and other half imaginary)
// Strings: "Hello World"
// Boleans: true, false

// FMT Cheat Sheet
// 		General: %v value, %T type, %% (literal %)
// 		Boolean: %t (true or false)
//		Integer: %b(base 2), %o(base8). %d(base10), %x(base16)
//		Floating: %e (scientific notation), %f/%F (decimal no exponent), %g (for large exponent)
// 		String : %s(string), %q (string with quotes)

// Width & Precision
// 		%f(default with, default precision)
// 		%9f(witdth 9, default precision)
// 		%.2f(defult width, 2 precision)
// 		%9.f(9 width, 0 precision)
//

// Padding
//		%09d (pads digit to length 9 with prceeding 0's)
//		%-4d (pads with spaces (witdth 4, left justified))
//
//
