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
}
