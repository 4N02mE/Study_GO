package main

import "fmt"

const a, b int = 10, 3

var c, i int

func main() {
	/*
		연산자

		Go 언어는 다른 언어들과 비슷하게
		산술연산자, 관계연산자, 논리연산자, Bitwise연산자,
		할당연산자, 포인터연산자 등을 지원한다.

		산술연산자는 사칙연산자(+, -, *, /, %)와 증감연산자(++, --)를
		사용한다.
	*/
	c = (a + b) / 5
	i++

	//관계연산자는 서로의 크기를 비교하거나 동일함을 체크하는데 사용한다.
	result := a == b
	result = a != c
	result = a >= b
	if result != false {
		return
	}

	//논리연산자는 AND, OR, NOT을 표현하는데 사용된다.
	c = (a & b) << 5

	/*
		할당연산자는 값을 할당하는 = 연산자 외에 사칙연산, 비트연산을
		축약한 +=, &=, <<= 같은 연산자들을 말한다.
	*/
	d := 100
	c += d
	c >>= 2
	d |= d

	/*
		포인터연산자는 & 혹은 * 을 사용하여 해당 변수의 주소를
		얻어내거나 이를 반대로 Dereference 할 때 사용한다.
		포인터연산자는 제공하지만 포인터에 더하고 빼는 기능은
		제공하지 않는다.
	*/
	var k int = 10
	var p = &k      // k의 주소를 할당
	fmt.Println(*p) // p가 가리키는 주소에 있는 실제 내용을 출력 = k의 값
}
