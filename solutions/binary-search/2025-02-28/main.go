package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

	// haveCableNum: 현재 가지고 있는 랜선 개수, needCableNum: 필요한 랜선 개수
	var haveCableNum, needCableNum, max int

	fmt.Fscan(reader, &haveCableNum, &needCableNum)
	cable := make([]int, haveCableNum)

	for i := 0; i < haveCableNum; i++ {
		fmt.Fscan(reader, &cable[i])
		if cable[i] > max {
			max = cable[i] // 최대 길이 저장
		}
	}

	fmt.Fprintln(writer, FindMaximumCable(max, needCableNum, cable))
}

// 이분 탐색을 사용하여 최대 길이를 찾음
func FindMaximumCable(max, need int, cable []int) int {
	res := 0
	left, right := 1, max

	for left <= right {
		mid := (left + right) / 2
		total := 0

		// 주어진 길이로 만들 수 있는 랜선 개수를 계산하는 반복문
		for i := 0; i < len(cable); i++ {
			total += cable[i] / mid
		}
		if total >= need {
			res = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return res
}
