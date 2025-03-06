# 쇠막대기와 레이저 문제 풀이
👉🏻[문제 링크](https://www.acmicpc.net/problem/10799)

## 문제 설명
여러 개의 쇠막대기를 레이저로 절단하려고 한다. 쇠막대기는 아래에서 위로 겹쳐 놓고, 레이저는 위에서 수직으로 발사하여 쇠막대기들을 자른다.

쇠막대기와 레이저의 배치는 다음과 같은 규칙을 따른다.
- 레이저는 `()`로 표현된다.
- 쇠막대기의 시작은 `(`, 끝은 `)`로 표현된다.
- 레이저가 지나가는 곳에서 쇠막대기가 조각난다.

입력으로 주어진 괄호 표현에서 잘려진 쇠막대기 조각의 총 개수를 구하는 프로그램을 작성한다.

## 입력
- 괄호로 이루어진 문자열이 주어진다. (최대 길이: 100,000)

## 출력
- 잘려진 쇠막대기 조각의 총 개수를 출력한다.

## 예제 입력/출력
### 예제 1
**입력**
```
()(((()())(())()))(())
```
**출력**
```
17
```

### 예제 2
**입력**
```
(((()(()()))(())()))(()())
```
**출력**
```
24
```

---

## 해결 방법
### 1. 변수 활용 방식
- 여는 괄호 `(`를 만나면 **현재 쇠막대기 개수 증가**
- 닫는 괄호 `)`를 만났을 때:
  - **이전 문자가 `(`인 경우 (레이저)** → 쇠막대기 개수 감소 후 **현재 쇠막대기 개수만큼 조각 증가**
  - **이전 문자가 `)`인 경우 (쇠막대기 끝)** → 쇠막대기 개수 감소 후 **조각 하나 추가**

#### 코드 (변수 활용 방식)
```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func CutPipe(input string) int {
	res := 0
	var pipe int
	var prev byte

	for i := 0; i < len(input); i++ {
		if input[i] == '(' {
			pipe++
			prev = '('
		} else {
			if prev == '(' {
				pipe--
				res += pipe
			} else {
				pipe--
				res++
			}
			prev = ')'
		}
	}

	return res
}

func main() {
	defer writer.Flush()

	input := ""
	fmt.Fscan(reader, &input)

	fmt.Fprintln(writer, CutPipe(input))
}
```

---

### 2. 스택을 활용한 방식
- 여는 괄호 `(`를 만나면 스택에 push
- 닫는 괄호 `)`를 만났을 때:
  - **이전 문자가 `(`인 경우 (레이저)** → pop 후 **현재 스택 크기만큼 조각 증가**
  - **이전 문자가 `)`인 경우 (쇠막대기 끝)** → pop 후 **조각 하나 추가**

#### 코드 (스택 활용 방식)
```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func CutPipeUsingStack(input string) int {
	stack := []rune{} // 스택 역할
	res := 0

	for i, ch := range input {
		if ch == '(' {
			stack = append(stack, ch) // 여는 괄호는 스택에 push
		} else { // 닫는 괄호일 때
			stack = stack[:len(stack)-1] // pop

			if input[i-1] == '(' { // 레이저인 경우
				res += len(stack) // 현재 스택 크기만큼 조각 증가
			} else { // 쇠막대기의 끝인 경우
				res++
			}
		}
	}

	return res
}

func main() {
	defer writer.Flush()

	input := ""
	fmt.Fscan(reader, &input)

	fmt.Fprintln(writer, CutPipeUsingStack(input))
}
```

---

## 방법 비교
| 방법 | 시간 복잡도 | 공간 복잡도 | 특징 |
|------|------------|------------|--------|
| 변수 활용 방식 | O(N) | O(1) | `pipe` 변수를 이용해 열린 막대 개수를 추적 |
| 스택 활용 방식 | O(N) | O(N) | 실제 스택을 사용하여 구현 |

### **결론**
- **변수 활용 방식**은 **공간 복잡도가 낮아** 더 효율적이다.
- **스택 활용 방식**은 **직관적이며 괄호 구조를 그대로 반영**하여 이해하기 쉽다.
- 두 방법 모두 시간 복잡도는 O(N)으로 동일하다.

상황에 따라 두 방법 중 적절한 방식을 선택하면 된다. 🚀

