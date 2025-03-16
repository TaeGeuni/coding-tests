# 늑대와 올바른 단어
👉🏻[문제 링크](https://www.acmicpc.net/problem/13022)

## 문제 설명
늑대 나라에서 사용하는 올바른 단어의 규칙은 다음과 같다.

1. 임의의 양의 정수 `n`에 대해서, `w`가 `n`번 나온 후, `o`가 `n`번, `l`이 `n`번, `f`가 `n`번 나오는 단어는 올바른 단어이다.
2. 올바른 단어 두 개를 이은 단어도 올바른 단어이다.
3. 위 두 조건으로 만들 수 있는 단어만 올바른 단어이다.

### 예제
#### 올바른 단어 예시
- "wolf"
- "wwoollff"
- "wwwooolllfff"
- "wolfwwoollff"
- "wolfwwoollffwolf"

#### 올바르지 않은 단어 예시
- "wfol" (순서가 올바르지 않음)
- "wwolfolf" (중간에 다른 문자열이 들어감)
- "wwwoolllfff" (`o`의 개수가 다름)

## 입력
- 단어는 `w`, `o`, `l`, `f`로만 이루어져 있으며, 길이는 50을 넘지 않는다.

## 출력
- 입력으로 주어진 단어가 올바른 단어이면 `1`, 아니면 `0`을 출력한다.

## 코드 구현
```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func Ookami(input string) bool {
	var w, o, l, f int

	for i := 0; i < len(input); i++ {
		switch {
		case input[i] == 'w':
			if o == 0 && l == 0 && f == 0 {
				w++
			} else {
				return false
			}
		case input[i] == 'o':
			if w > o && l == 0 && f == 0 {
				o++
			} else {
				return false
			}
		case input[i] == 'l':
			if w == o && l < w && f == 0 {
				l++
			} else {
				return false
			}
		case input[i] == 'f':
			if w == o && w == l && f < w {
				f++
			} else {
				return false
			}
			if w == o && w == l && w == f {
				w, o, l, f = 0, 0, 0, 0
			}
		}
	}

	return w == 0 && o == 0 && l == 0 && f == 0
}

func main() {
	defer writer.Flush()

	var input string
	fmt.Fscan(reader, &input)

	if Ookami(input) {
		fmt.Fprintln(writer, 1)
	} else {
		fmt.Fprintln(writer, 0)
	}
}
```

## 코드 설명
- `Ookami(input string) bool` 함수는 입력된 단어가 올바른 단어인지 판별한다.
- `w, o, l, f` 변수는 각각 문자 `w`, `o`, `l`, `f`의 개수를 추적한다.
- 순서대로 `w → o → l → f`가 나타나야 하며, 각 문자 그룹의 개수가 같아야 한다.
- `f`까지 조건이 만족되면 `w, o, l, f`를 초기화하여 새로운 올바른 단어를 인식할 수 있도록 한다.
- 마지막에 모든 변수가 `0`이 되어야 올바른 단어로 판별된다.

## 예제 실행
### 입력 1
```
wolf
```
### 출력 1
```
1
```
### 입력 2
```
wwolfolf
```
### 출력 2
```
0
```