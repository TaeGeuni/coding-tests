# 풍선 터뜨리기 문제
👉🏻[문제 링크](https://www.acmicpc.net/problem/2346)

## 문제 설명

1번부터 N번까지 N개의 풍선이 원형으로 놓여 있고, i번 풍선의 오른쪽에는 i+1번 풍선이, 왼쪽에는 i-1번 풍선이 있습니다. 단, 1번 풍선의 왼쪽에는 N번 풍선이, N번 풍선의 오른쪽에는 1번 풍선이 있습니다. 각 풍선 안에는 -N보다 크거나 같고, N보다 작거나 같은 정수가 적힌 종이가 들어 있습니다. 

이 풍선들을 다음과 같은 규칙으로 터뜨립니다:

1. 제일 처음에는 1번 풍선을 터뜨립니다.
2. 터진 풍선 안의 종이에 적힌 값을 읽고, 해당 값만큼 이동하여 다음 풍선을 터뜨립니다. 양수는 오른쪽으로, 음수는 왼쪽으로 이동합니다.
3. 이동 시 이미 터진 풍선은 건너뜁니다.

### 입력

- 첫째 줄: 자연수 N (1 ≤ N ≤ 1,000)
- 둘째 줄: 각 풍선 안의 종이에 적힌 정수 (공백으로 구분된 N개의 정수, 0은 없음)

### 출력

- 첫째 줄: 터진 풍선의 번호를 터진 순서대로 출력

### 예제 입력 1

```text
5
3 2 1 -3 -1
```

### 예제 출력 1

```text
1 4 5 3 2
```

## 풀이 코드

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)
	balloons := make([][]int, n+1)

	for i := 1; i < n+1; i++ {
		v := 0
		fmt.Fscan(reader, &v)
		balloons[i] = append(balloons[i], v)
	}

	start := 1
	move := 0
	res := make([]int, 0)

	for len(res) < n {
		if move == 0 {
			res = append(res, start)
			move = balloons[start][0]
			balloons[start][0] = 0
		} else if move > 0 {
			if start == n {
				start = 1
			} else {
				start++
			}

			if balloons[start][0] != 0 {
				move--
			}
		} else if move < 0 {
			if start == 1 {
				start = n
			} else {
				start--
			}

			if balloons[start][0] != 0 {
				move++
			}
		}
	}

	ans := strings.Trim(fmt.Sprint(res), "[]")
	fmt.Fprintln(writer, ans)
}
```

## 풀이 설명

1. 입력된 풍선의 정보를 `balloons` 배열에 저장합니다.
2. 1번 풍선부터 시작하여, 이동해야 할 거리를 `move` 변수로 기록합니다.
3. 현재 풍선 번호를 결과 배열 `res`에 추가한 뒤, 해당 풍선의 값을 0으로 설정하여 더 이상 참조되지 않도록 합니다.
4. 이동해야 할 거리 `move`의 부호에 따라 오른쪽 또는 왼쪽으로 이동하며, 이미 터진 풍선을 건너뜁니다.
5. 모든 풍선이 터질 때까지 위 과정을 반복합니다.
6. 최종 결과를 출력 형식에 맞게 변환하여 출력합니다.

## 동작 원리

- 풍선의 순환 구조는 `start` 값을 적절히 갱신하여 구현합니다.
- 이미 터진 풍선은 값이 0인지 확인하여 건너뜁니다.
- 이동 거리를 처리할 때 이동 방향(양수/음수)에 따라 분기하여 처리합니다.

## 시간 복잡도

- 풍선이 N개일 때, 각 풍선이 최대 한 번씩 처리되므로 시간 복잡도는 **O(N)**입니다.
