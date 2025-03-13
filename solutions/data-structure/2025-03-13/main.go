package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func ReceiveSignal(tower []int) []int {
	res := make([]int, len(tower))
	stack := make([][2]int, 0) // 스택: (탑의 인덱스, 탑의 높이)

	for i := 0; i < len(tower); i++ {
		// 현재 탑이 이전 탑보다 높으면, 이전 탑 제거
		for len(stack) > 0 && stack[len(stack)-1][1] < tower[i] {
			stack = stack[:len(stack)-1]
		}

		// 신호를 수신하는 탑이 있다면 인덱스를 저장
		if len(stack) > 0 {
			res[i] = stack[len(stack)-1][0] + 1
		} else {
			res[i] = 0
		}

		// 현재 탑을 스택에 추가
		stack = append(stack, [2]int{i, tower[i]})
	}

	return res
}

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	tower := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &tower[i])
	}

	res := ReceiveSignal(tower)

	for i := 0; i < n; i++ {
		fmt.Fprintf(writer, "%d ", res[i])
	}
}
