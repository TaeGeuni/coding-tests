package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

	var n int
	fileExtension := make(map[string]int)
	files := make([]string, 0)

	fmt.Fscan(reader, &n)

	for i := 0; i < n; i++ {
		input := ""
		fmt.Fscan(reader, &input)
		fileName := strings.Split(input, ".")
		if fileExtension[fileName[1]] == 0 {
			files = append(files, fileName[1])
		}
		fileExtension[fileName[1]]++
	}

	sort.Slice(files, func(i, j int) bool {
		return files[i] < files[j]
	})

	for i := 0; i < len(files); i++ {
		fmt.Fprintln(writer, files[i], fileExtension[files[i]])
	}

}
