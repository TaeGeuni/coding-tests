# 밑 줄
👉🏻[문제 링크](https://www.acmicpc.net/problem/1474)

## 📝 문제 설명

세준이는 **N개의 영어 단어**를 이어 붙여 **길이가 M인 단어**를 만들려고 한다. 이때 단어들 사이에는 `_`(언더바)만 넣을 수 있다. 조건은 다음과 같다:

- 단어 사이에만 `_`를 추가할 수 있다.
- 단어와 단어 사이의 `_`의 개수는 모두 같아야 한다.
- 모두 같게 만드는 것이 불가능한 경우 `_` 개수의 **최댓값과 최솟값의 차이가 1**이 되어야 한다.
- 가능한 단어 중 **사전 순으로 가장 앞서는 단어**를 출력해야 한다.

## 📥 입력

- 첫 줄: 정수 `N`, `M` (2 ≤ N ≤ 10, 3 ≤ M ≤ 200)
- 둘째 줄부터 `N`줄: 영어 단어 (길이 1~10의 알파벳 대소문자)

## 📤 출력

- 길이가 정확히 M인 새로운 단어 중, 사전 순으로 가장 앞서는 단어 출력

## 🔤 알파벳 정렬 순서

```
'A' < 'B' < ... < 'Z' < '_' < 'a' < 'b' < ... < 'z'
```

## 🧠 접근 방식

1. 모든 단어를 붙이고 남은 길이 `r` 만큼 `_`을 나눠야 한다.
2. 기본적으로 `(r / (N - 1))` 개의 `_`를 각 사이에 배치.
3. 나머지 언더바(`remain`)는 사전 순서를 고려해 조절:
   - 남은 `_`를 가능한 사전 순서가 앞서도록 분배 (앞 단어가 `_`보다 크면 그 사이에 하나 더 추가).

## ✅ 코드 (Go)

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

func MakeNewString(n, r int, sli []string) string {
	res, underline := sli[0], ""
	for i := 0; i < (r / (n - 1)); i++ {
		underline += "_"
	}
	remain := r - (n-1)*(r/(n-1))

	for i := 1; i < n; i++ {
		res += underline

		if remain != 0 {
			if n-i == remain || sli[i][0] > '_' {
				res += "_"
				remain--
			}
		}

		res += sli[i]
	}

	return res
}

func main() {
	defer writer.Flush()

	var n, m int
	fmt.Fscan(reader, &n, &m)
	r := m

	sli := make([]string, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &sli[i])
		r -= len(sli[i])
	}
	fmt.Fprintln(writer, MakeNewString(n, r, sli))
}
```

## 🔍 예제

### 입력 1
```
9 50
A
quick
brown
fox
jumps
over
the
lazy
dog
```

### 출력 1
```
A___quick__brown__fox__jumps__over__the__lazy__dog
```

### 입력 2
```
5 32
Alpha
Beta
Gamma
Delta
Epsilon
```

### 출력 2
```
Alpha_Beta_Gamma__Delta__Epsilon
```

### 입력 3
```
4 29
Hello
world
John
said
```

### 출력 3
```
Hello____world___John____said
```

## ⏱️ 시간 복잡도

- 정렬 없음, 단순 순회 및 문자열 조작 → **O(N)**