# 안정적인 문자열
👉🏻[문제 링크](https://www.acmicpc.net/problem/4889)

## 문제 설명

여는 괄호 `{`와 닫는 괄호 `}`만으로 이루어진 문자열이 주어진다. 여기서 **안정적인 문자열**을 만들기 위한 최소 연산의 수를 구하려고 한다.

### 안정적인 문자열의 정의:

- 빈 문자열은 안정적이다.
- `S`가 안정적이라면, `{S}`도 안정적인 문자열이다.
- `S`와 `T`가 안정적이라면, `ST`(두 문자열의 연결)도 안정적이다.

예: `{}`, `{}{}{}`, `{{{}}}` 는 안정적이다. 하지만 `}{`, `{{}{`, `{}{` 는 안정적이지 않다.

### 연산:

문자열을 안정적으로 만들기 위해 할 수 있는 연산은 다음 두 가지 중 하나이다.

1. 여는 괄호를 닫는 괄호로 바꾸기
2. 닫는 괄호를 여는 괄호로 바꾸기

## 입력

- 한 줄에 하나씩 여는 괄호와 닫는 괄호로 이루어진 문자열이 주어진다.
- 문자열의 길이는 2000자를 넘지 않으며, 항상 길이는 짝수이다.
- 마지막 줄은 하이픈(`-`)이 한 개 이상 주어진다.

## 출력

각 테스트 케이스마다, 테스트 케이스 번호와 입력으로 주어진 문자열을 안정적으로 바꾸는 데 필요한 최소 연산의 수를 출력한다.

## 예제 입력

```
}{
{}{}{}
{{{}
---
```

## 예제 출력

```
1. 2
2. 0
3. 1
```

---

## 해결 방법 (Go 코드)

```go
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

func isThisStable(s string) int {
	res := 0
	count := 0

	for i := 0; i < len(s); i++ {
		if count == 0 && s[i] == '{' {
			count++
		} else if count == 0 && s[i] == '}' {
			res++
			count++
		} else if len(s)-i == count && s[i] == '{' {
			res++
			count--
		} else if len(s)-i == count && s[i] == '}' {
			count--
		} else if s[i] == '{' {
			count++
		} else if s[i] == '}' {
			count--
		} else if s[i] == '-' {
			return -1
		}
	}

	return res
}

func main() {
	defer writer.Flush()

	s := ""
	count := 0

	for {
		fmt.Fscan(reader, &s)

		n := isThisStable(s)

		if n < 0 {
			break
		}
		count++
		fmt.Fprintf(writer, "%d. %d\n", count, n)
	}
}
```

## 설명

- 괄호 문자열을 탐색하면서 필요한 수정을 세는 방식.
- 잘못된 닫는 괄호가 먼저 등장하는 경우를 고려하여 바꿔야 할 수를 센다.
- 전체 문자열을 스캔하며 안정적인 상태가 유지되도록 `count`로 추적.