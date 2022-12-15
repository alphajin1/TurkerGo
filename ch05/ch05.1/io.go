package main

import "fmt"

func main() {
	var a int = 10
	var b int = 20
	var f float64 = 32799438743.8297

	fmt.Print("[Print]", "a:", a, "b:", b, "f:", f)
	fmt.Println("[Println]", "a:", a, "b:", b, "f:", f)
	fmt.Printf("[Printf] a: %d b: %d f:%f\n", a, b, f)
	// 서식지정자
}
