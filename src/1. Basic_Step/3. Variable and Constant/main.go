package main

import (
	"fmt"
)

func main() {
	/*
		자료형 및 변수
		Go -> 정적 타입 언어
		: 컴파일 타임에 타입이 결정되며, 런타임 중 동적으로 타입 변경 X
		  런타임 중 타입 추론 및 최적화 진행 상황 X -> 대체로 빠른 속도
		Go는 사용하지 않는 변수가 있는 경우 에러 표시

		변수 선언 방법:
		1. var {name} {type} -> {name} = {value}
		Ex)	var example int
			example = 0

		2. var {name} {type} = {value}
		Ex) var example int = 0

		3. var {name} = {value}
		Ex) var example = 0

		4. {name} := {value} (변수 단축 선언 방법)
		Ex) example := {value}

		* 만약 타입이 명시되지 않은 경우 컴파일 시 자동으로 타입을 추론 *
	*/

	//int
	var num_1 int = 1
	num_2 := 2
	fmt.Println("num_1: ", num_1, "num_2:", num_2)

	//string
	var str_1 string = "str_variable 1"
	str_2 := "str_variable 2" // string
	fmt.Println("str_1:", str_1, "str_2:", str_2)

	//boolean
	var bool_1 bool = true
	var bool_2 bool = false
	fmt.Println("bool_1:", bool_1, "bool_2:", bool_2)

	/*
		여러 개의 변수 방법:
		1. {name1}, {name2}, {name3} := {value1}, {value2}, {value3}
		Ex) num_, str_, bool_ = 100, "string", true

		2. var ({name1} {type1}\n{name2} {type2}\n{name3} {type3})
		Ex) var {
				num_ int
				str_ string
				bool_ bool
			}

		변수를 선언하고 초기화 할 때에는
		반드시 선언하는 변수와 초기화하는 값의 개수가 같아야 하며,
		타입은 똑같지 않아도 된다bool_1
	*/

	num_3, str_3, bool_3 := 250, "str_variable 3", false
	fmt.Println("num_3:", num_3, "str_3:", str_3, "bool_3:", bool_3)

	var (
		num_4  int
		str_4  string
		bool_4 bool
	)
	fmt.Println("num_4:", num_4, "str_4:", str_4, "bool_4:", bool_4)

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
	num_5, str_5, bool_5 := 256, "str_variable 5", true
	fmt.Println("num_5:", num_5, "str_5:", str_5, "bool_5:", bool_5)

	var (
		num_6  int    = 512
		str_6  string = "str_variable 6"
		bool_6 bool   = true
	)
	fmt.Println("num_6:", num_6, "str_6:", str_6, "bool_6:", bool_6)
}

/*
	전역 스코프에서 전역 변수 선언 시 주의 할 점
	1. 단축 구문 사용시 문법 에러가 발생 -> "var a int = 1" O / "a := 1" X
	2. 변수를 선언하고 따로 할당하는 것이 불가능
*/
var global string = "This is Global Variable"
