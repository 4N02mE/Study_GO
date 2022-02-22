package main

/*
package main -> Entry Point (진입점)
코드 실행 시 package main -> func main() 순서로 찾아내어 실행한다.
*/

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
	/*
		fmt -> 기본적으로 제공하는 패키지
		fmt.Println(): stdout으로 표준 문자열 출력을 제공하는 함수
		println(): 내장된 함수로 stderr를 출력해주는 내장함수
		           때문에 주로 디버깅할 때에 유용
	*/
}
