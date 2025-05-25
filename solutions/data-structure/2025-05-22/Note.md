# 괄호의 값
👉🏻[문제 링크](https://www.acmicpc.net/problem/2504)

## 문제 설명

4개의 기호 ‘(’, ‘)’, ‘[’, ‘]’를 이용해서 만들어지는 괄호열 중에서 올바른 괄호열이란 다음과 같이 정의된다.

1. 한 쌍의 괄호로만 이루어진 ‘()’와 ‘[]’는 올바른 괄호열이다.  
2. 만일 X가 올바른 괄호열이면 ‘(X)’이나 ‘[X]’도 모두 올바른 괄호열이 된다.  
3. X와 Y 모두 올바른 괄호열이라면 이들을 결합한 XY도 올바른 괄호열이 된다.

예: `(()[[]])`, `(())[][]` 는 올바른 괄호열  
반례: `([)]`, `(()()[]` 는 올바르지 않음

올바른 괄호열 X에 대하여 그 괄호열의 값(괄호값)은 다음과 같이 정의된다.

- `()`의 값은 2  
- `[]`의 값은 3  
- `(X)`의 값은 2 × 값(X)  
- `[X]`의 값은 3 × 값(X)  
- XY는 값(X) + 값(Y)

### 입력

첫째 줄에 괄호열을 나타내는 문자열 S가 주어진다.  
(1 ≤ S의 길이 ≤ 30)

### 출력

괄호열이 올바르면 그 괄호열의 값을 출력한다.  
올바르지 않으면 0을 출력한다.

---

## 예제

### 입력 1
```
(()[[]])([])
```
### 출력 1
```
28
```

### 입력 2
```
[][]((])
```
### 출력 2
```
0
```

---

## 풀이 (Go)

이 문제는 **스택**을 이용하여 괄호열을 탐색하고, 조건에 따라 괄호 값을 계산하는 문제입니다. 괄호가 열릴 때마다 스택에 기대하는 닫는 괄호를 push 하고, 닫는 괄호가 등장할 때마다 스택을 검사합니다.

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
                num2 *= 2
            }
            for i := 0; i < check['[']; i++ {
                num3 *= 3
            }
            res += num2 * num3
            check['(']--
            pos = false
            stack = stack[:len(stack)-1]
        } else if s[i] == ']' && pos {
            num2, num3 := 1, 1
            for i := 0; i < check['(']; i++ {
                num2 *= 2
            }
            for i := 0; i < check['[']; i++ {
                num3 *= 3
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
```

## 시간 복잡도

- 입력 문자열 길이 n에 대해 O(n)
- 최대 길이 30이므로 매우 효율적

## 핵심 아이디어

- 스택으로 괄호 열림/닫힘을 관리
- 열린 괄호가 닫힐 때 값을 계산해서 누적
- 중첩될 경우 각 괄호마다 값을 누적해 곱한 뒤 합산
- 올바르지 않은 괄호열을 조기에 감지하여 0 반환

---