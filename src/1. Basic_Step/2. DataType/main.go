package main

import "fmt"

func main() {
	/*
		Go 프로그래밍 언어의 데이터 타입

		1. Boolean
			bool

		2. String
			string -> 한 번 생성되면 수정될 수 없는 immutable 타입

		3. Integer
			int, int8, int16, int32, int64
			unit, unit8, unit16, unit32, unit64 uintptr

		4. Float 및 복소수
			float32, float64, complex64, complex128

		5. 기타
			byte: unit8과 동일하며 바이트 코드에 사용
			rune int32와 동일하며 유니코드 코드포인트에 사용
	*/
	bool_ := true    // var bool_ bool = true
	str_ := "string" // var str_ string = "string"
	int_ := 100      // var int_ int = 100
	float_ := 0.125
	byte_ := 24
	rune_ := 64

	fmt.Println("Boolean:", bool_, "String:", str_, "Integer:", int_)
	fmt.Println("Float:", float_, "Byte_:", byte_, "Rune:", rune_)

	/*
		문자열

		문자열 리터럴은 ` ` 혹은 " "를 이용하여 표현 가능하다.

		1. Back Quote(` `)
		` `로 둘러 싸인 문자열은 Raw String Literal이라고 부른다.
		문자열은 별도로 해석되지 않고 Raw String 그대로의 값을 가진다.
		예로 "\n"의 경우 NewLine으로 해석되지 않고 문자열 "\n"값을 가진다.
		복수 라인이 문자열을 표현할 때 자주 사용된다.

		2. 이중인용부호(" ")
		" "로 둘러 싸인 문자열은 Interpreted String Literal이라 부른다.bool_
		복수 라인에 걸쳐 쓸 수 없으며, Escape 문자열들은 특별한 의미로 해석된다.
		문자열 안에 "\n"이 있을 경우 NewLine으로 해석된다.
		이중인용부호를 이용해 여러 라인에 걸쳐쓰기 위해서는
		+ 연산자를 이용해 결합하여 사용한다.
	*/

	rawLiteral := `Line 1
	line 2
	Escape -> X | Ex \n
	Test`

	interLiteral := "Esxape -> O | Ex) \n" + "Test"

	fmt.Println(rawLiteral)
	fmt.Println(interLiteral)

	/*
		데이터 타입 변환

		하나의 데이터 타입에서 다른 데이터 타입으로 변환하기 위해서는
		Type(value)과 같이 표현하고 이를 Type Conversion이라 부른다.
		Ex)	float_ = float32(100) 	| float32 -> int
			bytes_ = []byte("ABC") 	| String -> Byte 배열

		Go 에서는 타입간 변환은 명시적으로 지정해 주어야 한다.
		예로 정수형 int에서 uint로 변환할 때 암묵적으로 변환이
		일어나지 않으므로 아래 예시와 같이 반드시 변환을 지정해
		주어야하며 명시적 지정없이 변환을 하면 런타임 에러가 발생한다.bool_
		Ex) var i int = 100; var u uint = uint(i)
	*/
	var i int = 100
	var u uint = uint(i)
	var f float32 = float32(i)
	fmt.Println(u, f)

	var _str string = "ABC"
	var _bytes []byte = []byte(_str)
	var _str2 string = string(_bytes)
	fmt.Println(_bytes, _str2)
}
