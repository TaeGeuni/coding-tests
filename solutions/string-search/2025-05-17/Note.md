# IOIOI 패턴 찾기 문제 풀이
👉🏻[문제 링크](https://www.acmicpc.net/problem/5525)

## 🧩 문제 설명

> N+1개의 I와 N개의 O로 이루어진 문자열을 PN이라고 한다.
>
> 예시:
>
> - P1: IOI  
> - P2: IOIOI  
> - P3: IOIOIOI  
> - PN: IOIOI...OI (O가 N개)
>
> I와 O로만 이루어진 문자열 S와 정수 N이 주어졌을 때, 문자열 S 안에 PN 패턴이 몇 번 등장하는지 계산하는 프로그램을 작성하라.

### 입력

- 첫째 줄: 정수 N (1 ≤ N ≤ 1,000,000)
- 둘째 줄: 문자열 S의 길이 M (2N+1 ≤ M ≤ 1,000,000)
- 셋째 줄: 문자열 S (길이 M, I와 O로만 구성)

### 출력

- PN이 S에 몇 번 포함되어 있는지 출력한다.

---

## ✅ 예제

### 입력

```
1
13
OOIOIOIOIIOII
```

### 출력

```
4
```

---

## 🔍 접근 방식

- PN의 기본 구조는 `IOIOIO...I`로 길이는 `2N + 1`이다.
- 문자열을 앞에서부터 순회하면서 `IOI` 패턴을 연속적으로 몇 번 반복하는지를 계산한다.
- 연속적인 `IOI` 패턴이 N번 이상이면 PN을 만족한다.
- 연속된 IOI 개수로부터 PN의 개수를 카운트한다.

---

## 💡 핵심 아이디어

- 슬라이딩 윈도우 또는 투 포인터 방식도 가능하지만,
- 본 코드는 상태 플래그(`ok`, `next`)와 카운트를 사용하여 연속된 IOI 패턴 수를 효율적으로 계산함.

---

## 📦 Go 코드

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

func IOIOI(n, m int, text string) int {
	res := 0
	count := 0
	ok := false
	next := false

	ioi := (2 * n) + 1

	for i := 0; i < m; i++ {
		if ok && next {
			if text[i] == 'O' {
				next = false
			} else {
				ok = false
				next = false
				if count == ioi {
					res++
				} else if count > ioi {
					count -= ioi
					res++
					res += count / 2
				}
				count = 0
			}
		} else if ok && !next {
			if text[i] == 'I' {
				next = true
				count += 2
			} else {
				ok = false
				next = false
				if count == ioi {
					res++
				} else if count > ioi {
					count -= ioi
					res++
					res += count / 2
				}
				count = 0
			}
		}
		if text[i] == 'I' && !ok {
			ok = true
			next = true
			count++
		}
		if i == m-1 {
			if count == ioi {
				res++
			} else if count > ioi {
				count -= ioi
				res++
				res += count / 2
			}
		}
	}

	return res
}

func main() {
	defer writer.Flush()

	var n, m int
	fmt.Fscan(reader, &n, &m)
	text := ""
	fmt.Fscan(reader, &text)

	res := IOIOI(n, m, text)
	fmt.Fprintln(writer, res)
}
```

---

## ✅ 결과

이 코드는 O(M) 시간복잡도로, 주어진 입력 크기 제한 내에서 효율적으로 동작한다.
