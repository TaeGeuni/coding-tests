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

func ValueOfParentheses(s string) int {
	res := 0
	pos := false
	stack := make([]byte, 0)

	check := make(map[byte]int)
	check['('] = 0
	check['['] = 0

	for i := 0; i < len(s); i++ {
		if (check['('] == 0 && s[i] == ')') || (check['['] == 0 && s[i] == ']') {
			return 0
		}
		if s[i] == '(' {
			check['(']++
			pos = true
			stack = append(stack, ')')
		} else if s[i] == '[' {
			check['[']++
			pos = true
			stack = append(stack, ']')
		} else if (s[i] == ')' || s[i] == ']') && stack[len(stack)-1] != s[i] {
			return 0
		} else if s[i] == ')' && pos {
			num2, num3 := 1, 1
			for i := 0; i < check['(']; i++ {
				num2 = num2 * 2
			}
			for i := 0; i < check['[']; i++ {
				num3 = num3 * 3
			}
			res += num2 * num3
			check['(']--
			pos = false
			stack = stack[:len(stack)-1]
		} else if s[i] == ']' && pos {
			num2, num3 := 1, 1
			for i := 0; i < check['(']; i++ {
				num2 = num2 * 2
			}
			for i := 0; i < check['[']; i++ {
				num3 = num3 * 3
			}
			res += num2 * num3
			check['[']--
			pos = false
			stack = stack[:len(stack)-1]
		} else if s[i] == ')' && !pos {
			check['(']--
			stack = stack[:len(stack)-1]
		} else if s[i] == ']' && !pos {
			check['[']--
			stack = stack[:len(stack)-1]
		}

		if i == len(s)-1 && len(stack) != 0 {
			return 0
		}

	}

	return res
}

func main() {
	defer writer.Flush()

	var s string
	fmt.Fscan(reader, &s)

	fmt.Fprintln(writer, ValueOfParentheses(s))

}
