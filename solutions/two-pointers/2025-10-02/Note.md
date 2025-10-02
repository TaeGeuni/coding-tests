# 부분합 문제 풀이
👉🏻[문제 링크](https://www.acmicpc.net/problem/1806)

## 문제 설명

10,000 이하의 자연수로 이루어진 길이 N짜리 수열이 주어진다.  
이 수열에서 연속된 수들의 부분합 중에 그 합이 S 이상이 되는 것 중, 가장 짧은 것의 길이를 구하는 프로그램을 작성하시오.

- **입력**
  - 첫째 줄에 N (10 ≤ N < 100,000)과 S (0 < S ≤ 100,000,000)가 주어진다.
  - 둘째 줄에는 수열이 주어진다. 수열의 각 원소는 공백으로 구분되어 있으며, 10,000 이하의 자연수이다.

- **출력**
  - 구하고자 하는 최소의 길이를 출력한다.
  - 그러한 합을 만드는 것이 불가능하다면 0을 출력한다.

---

## 예제 입력 1
```
10 15
5 1 3 5 10 7 4 9 2 8
```

## 예제 출력 1
```
2
```

---

## 해결 방법

이 문제는 **투 포인터(two-pointer) 알고리즘**을 활용하여 해결할 수 있다.

1. `up` 포인터를 오른쪽으로 이동시키며 부분합을 증가시킨다.
2. 부분합이 S 이상이 되면 `down` 포인터를 이동시키며 길이를 최소화한다.
3. 이 과정을 반복하여 최소 길이를 구한다.
4. 만약 조건을 만족하는 부분합이 없다면 `0`을 출력한다.

이 방식은 모든 원소를 한 번씩만 확인하므로 **O(N)** 시간 복잡도로 해결 가능하다.

---

## 코드 구현 (Golang)

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

func twoPointer(n, s int, sequence []int) int {
    res := 100_001
    down := 0
    sum := 0
    for up := 0; up < n; up++ {
        sum += sequence[up]

        for sum >= s {
            length := up - down + 1
            if res > length {
                res = length
            }
            sum -= sequence[down]
            down++
        }
    }

    if res == 100_001 {
        return 0
    }

    return res
}

func main() {
    defer writer.Flush()
    var n, s int

    fmt.Fscan(reader, &n, &s)

    sequence := make([]int, n)

    for i := range sequence {
        fmt.Fscan(reader, &sequence[i])
    }

    fmt.Fprintln(writer, twoPointer(n, s, sequence))
}
```

---

## 풀이 요약

- **알고리즘**: 투 포인터
- **시간 복잡도**: O(N)
- **핵심 아이디어**: 부분합이 S 이상일 때, 왼쪽 포인터를 줄여가며 최소 길이를 탐색한다.

