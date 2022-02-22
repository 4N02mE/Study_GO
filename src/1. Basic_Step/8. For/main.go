package main

import "fmt"

func main() {
	/*
		Go 프로그래밍 언어에서 반목문은 for문 하나 밖에 없다.
		for 초기값; 조건식; 증감 {...} 의 형식을 따르며
		경우에 따라 초기값, 조건식, 증감식 등은 생략이 가능하다.
	*/
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)

	/*
		Go에서 for루프는 초기값과 증감식을 생략하고 조건식만 사용할 수 있다.
		다른 언어의 while루프와 같이 쓰이도록 한다.
	*/
	n := 1
	for n < 100 {
		n *= 2
		//
	}
	fmt.Println(n)

	/*
		for루프로 무한루프를 만들려면 초기값, 조건식, 증감을 모두 생략하면 된다.
		무한루프 중 빠져나오기 위해선 Ctrl+C를 누르면 된다.
	*/
	x := 0
	for {
		x++
		println(x)
		if x > 5 {
			break
		}
	}

	/*
		for range문은 컬렉션에서 한 요소씩 가져와 차례대로
		for문의 문장들을 실행한다. 타 언어의 foreach와 비슷한 용법

		for range문은 for 인덱스, 요소값 := range 컬렉션 같이 for 루프를
		구성하고 컬렉션에서 요소와 요소의 위치 인덱스 값을
		각각 for 키워드 이후 2개의 변수에 할당한다.
	*/
	names := []string{"gragas", "gnar", "sion"}
	for index, name := range names {
		fmt.Println("index: ", index, "| name: ", name)
	}

	/*
		for 루프 내에서 빠져나오기 위해 사용하는 것들이 있다.
		즉시 빠져나오기 위해서는 break,
		루프 시작부분으로 바로 돌아가기 위해서는 continue,
		임의의 문장으로 이동하기 위해서는 goto를 사용할 수 있다.

		goto는 for 루프와 관련없이 사용될 수 있다.
		break문은 for문 이외에 switch문이나 select문에서도 사용할 수 있다.
		continue문은 for문과 연관되어 사용한다.
	*/
	var a int = 1
	for a < 15 {
		if a == 5 {
			a += a
			continue
		}
		a++
		if a > 10 {
			break
		}
	}
	if a == 11 {
		goto END
	}
	fmt.Println(a)
END:
	fmt.Println("End")

	/*
		break문은 보통 단독으로 사용되지만 경우에 따라 "break 레이블"과 같이
		사용하여 지정된 레이블로 이동할 수 있다.
		break의 "레이블"은 보통 현재의 for 루프를 바로 위에 적게 되는데
		이러한 "break 레이블"은 현재의 루프를 빠져나와 지정된 레이블로 이동하고
		break문의 직속 for루프 전체의 다음 문장을 실행하게 한다.
	*/
	y := 0
L1:
	for {
		if y == 0 {
			break L1
		}
	}
	fmt.Println("OK")
}
