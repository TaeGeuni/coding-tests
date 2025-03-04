package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

	var N, L, res int
	fmt.Fscan(reader, &N, &L)
	pipe := make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Fscan(reader, &pipe[i])
	}

	// 물이 새는 위치를 정렬
	sort.Slice(pipe, func(i, j int) bool {
		return pipe[i] < pipe[j]
	})

	lastTape := 0 // 마지막으로 붙인 테이프의 끝 위치

	for i := 0; i < N; i++ {
		// 현재 위치가 마지막 테이프의 커버 범위를 벗어나면 새 테이프 사용
		if pipe[i] > lastTape {
			lastTape = pipe[i] + L - 1
			res++
		}
	}

	fmt.Fprintln(writer, res)
}
