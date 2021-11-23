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
	fmt.Printf("You will be %d year old at the end of 2021\n", 2021-input)
	// scanner.Scan()
	// val,_ := strconv.ParseInt(scanner.Text(),10,64)
	// if val == 10{
	// 	fmt.Println("Equal to 10")
	// } else if val > 10 {
	// 	fmt.Println("Greator than 10")
	// } else if(9 < 10){
	// 	fmt.Println("Smaller than 10")
	// }
	// loop
	for i:=0; i < 5; i+=1{
		fmt.Println("in loop, count:",i)
	}

	i :=3
	switch i {
	case 1:
		fmt.Println("1. one")
	case 2:
		fmt.Println("2. two")
	default:
		fmt.Println("defualt")
	}

	// Array
	var arr [5]int 
	arr[0] = 100
	fmt.Println(arr)

	newarr := [3]int{4,5,6}
	fmt.Println(newarr, " len: ",len(newarr))

	arr2D := [2][3]int{{1,2,3},{4,5,6}}
	fmt.Println(arr2D)


	// Slices
	var s []int = newarr[:]
	s = append(s,4)
	
	fmt.Printf("capacity: %v, lenth: %v ", cap(newarr),len(newarr))
	fmt.Println(newarr)
	fmt.Printf("capacity: %v, lenth: %v ", cap(s),len(s))
	fmt.Println(s)


	// Range 
	for x, ele := range s{
		fmt.Printf("%v: %v\n",x,ele)
	}

	// Maps
	var mymap map[string]int = map[string]int{
		"one":1,
		"two":2,
		"three":3,
		"four":4,
	} 

	fmt.Println(mymap)
	mymap["five3"] = 5
	delete(mymap,"zero")
	delete(mymap,"one")
	val, ok := mymap["two"]

	fmt.Println(mymap)
	fmt.Println(val,ok)


	// function 
	res := apsolute(-5)
	fmt.Println(res)
}


 func apsolute(x int) int {
	if(x < 0){
		return -x
	}
	return x
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
