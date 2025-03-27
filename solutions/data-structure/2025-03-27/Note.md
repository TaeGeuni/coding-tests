# 문자열 단어 뒤집기 문제 풀이
👉🏻[문제 링크](https://www.acmicpc.net/problem/17413)

## 문제 설명
주어진 문자열에서 단어를 뒤집는 프로그램을 작성합니다. 단, 태그(`<`, `>`)로 둘러싸인 부분은 뒤집지 않고 그대로 유지하며, 공백으로 구분된 단어만 뒤집습니다.

### 입력 조건
- 문자열은 알파벳 소문자(`a-z`), 숫자(`0-9`), 공백(` `), 특수문자(`<`, `>`)로 구성됩니다.
- 문자열의 시작과 끝은 공백이 없습니다.
- `<`와 `>`는 번갈아 나타나며, 태그 내부에는 알파벳 소문자와 공백만 존재합니다.

### 출력 조건
- 태그 내부는 그대로 유지하고, 태그 외부의 단어만 뒤집습니다.

---

## 코드 구현

### 코드
```go
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
```

---

## 코드 설명

### 주요 로직
1. **`reverse` 함수**
   - 입력된 슬라이스(`r`)를 뒤집어 `result`에 추가합니다.

2. **`WordReverse` 함수**
   - `result`: 최종 결과를 저장하는 바이트 슬라이스.
   - `r`: 단어를 임시로 저장하는 바이트 슬라이스.
   - `bracket`: 태그 내부 여부를 나타내는 플래그.

3. **문자열 처리**
   - `<`를 만나면:
     - 현재까지의 단어(`r`)를 뒤집어 `result`에 추가.
     - `bracket` 플래그를 `true`로 설정.
   - `>`를 만나면:
     - `bracket` 플래그를 `false`로 설정.
   - 공백을 만나면:
     - 현재까지의 단어를 뒤집어 `result`에 추가.
     - 공백 문자도 `result`에 추가.
   - 마지막 문자까지 처리한 후, 남은 단어를 뒤집어 추가.

---

## 예제 실행

### 입력 1
```
baekjoon online judge
```

### 출력 1
```
noojkeab enilno egduj
```

### 입력 2
```
<int><max>2147483647<long long><max>9223372036854775807
```

### 출력 2
```
<int><max>7463847412<long long><max>7085774586302733229
```

---

## 주요 특징
- **효율적인 메모리 사용**: `[]byte`를 사용하여 문자열 덧셈(`+=`)의 비효율성을 해결.
- **태그 처리**: `<`와 `>`로 둘러싸인 부분을 정확히 유지.
- **단어 뒤집기**: 공백 기준으로 단어를 분리하고, `reverse` 함수를 통해 뒤집기.

---

## 시간 및 공간 복잡도
- **시간 복잡도**: O(n)
  - 문자열을 한 번 순회하며, 단어를 뒤집는 작업이 O(n)에 해당.
- **공간 복잡도**: O(n)
  - `result`와 `r`에 문자열 데이터를 저장하므로, 추가 공간 사용은 입력 문자열의 길이에 비례.

---

## 결론
이 코드는 태그와 단어를 정확히 구분하며 효율적으로 처리합니다. 문자열 덧셈 대신 `[]byte`를 사용해 메모리 사용량과 속도 면에서 최적화된 방식으로 동작합니다.

