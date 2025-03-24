package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

func DrawStarDot(num, x, y int, stars [][]string) {
	if num == 1 {
		stars[x][y] = "*"
		return
	}

	size := 1 + (4 * (num - 1))
	endx, endy := (x + size), (y + size)

	// 테두리 그리기
	for i := y; i < endy; i++ {
		stars[x][i] = "*"
		stars[endx-1][i] = "*"
	}
	for i := x; i < endx; i++ {
		stars[i][y] = "*"
		stars[i][endy-1] = "*"
	}

	// 내부 재귀 호출
	DrawStarDot(num-1, x+2, y+2, stars)
}

func main() {
	defer writer.Flush()

	var num int
	fmt.Fscan(reader, &num)

	size := 1 + (4 * (num - 1))
	stars := make([][]string, size)
	for i := 0; i < size; i++ {
		stars[i] = make([]string, size)
		for j := 0; j < size; j++ {
			stars[i][j] = " "
		}
	}

	DrawStarDot(num, 0, 0, stars)

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			fmt.Fprint(writer, stars[i][j])
		}
		fmt.Fprintln(writer)
	}
}
