package main

import (
	"bufio"
	"fmt"
	"os"
)

type MoveAndTruck struct {
	truckNumber int
	move        int
}

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

func CrossBridge(n, w, l int, truck []int) int {
	res := 0
	nowWeight := 0
	next := 0
	end := false

	bridge := make([]MoveAndTruck, 0)
	bridge = append(bridge, MoveAndTruck{truckNumber: next, move: 0})
	nowWeight += truck[next]
	if next == n-1 {
		end = true
	} else {
		next++
	}

	for len(bridge) > 0 {
		for i := 0; i < len(bridge); i++ {
			bridge[i].move++
		}
		if bridge[0].move > w {
			nowWeight -= truck[bridge[0].truckNumber]
			bridge = bridge[1:]
		}

		if !end {
			if nowWeight+truck[next] <= l {
				bridge = append(bridge, MoveAndTruck{truckNumber: next, move: 1})
				nowWeight += truck[next]
				if next == n-1 {
					end = true
				} else {
					next++
				}
			}
		}
		res++
	}

	return res
}

func main() {
	defer writer.Flush()

	var n, w, l int
	fmt.Fscan(reader, &n, &w, &l)

	truck := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &truck[i])
	}

	result := CrossBridge(n, w, l, truck)

	fmt.Fprintln(writer, result)
}
