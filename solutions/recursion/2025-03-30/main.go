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

func Draw(n, x, y int, canvas [][]string) {
	if n == 1 {
		canvas[x][y] = "*"
		return
	}

	ySize := (n * 4) - 3
	xSize := ySize + 2

	maxX, maxY := (xSize + x), (ySize + y)

	for i := y; i < maxY; i++ {
		canvas[x][i] = "*"
		canvas[maxX-1][i] = "*"
	}
	for i := x; i < maxX; i++ {
		canvas[i][y] = "*"
		if i == x+1 {
			canvas[i+1][maxY-2] = "*"
		} else {
			canvas[i][maxY-1] = "*"
		}
	}

	if n == 2 {
		canvas[x+3][y+2] = "*"
		canvas[x+4][y+2] = "*"
	}

	Draw(n-1, x+2, y+2, canvas)

}

func main() {
	defer writer.Flush()

	var n, xSize, ySize int
	fmt.Fscan(reader, &n)
	ySize = (n * 4) - 3
	if n == 1 {
		xSize = 1
	} else {
		xSize = ySize + 2
	}
	canvas := make([][]string, xSize)
	for i := 0; i < xSize; i++ {
		canvas[i] = make([]string, ySize)
	}
	for i := 0; i < xSize; i++ {
		for j := 0; j < ySize; j++ {
			canvas[i][j] = " "
		}
	}
	Draw(n, 0, 0, canvas)
	for i := 0; i < xSize; i++ {
		if i == 1 {
			fmt.Fprintln(writer, canvas[i][0])
			continue
		}
		for j := 0; j < ySize; j++ {
			fmt.Fprint(writer, canvas[i][j])
		}
		fmt.Fprintln(writer)
	}
}
