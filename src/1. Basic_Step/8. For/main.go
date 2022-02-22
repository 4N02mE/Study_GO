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
}
