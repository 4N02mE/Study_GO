package main

import "fmt"

func main() {
	/*
		여러 값을 비교해야 하는 경우 혹은 다수의 조건식을
		체크해야 하는 경우 switch문을 사용한다.

		switch문 뒤에 하나의 변수(혹은 Expression)을
		지정하고, case문에 해당 변수가 가질 수 있는 값들을 지정한다.
		같은 값의 case문이 있을 경우 묶어서 작성이 가능하다.
	*/

	var age int = 10
	switch age {
	case 7:
		{
			fmt.Println("7")
		}
	case 8, 9:
		{
			fmt.Println("8 or 9")
		}
	case 10:
		{
			fmt.Println("10")
		}
	default:
		{
			fmt.Println("11~")
		}
	}

	/*
		Expression을 사용한 경우
		var y = 10
		switch x := y + 2; x * 3 {
			...
		}
	*/

	/*
		Go의 단순한 Switch문 용법은 타 언어와 비슷하지만
		다른 언어와 다른 Go만의 특별한 용법이 있다.

		1. Switch 뒤에 Expression이 없을 수 있음
			타 언어는 switch 키워드 뒤에 변수나 Expression이 반드시 요하지만
			Go는 이를 true로 생각하고 첫 번째 case문으로 이동하여 검사한다.

		2. case문에 Expression을 쓸 수 있음
			타 언어의 case문은 일반적인 리터럴 값만을 갖지만
			Go는 case문에 복잡한 Expression을 가질 수 있다.

		3. No default fall Through
			break문을 사용하지 않으면 다음 case문을 실행시키는 타 언어와 달리
			Go는 다음 case로 이동하지 않는다.
			이는 Go 컴파일러가 자동으로 break문을 각 case문에 추가하여
			break문의 유무에 상관없이 break를 수행한다.
			타 언어와 같이 다음 문장으로 넘어가게 하려 한다면
			fallthrough를 명시하면 타 언어와 같이 다음 case도 수행하게 된다.

		4. Type switch
			타 언어의 switch는 일반적으로 변수의 값을 기준으로 하여
			case로 분기하지만, Go는 그 변수의 Type에 따라 case를 분기할 수 있음
	*/

	// 1. Switch 뒤에 Expression이 없을 수 있음
	// 2. case문에 Expression을 쓸 수 있음
	var score int = 85
	switch {
	case score >= 90:
		fmt.Println("A")
	case score >= 80:
		fmt.Println("B")
	default:
		fmt.Print("C")
	}

	// 3. No default fall through
	var x int = 3
	switch {
	case x >= 1:
		fmt.Println("1-1")
	case x >= 2:
		fmt.Println("1-2")
	case x >= 3:
		fmt.Println("1-3")
	default:
		fmt.Println("None")
	}

	switch {
	case x >= 1:
		fmt.Println("2-1")
		fallthrough
	case x >= 2:
		fmt.Println("2-2")
		fallthrough
	case x >= 3:
		fmt.Println("2-3")
		fallthrough
	default:
		fmt.Println("None")
	}

	// 4. Type Switch
	typeSwitch(10)
	typeSwitch("str_")
	typeSwitch(true)
}

func typeSwitch(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("unknown")
	}
}
