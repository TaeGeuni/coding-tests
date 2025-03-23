package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)
	balloons := make([][]int, n+1)

	for i := 1; i < n+1; i++ {
		v := 0
		fmt.Fscan(reader, &v)
		balloons[i] = append(balloons[i], v)
	}

	start := 1
	move := 0
	res := make([]int, 0)

	for len(res) < n {
		if move == 0 {
			res = append(res, start)
			move = balloons[start][0]
			balloons[start][0] = 0

		} else if move > 0 {
			if start == n {
				start = 1
			} else {
				start++
			}

			if balloons[start][0] != 0 {
				move--
			}

		} else if move < 0 {
			if start == 1 {
				start = n
			} else {
				start--
			}

			if balloons[start][0] != 0 {
				move++
			}
		}
	}

	ans := strings.Trim(fmt.Sprint(res), "[]")
	fmt.Fprintln(writer, ans)
}
