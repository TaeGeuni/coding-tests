# A와 B
👉🏻[문제 링크](https://www.acmicpc.net/problem/12904)

## 🧩 문제 설명

수빈이는 A와 B로만 이루어진 영어 단어가 존재한다는 사실에 놀랐다. 대표적인 예로 AB, BAA, AA, ABBA 등이 있다.

이런 사실에 놀란 수빈이는 간단한 게임을 만들기로 했다. 두 문자열 S와 T가 주어졌을 때, S를 T로 바꾸는 게임이다. 문자열을 바꿀 때는 다음과 같은 두 가지 연산만 가능하다.

1. 문자열의 뒤에 A를 추가한다.
2. 문자열을 뒤집고 뒤에 B를 추가한다.

주어진 조건을 이용해서 S를 T로 만들 수 있는지 없는지 알아내는 프로그램을 작성하시오.

---

## 📥 입력

- 첫째 줄에 문자열 S
- 둘째 줄에 문자열 T

* (1 ≤ S의 길이 ≤ 999, 2 ≤ T의 길이 ≤ 1000, S의 길이 < T의 길이)

---

## 📤 출력

- S를 T로 바꿀 수 있으면 `1`을, 없으면 `0`을 출력한다.

---

## 🧪 예제 입력 1
```
B
ABBA
```

## ✅ 예제 출력 1
```
1
```

## 🧪 예제 입력 2
```
AB
ABB
```

## ✅ 예제 출력 2
```
0
```

---

## 💡 풀이 설명

이 문제는 문자열 S에서 시작해서 T를 만드는 것이 아니라, **T에서 시작해서 S로 도달할 수 있는지 역추적**하는 방식으로 해결하는 것이 핵심이다. 왜냐하면 T에서 가능한 연산이 역방향으로는 유일하게 결정되기 때문이다.

### 🔁 가능한 역연산

- T의 마지막 문자가 'A'라면 마지막 A를 제거한다.
- T의 마지막 문자가 'B'라면 마지막 B를 제거하고 문자열을 뒤집는다.

### 🔍 알고리즘

1. 문자열 T의 길이가 S의 길이와 같아질 때까지 위의 연산을 반복한다.
2. 최종적으로 S와 T가 같다면 1을 반환, 그렇지 않으면 0을 반환한다.

---

## 🧠 Go 코드

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

func AB(s, t string) int {
	for len(s) != len(t) {
		if t[len(t)-1] == 'A' {
			t = t[:len(t)-1]
		} else {
			t = t[:len(t)-1]
			sli := make([]byte, len(t))
			for i := 0; i < len(t); i++ {
				sli[i] = t[len(t)-1-i]
			}
			t = string(sli)
		}
	}

	if s == t {
		return 1
	}
	return 0
}

func main() {
	defer writer.Flush()

	var s, t string
	fmt.Fscan(reader, &s)
	fmt.Fscan(reader, &t)

	fmt.Fprintln(writer, AB(s, t))
}
```

---

## 🏁 시간 복잡도

- O(N²): 매번 뒤집기 연산을 할 때 O(N) 시간이 걸리며, 최대 N번 수행된다.
- 하지만 N ≤ 1000이므로 충분히 통과 가능하다.

---

## 📌 핵심 포인트

- 정방향이 아닌 **역방향으로 문제를 해결**해야 풀 수 있다.
- 문자열의 끝을 보고 판단하는 **그리디한 접근**이 유효하다.
