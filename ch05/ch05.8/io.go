package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	stdin := bufio.NewReader(os.Stdin)

	var a int
	var b int

	// GoLang은 선언한 변수는 반드시 사용해야 함.
	n, err := fmt.Scanln(&a, &b)
	if err != nil {
		fmt.Println(err)
		// \n 가 나올 때까지 READ, 입력버퍼를 비우게 됨
		stdin.ReadString('\n')
	} else {
		fmt.Println(n, a, b)
	}
	//===INPUT===
	//hello 4
	//===result===
	//expected integer

	//반복
	n, err = fmt.Scanln(&a, &b)
	if err != nil {
		fmt.Println(err)
		stdin.ReadString('\n')
	} else {
		fmt.Println(n, a, b)
	}
	//===INPUT===
	//3 4
	//===result===
	//2 3 4
}
