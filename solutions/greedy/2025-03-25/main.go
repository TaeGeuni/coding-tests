package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

func FindMin(k int, sensors []int) int {
	res := 0
	if len(sensors) == 1 {
		return res
	}
	distance := make([]int, 0)

	sort.Ints(sensors)

	for i := 1; i < len(sensors); i++ {
		distance = append(distance, sensors[i]-sensors[i-1])
	}

	sort.Ints(distance)

	if k > 1 {
		distance = distance[:len(distance)-k+1]
	}

	for i := 0; i < len(distance); i++ {
		res += distance[i]
	}

	return res
}

func main() {
	defer writer.Flush()

	var n, k int
	fmt.Fscan(reader, &n)
	fmt.Fscan(reader, &k)

	sensors := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &sensors[i])
	}

	fmt.Fprintln(writer, FindMin(k, sensors))

}
