package main

import "fmt"

func Divide(a, b int) (int, bool) {
	if b == 0 {
		return 0, false
	}

	return a / b, true
}

func main() {
	c, success := Divide(9, 3)
	fmt.Println(c, success)

	// 하나라도 다른 (d)게 있으면 선언대입문을 허용해 줌
	d, success := Divide(9, 0)
	fmt.Println(d, success)
	//===result===
	//3 true
	//0 false

}
