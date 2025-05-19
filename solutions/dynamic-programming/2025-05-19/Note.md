# 전깃줄 - 교차하지 않게 하기 (LIS 문제)
👉🏻[문제 링크](https://www.acmicpc.net/problem/2565)

## 문제 설명

두 전봇대 A와 B 사이에 여러 개의 전깃줄이 설치되어 있다. 전깃줄들이 서로 교차하면 합선의 위험이 있기 때문에, 몇 개의 전깃줄을 제거하여 **서로 교차하지 않도록** 만들고자 한다.

### 입력
- 첫 줄에 전깃줄의 개수 `N` (1 ≤ N ≤ 100)
- 다음 N개의 줄에는 각 줄마다 A 전봇대의 위치 `a`와 B 전봇대의 위치 `b`가 주어진다.
- A, B 전봇대의 위치 번호는 500 이하의 자연수이며, 같은 위치에 두 전깃줄이 연결되지 않는다.

### 출력
- 남아있는 전깃줄이 **서로 교차하지 않게 하기 위해 없애야 하는 전깃줄의 최소 개수**를 출력한다.

---

## 예제 입력

```
8
1 8
3 9
2 2
4 1
6 4
10 10
9 7
7 6
```

## 예제 출력

```
3
```

---

## 문제 풀이

이 문제는 **최장 증가 부분 수열 (LIS)** 을 응용해서 해결할 수 있다.

1. 입력 받은 전깃줄 정보를 A 전봇대 기준으로 정렬한다.
2. A 전봇대의 위치를 기준으로 정렬한 뒤, B 전봇대의 위치를 가지고 LIS를 구한다.
3. 전체 전깃줄 수 `N`에서 LIS의 길이를 빼면, 제거해야 하는 전깃줄의 수가 된다.

### 핵심 아이디어

- A 전봇대를 기준으로 정렬하면, B 전봇대의 위치들만 고려해서 교차 여부를 판단할 수 있다.
- 교차하지 않게 전깃줄을 최대한 남기는 것은 **B 전봇대의 연결 위치들이 증가하는 수열을 만드는 것**과 동일하다.

---

## Go 코드

```go
package main

import (
    "bufio"
    "fmt"
    "os"
    "sort"
)

var (
    reader = bufio.NewReader(os.Stdin)
    writer = bufio.NewWriter(os.Stdout)
)

func LIS(n int, pole [][]int) int {
    tails := make([]int, 0)
    tails = append(tails, 0)

    for i := 0; i < n; i++ {
        if tails[len(tails)-1] < pole[i][1] {
            tails = append(tails, pole[i][1])
        } else {
            left, right := 0, len(tails)-1
            res := 0
            for left <= right {
                mid := (left + right) / 2
                num := tails[mid]
                if num >= pole[i][1] {
                    res = mid
                    right = mid - 1
                } else {
                    left = mid + 1
                }
            }
            tails[res] = pole[i][1]
        }
    }

    return (n - (len(tails) - 1))
}

func main() {
    defer writer.Flush()

    var n int
    fmt.Fscan(reader, &n)

    pole := make([][]int, n)

    for i := 0; i < n; i++ {
        a, b := 0, 0
        fmt.Fscan(reader, &a, &b)
        pole[i] = append(pole[i], a)
        pole[i] = append(pole[i], b)
    }

    sort.Slice(pole, func(i, j int) bool {
        return pole[i][0] < pole[j][0]
    })

    fmt.Fprintln(writer, LIS(n, pole))
}
```

---

## 시간복잡도

- LIS 탐색: O(N log N)
- 입력 정렬: O(N log N)
- 전체 시간복잡도는 O(N log N)이며, N ≤ 100이므로 매우 효율적으로 동작한다.

---

## 결론

- 이 문제는 전형적인 LIS(최장 증가 부분 수열) 응용 문제다.
- 전깃줄의 시작 지점을 정렬하고, 끝 지점을 기준으로 LIS를 구하는 패턴을 익혀두자.
