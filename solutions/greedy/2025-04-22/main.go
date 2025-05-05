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

func TurnOver(n, m int, coin [][]byte) int {
	res := 0

	allZero := false

	for !allZero {
		x, y := 0, 0
		find := false
		for i := n - 1; i >= 0; i-- {
			for j := m - 1; j >= 0; j-- {
				if coin[i][j] == '1' {
					x, y = i, j
					find = true
				}
				if find {
					break
				}
			}
			if find {
				break
			}
		}

		if find {
			res++
			for i := 0; i <= x; i++ {
				for j := 0; j <= y; j++ {
					if coin[i][j] == '1' {
						coin[i][j] = '0'
					} else {
						coin[i][j] = '1'
					}
				}
			}
		} else {
			allZero = true
		}

	}

	return res
}

func main() {
	defer writer.Flush()

	var n, m int
	fmt.Fscan(reader, &n, &m)
	coin := make([][]byte, n)

	for i := 0; i < n; i++ {
		input := ""
		fmt.Fscan(reader, &input)
		for j := 0; j < m; j++ {
			coin[i] = append(coin[i], input[j])
		}
	}
	res := TurnOver(n, m, coin)
	fmt.Fprintln(writer, res)

}
