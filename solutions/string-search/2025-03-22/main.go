package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

// 팰린드롬 생성 함수
func Palindrome(name string, alphabetList []string, howManyAlphabet map[string]int) string {
	res := ""

	if len(name)%2 != 0 {
		// 홀수 길이: 홀수 개수 문자가 정확히 1개여야 함
		check := 0
		v := ""
		revers := ""
		for i := 0; i < len(alphabetList); i++ {
			if howManyAlphabet[alphabetList[i]]%2 != 0 {
				check++
				v = alphabetList[i] // 중앙에 올 문자를 저장
			}
			howManyAlphabet[alphabetList[i]] /= 2 // 절반만 사용
		}
		if check != 1 {
			return "I'm Sorry Hansoo"
		}
		for i := 0; i < len(alphabetList); i++ {
			res += strings.Repeat(alphabetList[i], howManyAlphabet[alphabetList[i]])
		}
		revers = reverse(res)
		res += v + revers
	} else {
		// 짝수 길이: 모든 알파벳의 개수가 짝수여야 함
		check := 0
		revers := ""
		for i := 0; i < len(alphabetList); i++ {
			if howManyAlphabet[alphabetList[i]]%2 != 0 {
				check++
			}
			howManyAlphabet[alphabetList[i]] /= 2
		}
		if check != 0 {
			return "I'm Sorry Hansoo"
		}
		for i := 0; i < len(alphabetList); i++ {
			res += strings.Repeat(alphabetList[i], howManyAlphabet[alphabetList[i]])
		}
		revers = reverse(res)
		res += revers
	}
	return res
}

// 문자열 뒤집기 함수
func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	defer writer.Flush()

	var name string
	fmt.Fscan(reader, &name)

	alphabetList := make([]string, 0)
	howManyAlphabet := make(map[string]int)

	for i := 0; i < len(name); i++ {
		char := string(name[i])
		if howManyAlphabet[char] == 0 {
			alphabetList = append(alphabetList, char)
		}
		howManyAlphabet[char]++
	}

	sort.Slice(alphabetList, func(i, j int) bool {
		return alphabetList[i] < alphabetList[j]
	})

	fmt.Fprintln(writer, Palindrome(name, alphabetList, howManyAlphabet))
}
