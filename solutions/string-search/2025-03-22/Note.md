# 팰린드롬 만들기 문제 풀이
👉🏻[문제 링크](https://www.acmicpc.net/problem/1213)

## 문제 설명
임한수와 임문빈은 서로 사랑하는 사이입니다. 임한수는 팰린드롬 문자열을 좋아하기 때문에, 임문빈은 그의 이름으로 팰린드롬을 만들려 합니다. 임문빈이 임한수의 영어 이름의 알파벳 순서를 적절히 바꿔 팰린드롬을 만들 수 있도록 도와주는 프로그램을 작성하세요.

### 입력 조건
- 첫째 줄에 임한수의 영어 이름이 주어집니다.
- 이름은 알파벳 대문자로만 이루어져 있으며, 최대 50글자입니다.

### 출력 조건
- 팰린드롬을 만들 수 있으면, 사전순으로 가장 앞서는 팰린드롬을 출력합니다.
- 만들 수 없다면 `I'm Sorry Hansoo`를 출력합니다.

## 예제
### 입력 1
```
AABB
```
### 출력 1
```
ABBA
```

### 입력 2
```
AAABB
```
### 출력 2
```
ABABA
```

### 입력 3
```
ABACABA
```
### 출력 3
```
AABCBAA
```

### 입력 4
```
ABCD
```
### 출력 4
```
I'm Sorry Hansoo
```

## 풀이 과정

### 문제 핵심
팰린드롬이란 앞에서 읽든 뒤에서 읽든 같은 문자열입니다. 이를 만들기 위해서는 다음 조건이 만족되어야 합니다:

- **문자열의 길이가 짝수인 경우**: 모든 문자의 개수가 짝수여야 팰린드롬을 만들 수 있습니다.
- **문자열의 길이가 홀수인 경우**: 단 한 개의 문자만 개수가 홀수여야 하고, 나머지는 짝수여야 합니다. 홀수 개수의 문자는 중앙에 위치합니다.

### 알고리즘 설명
1. 문자열을 순회하며 각 알파벳의 등장 횟수를 센다.
2. 알파벳 리스트를 사전순으로 정렬한다. (사전순 출력을 위해)
3. 홀수 개수의 알파벳 개수를 파악한다.
   - 짝수 길이 문자열에서 홀수 개수 문자가 하나라도 있으면 실패.
   - 홀수 길이 문자열에서 홀수 개수 문자가 2개 이상이면 실패.
4. 실패 조건이 아니면, 각 문자의 개수를 절반으로 나누어 앞 절반 문자열을 만든다.
5. 이 문자열을 뒤집어서 뒤 절반으로 만들고, (홀수 개수 문자라면 가운데에 삽입)
6. 완성된 팰린드롬을 출력한다.

### 시간 복잡도 분석
- 입력 길이: `n` (최대 50)
- 문자 카운트: `O(n)`
- 정렬: `O(26 log 26)` → 상수 시간 (알파벳 개수는 고정)
- 결과 문자열 생성: `O(n)`

따라서 전체 시간 복잡도는 **O(n + k log k)** 입니다.

## 구현 코드

```go
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
```

## 결론
- 문제 해결을 위해 각 알파벳의 개수를 세고 조건을 만족하는지 판단한 뒤 팰린드롬을 생성했습니다.
- 팰린드롬이 가능한 조건만 잘 판단하면 구현은 어렵지 않으며, 시간 복잡도도 매우 효율적입니다.
- 사전순 출력을 위해 정렬을 반드시 수행해야 하며, 이는 고정 크기(26개) 내에서 수행되므로 부담이 없습니다.
