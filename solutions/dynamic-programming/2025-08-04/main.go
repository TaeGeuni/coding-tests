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

type value struct {
	count   int
	number  int
	history []int
}

func make1(n int) (int, []int) {

	if n == 1 {
		return 0, []int{1}
	}

	queue := make([]value, 0)
	queue = append(queue, value{count: 0, number: n, history: []int{n}})

	visited := make(map[int]bool)
	visited[n] = true

	var res value

	for len(queue) > 0 {
		node := queue[0]
		var v value

		if node.number%3 == 0 && !visited[node.number/3] {
			v.count = node.count + 1
			v.number = node.number / 3
			newHistory := make([]int, len(node.history))
			copy(newHistory, node.history)
			newHistory = append(newHistory, v.number)
			v.history = newHistory

			visited[v.number] = true
			queue = append(queue, v)

			if v.number == 1 {
				res = v
				break
			}

		}
		if node.number%2 == 0 && !visited[node.number/2] {
			v.count = node.count + 1
			v.number = node.number / 2
			newHistory := make([]int, len(node.history))
			copy(newHistory, node.history)
			newHistory = append(newHistory, v.number)
			v.history = newHistory

			visited[v.number] = true
			queue = append(queue, v)

			if v.number == 1 {
				res = v
				break
			}

		}
		if node.number-1 >= 1 && !visited[node.number-1] {
			v.count = node.count + 1
			v.number = node.number - 1
			newHistory := make([]int, len(node.history))
			copy(newHistory, node.history)
			newHistory = append(newHistory, v.number)
			v.history = newHistory

			visited[v.number] = true
			queue = append(queue, v)

			if v.number == 1 {
				res = v
				break
			}

		}
		queue = queue[1:]
	}

	return res.count, res.history

}

func main() {
	defer writer.Flush()

	n := 0
	fmt.Fscan(reader, &n)

	count, history := make1(n)
	fmt.Fprintln(writer, count)
	for i := 0; i < len(history); i++ {
		fmt.Fprintf(writer, "%d ", history[i])
	}
}
