package main

import "fmt"

func main() {
	var a float32 = 1234.523
	var b float32 = 3456.123
	var c float32 = a * b
	var d float32 = c * 3
	fmt.Println(a)
	fmt.Println(b)
	// 1234.523
	// 3456.123

	fmt.Println(c)
	fmt.Println(d)
	// 4266663.334334329 = 1234.523 * 3456.123
	// 부동소숫점 오차
	// 4.266663e+06
	// 1.2799989e+07

}
