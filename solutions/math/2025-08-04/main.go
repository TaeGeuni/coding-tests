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

type Point struct {
	x int
	y int
}

func ccw(p1, p2, p3 Point) int {
	res := (p2.x-p1.x)*(p3.y-p1.y) - (p3.x-p1.x)*(p2.y-p1.y)

	if res > 0 {
		return 1
	} else if res < 0 {
		return -1
	}
	return 0
}

func main() {
	defer writer.Flush()

	var p1, p2, p3 Point

	fmt.Fscan(reader, &p1.x, &p1.y)
	fmt.Fscan(reader, &p2.x, &p2.y)
	fmt.Fscan(reader, &p3.x, &p3.y)

	fmt.Fprintln(writer, ccw(p1, p2, p3))

}
