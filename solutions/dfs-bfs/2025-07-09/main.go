package main

import (
	"bufio"
	"fmt"
	"os"
)

const max = 100_000

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

func hideAndSeek(n, k int) int {
	visited := make([]bool, max+1)
	queue := make([][2]int, 0)
	queue = append(queue, [2]int{n, 0})
	visited[n] = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		pos, time := current[0], current[1]
		if pos == k {
			return time
		}

		next := [3]int{pos - 1, pos + 1, pos * 2}

		for i := 0; i < 3; i++ {
			if next[i] >= 0 && next[i] <= max && !visited[next[i]] {
				queue = append(queue, [2]int{next[i], time + 1})
				visited[next[i]] = true
			}
		}

	}

	return -1
}

func main() {
	defer writer.Flush()

	n, k := 0, 0
	fmt.Fscan(reader, &n, &k)
	fmt.Fprintln(writer, hideAndSeek(n, k))
}
