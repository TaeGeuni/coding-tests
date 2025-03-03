package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

// generatePascalsTriangle 함수는 주어진 R과 W를 기준으로 파스칼 삼각형을 생성한다.
// 삼각형의 각 요소는 (i, j) 좌표를 키로 하는 map에 저장된다.
func generatePascalsTriangle(R, W int) map[[2]int]int {
	pascMap := make(map[[2]int]int)
	for i := 1; i < R+W; i++ {
		for j := 1; j <= i; j++ {
			// 삼각형의 양 끝 값은 항상 1
			if j == 1 || j == i {
				pascMap[[2]int{i, j}] = 1
			} else {
				// 내부 값은 위의 두 수의 합
				pascMap[[2]int{i, j}] = pascMap[[2]int{i - 1, j - 1}] + pascMap[[2]int{i - 1, j}]
			}
		}
	}
	return pascMap
}

// sumTriangle 함수는 주어진 R, C, W에 해당하는 부분 삼각형의 합을 계산한다.
func sumTriangle(pascMap map[[2]int]int, R, C, W int) int {
	res := 0
	count := C
	for i := R; i < R+W; i++ {
		for j := C; j <= count; j++ {
			res += pascMap[[2]int{i, j}]
		}
		count++ // 다음 행에서 더할 요소가 하나씩 증가
	}
	return res
}

func main() {
	defer writer.Flush()

	var R, C, W int
	fmt.Fscan(reader, &R, &C, &W)

	// 파스칼 삼각형을 생성
	pascMap := generatePascalsTriangle(R, W)

	// 결과 출력
	fmt.Fprintln(writer, sumTriangle(pascMap, R, C, W))
}
