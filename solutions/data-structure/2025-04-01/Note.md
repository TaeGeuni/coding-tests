# 최소 힙 문제 풀이
👉🏻[문제 링크](https://www.acmicpc.net/problem/1927)

## 문제 설명
널리 잘 알려진 자료구조 중 최소 힙이 있다. 최소 힙을 이용하여 다음과 같은 연산을 지원하는 프로그램을 작성하시오.

1. 배열에 자연수 x를 넣는다.
2. 배열에서 가장 작은 값을 출력하고, 그 값을 배열에서 제거한다.

프로그램은 처음에 비어있는 배열에서 시작하게 된다.

## 입력
- 첫째 줄에 연산의 개수 N(1 ≤ N ≤ 100,000)이 주어진다.
- 다음 N개의 줄에는 연산에 대한 정보를 나타내는 정수 x가 주어진다.
- 만약 x가 자연수라면 배열에 x라는 값을 넣는(추가하는) 연산이고,
- x가 0이라면 배열에서 가장 작은 값을 출력하고 그 값을 배열에서 제거하는 경우이다.
- x는 $2^{31}$보다 작은 자연수 또는 0이고, 음의 정수는 입력으로 주어지지 않는다.

## 출력
- 입력에서 0이 주어진 횟수만큼 답을 출력한다.
- 만약 배열이 비어 있는 경우인데 가장 작은 값을 출력하라고 한 경우에는 0을 출력하면 된다.

## 예제 입력
```
9
0
12345678
1
2
0
0
0
0
32
```

## 예제 출력
```
0
1
2
12345678
0
```

## 코드 구현 (Go)
```go
package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	defer writer.Flush()

	var n int
	h := new(IntHeap)
	heap.Init(h)
	fmt.Fscan(reader, &n)

	for i := 0; i < n; i++ {
		input := 0
		fmt.Fscan(reader, &input)
		if input == 0 {
			if h.Len() == 0 {
				fmt.Fprintln(writer, 0)
			} else {
				fmt.Fprintln(writer, heap.Pop(h))
			}
		} else {
			heap.Push(h, input)
		}
	}
}
```

## 코드 설명
1. `IntHeap` 타입을 정의하여 `container/heap` 패키지에서 사용 가능한 최소 힙을 구현.
2. `Push`, `Pop`, `Len`, `Less`, `Swap` 메서드를 구현하여 힙의 동작을 정의.
3. `main` 함수에서 입력을 받아서 `heap.Init(h)`으로 힙을 초기화.
4. 입력값이 `0`이면 힙에서 최소값을 출력하고 제거하며, 힙이 비어 있으면 `0`을 출력.
5. 입력값이 자연수이면 힙에 추가.
6. `writer.Flush()`를 사용하여 출력 성능을 최적화.

## 시간 복잡도
- `heap.Push` 및 `heap.Pop` 연산은 O(log N) 시간이 걸린다.
- 따라서 최대 100,000개의 연산이 주어져도 충분히 빠르게 동작한다.
