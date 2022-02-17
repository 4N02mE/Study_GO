package main

import (
	"fmt"
	"log"
)

/*
함수는 로직이나 알고리즘을 묶은 코드의 집합이며,
함수형 프로그래밍에서는 기본 단위로 사용되기도 한다.

함수를 사용하는 가장 큰 의미는 "중복을 제거하는 일"

func {func name}을 통해 함수를 정의

func {func name}({입력 받는 값}) (return 해주는 값) {
	...
}
*/
func sayHello(message string) (string, error) {
	return message, nil
}

func sayHello_2() (helloGo string, err error) {
	helloGo = "Hello, Go!"
	return helloGo, err
}

func main() {
	helloGo_1, err_1 := sayHello("Hello, Go!")
	if err_1 != nil {
		log.Fatal(err_1)
	}
	// -> Hello, Go!
	fmt.Println(helloGo_1)
}
