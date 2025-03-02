package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var reader = bufio.NewReader(os.Stdin)  // 표준 입력을 빠르게 읽기 위한 버퍼 리더 생성
var writer = bufio.NewWriter(os.Stdout) // 표준 출력을 빠르게 출력하기 위한 버퍼 라이터 생성

func main() {
	defer writer.Flush() // main 함수 종료 시 출력 버퍼 비우기

	var guitarNum int
	fmt.Fscan(reader, &guitarNum) // 기타 개수 입력 받기

	serialNumbers := make([]string, guitarNum) // 시리얼 번호를 저장할 문자열 슬라이스 생성

	for i := 0; i < guitarNum; i++ {
		fmt.Fscan(reader, &serialNumbers[i]) // 시리얼 번호 입력 받기
	}

	sort.Slice(serialNumbers, func(i, j int) bool {
		// 1. 길이가 짧은 것이 먼저 온다
		if len(serialNumbers[i]) != len(serialNumbers[j]) {
			return len(serialNumbers[i]) < len(serialNumbers[j])
		}

		// 2. 숫자의 합을 계산하여 비교
		a := 0
		b := 0

		for m := 0; m < len(serialNumbers[i]); m++ {
			data := string(serialNumbers[i][m])
			num, _ := strconv.Atoi(data) // 숫자인 경우 변환, 알파벳은 변환 실패하여 0
			a += num
		}
		for m := 0; m < len(serialNumbers[j]); m++ {
			data := string(serialNumbers[j][m])
			num, _ := strconv.Atoi(data) // 숫자인 경우 변환, 알파벳은 변환 실패하여 0
			b += num
		}
		// 숫자의 합이 작은 것이 먼저 온다
		if a != b {
			return a < b
		}
		// 3. 사전순으로 비교
		return serialNumbers[i] < serialNumbers[j]

	})

	for i := 0; i < guitarNum; i++ {
		fmt.Fprintln(writer, serialNumbers[i])
	}
}
