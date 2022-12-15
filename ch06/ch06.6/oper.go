package main

import "fmt"
import "math"

func equal(a, b float64) bool {
	// NextAfter는 a > b로 1bit 만큼만 진행 함.
	// 즉, 아주 작은 오차에 대해서는 이 방식으로 보정이 가능하다.
	return math.Nextafter(a, b) == b
}

func main() {
	var a = 0.1
	var b = 0.2
	var c = 0.3
	fmt.Printf("%f + %f == %f : %v\n", a, b, c, a+b == c)
	fmt.Println(a + b)
	//===result===
	//0.100000 + 0.200000 == 0.300000 : false
	//0.30000000000000004

	fmt.Printf("%0.18f == %0.18f: %v\n", a+b, c, equal(a+b, c))
	fmt.Println(a + b)

	//a+b와 c는 bit 1개만 차이가 남.
	//===result===
	//0.300000000000000044 == 0.299999999999999989: true
	//0.30000000000000004

}
