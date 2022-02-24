package main

import "os"

func main() {
	// 파일 열기
	f_input, err := os.Open("f1.txt")
	if err != nil {
		panic(err)
	}

	defer f_input.Close()
}
