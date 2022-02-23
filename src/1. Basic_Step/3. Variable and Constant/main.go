package main

import (
	"fmt"
)

func main() {
	/*
		Go는 정적 타입 언어이다.
		컴파일 타임에 타입이 결정되며, 런타임 중 동적으로 타입변경이 불가하다.
		하지만 런타임 중 타입 추론 및 최적화 진행 상황이 없어
		대체로 빠른 속도를 갖는다.
		또한 Go는 사용하지 않는 변수가 있을 시 에러가 표시된다.

		변수/상수는 키워드 "var"/"const" 를 사용하여 선언한다.
		"var"/"const" 키워드 뒤에 변수의 이름과 타입을 기제하여 선언한다.
		Ex) var x int | const y string

		변수 선언 문법 4가지
		var number_1 int 	-> 변수 선언 후 값 초기화
		num_1 = 0

		var num_2 int = 0	-> 변수 선언과 동시에 초기화

		var num_3 = 0		-> 변수 타입을 생략하여 선언

		num_4 := 0			-> "var" 키워드와 변수 타입을 생략하여 선언

		[!]	변수 타입이 명시되지 않은 경우 컴파일 시 자동으로 타입을
			추론하여 타입이 붙게 된다.

		상수 선언 문법
		const num_1 int = 0

		[!]	상수 선언 시 주의점
			1. 한 번 선언되고 할당되면 값을 바꿀 수 없다.
			2. 선언과 초기화를 따로 할 수 없으며 선언할 때 초기화를
			   같이 해주어야 한다.
	*/

	// 첫 번째 선언 방법
	var number_1 int
	number_1 = 1

	// 두 번쨰 선언 방법
	var number_2 int = 2

	// 세 번째 선언 방법
	var number_3 = 3

	// 네 번째 선언 방법
	number_4 := 4

	fmt.Println(number_1, number_2, number_3, number_4)

	/*
		여러 개의 변수 선언 문법
		1. 한 줄로 선언
		num_1, str_1, bool_1 := 0, "string", true

		2. ()로 묶어서 선언
		var (
			num_2 int
			str_2 string
			bool_2 bool
		)

		여러 개의 상수 선업 문법
		1. 한 줄로 선언
		const num_1, str_1, bool_1 := 0, "string", true

		2. ()로 묶어서 선언
		const (
			num_2 int = 0
			str_2 string = "string"
			bool_2 bool = true
		)

		[!]	변수를 선언 후 초기화 시
			반드시 선언하는 변수와 초기화하는 값의 개수가 같아야한다.
			또한 타입이 똑같지 않아도 된다.
	*/
	number_5, str_5, bool_5 := 16, "다중선언", false
	
	const (
		num_6 int =
	)

	/*
		상수 선언 방법
		const num_5 int = 1

		여러 개의 상수 선언 방법:
		1. const {name1}, {name2}, {name3} = {value1}, {value2}, {value3}
		Ex) const num_, str_, bool_ = 0, "", false

		2. const ({name1} {type1}\n{name2} {type2}\n{name3} {type3})

		Ex) const (
				num_ int = 0
				str_ string = ""
				bool_ bool = true
			)

		상수 주의점
		1. 한 번 선언되고 할당되면 값을 바꿀 수 없다.
		2. 선언과 할당은 따로 할 수 없으며 선언할 떄 초기화를 같이 해주어야 한다.
	*/
}

/*
	전역 스코프에서 전역 변수 선언 시 주의 할 점
	1. 단축 구문 사용시 문법 에러가 발생 -> "var a int = 1" O / "a := 1" X
	2. 변수를 선언하고 따로 할당하는 것이 불가능
*/
var global string = "This is Global Variable"
