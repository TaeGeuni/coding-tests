package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

func reverse(r, result []byte) []byte {
	for i := len(r) - 1; i >= 0; i-- {
		result = append(result, r[i])
	}
	return result
}

func WordReverse(s string) string {
	result := make([]byte, 0)
	r := make([]byte, 0)
	bracket := false

	for i := 0; i < len(s); i++ {
		if s[i] == '<' {
			result = reverse(r, result)
			r = r[:0]
			bracket = true
		}
		if bracket {
			result = append(result, s[i])
		} else if !bracket && s[i] == ' ' {
			result = reverse(r, result)
			result = append(result, ' ')
			r = r[:0]
		} else {
			r = append(r, s[i])
		}
		if s[i] == '>' {
			bracket = false
		}
		if i == len(s)-1 {
			result = reverse(r, result)
			r = r[:0]
		}
	}
	return string(result)
}

func main() {
	defer writer.Flush()

	var s string
	s, _ = reader.ReadString('\n')
	s = strings.Trim(s, string('\n'))

	fmt.Fprintln(writer, WordReverse(s))
}
