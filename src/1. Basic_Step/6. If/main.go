package main

import "fmt"

func main() {
	/*
		if문은 해당 조건이 맞으면 {} 안의 내용을 실행한다.
		Go의 if 조건문은 ()로 둘러 싸지 않아도 된다.

		if문의 조건식은 반드시 Boolean식으로 표현되어야 한다.
	*/

	var a int = 10
	if a == 10 {
		fmt.Println("a is 10")
	}

	/*
		if문은 else if 혹은 else문을 할께 가질 수 있다.
		else if문은 if 조건문이 거짓일 때
		다시 다른 if 조건식을 검사하는데 사용되며,
		else문은 이전 if문들이 모두 거짓일 때 실행된다.
	*/

	var b int = 64
	if b == 16 {
		fmt.Println("b is 16")
	} else if b == 32 {
		fmt.Println("b is 32")
	} else if b == 64 {
		fmt.Println("b is 64")
	} else {
		fmt.Println("b is nothing")
	}

	/*
		if문에서 조건식을 사용하기 이전 간단한 문장(Optional Statement)을
		함께 실행할 수 있다.
		주의점은 이떄 정의된 변수는 해당 if문 내에서만 사용할 수 있어
		if문에서 벗어나면 사용할 수 없게 된다.
	*/
	var c int = 2
	if i := c * 2; i < b { // b = 32
		fmt.Println(i)
	}
	// fmt.Println(i) -> if문을 벗어나 사용할 수 없어 에러가 난다.
}
