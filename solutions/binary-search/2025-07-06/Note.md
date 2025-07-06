# 팀 목표레벨 (Team Target Level)
👉🏻[문제 링크](https://www.acmicpc.net/problem/16564)

## 문제 설명

성권이는 Heroes of the Storm 프로게이머 지망생이다.

이 게임에는 총 **N개의 캐릭터**가 있다. 현재 각 캐릭터의 레벨은 `X_i`이다.  
성권이는 앞으로 게임이 끝날 때까지, **레벨을 총합 K만큼 올릴 수 있다**.

**팀 목표레벨 T**는 다음과 같이 정의된다:

> T = min(Xi) (1 ≤ i ≤ N)

즉, 모든 캐릭터의 레벨 중 가장 낮은 레벨을 말한다.

성권이는 게임이 끝날 때까지 가능한 최대 팀 목표레벨 T를 달성하고 싶어 한다.

---

## 입력

- 첫째 줄: 정수 N(1 ≤ N ≤ 1,000,000), K(1 ≤ K ≤ 1,000,000,000)
- 둘째 줄부터 N개의 줄: 각 캐릭터의 현재 레벨 X₁, X₂, ..., Xₙ (1 ≤ Xi ≤ 1,000,000,000)

## 출력

- 가능한 최대 팀 목표레벨 T를 출력한다.

---

## 예제 입력 1

```
3 10
10
20
15
```

## 예제 출력 1

```
17
```

---

## 풀이 요약

이 문제는 **이분 탐색**을 이용해서 가능한 최대 팀 목표레벨 `T`를 찾는 문제입니다.

- 현재 각 캐릭터의 레벨을 정렬합니다.
- 최소 레벨을 기준으로, `T` 이상으로 올리기 위해 필요한 레벨 총합을 구합니다.
- 만약 그 합이 `K`보다 작거나 같으면 가능한 `T`입니다.
- 이를 통해 가능한 최대 `T`를 이분 탐색으로 탐색합니다.

---

## 사용한 코드 (Go)

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

func isItPossible(levels []int, k, mid int) bool {
    num := 0
    for i := 0; i < len(levels); i++ {
        if levels[i] < mid {
            num += mid - levels[i]
        }
        if num > k {
            return false
        }
    }
    return true
}

func binarySearchMaxTarget(levels []int, k int) int {
    res := 0
    lo, hi := levels[0], levels[len(levels)-1]+k

    for lo <= hi {
        mid := (lo + hi) / 2
        if isItPossible(levels, k, mid) {
            res = mid
            lo = mid + 1
        } else {
            hi = mid - 1
        }
    }

    return res
}

func main() {
    defer writer.Flush()

    var n, k int
    fmt.Fscan(reader, &n, &k)

    levels := make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Fscan(reader, &levels[i])
    }

    sort.Ints(levels)
    result := binarySearchMaxTarget(levels, k)
    fmt.Fprintln(writer, result)
}
```

---

## 시간 복잡도

- 정렬: O(N log N)
- 이분 탐색: O(log(max + K)) * O(N)
- 전체 시간 복잡도는 **O(N log N + N log(max + K))** 로 충분히 빠릅니다.

---

## 핵심 포인트

- **최솟값을 최대화**하는 문제는 이분 탐색의 전형적인 패턴입니다.
- 레벨을 올리기 위해 필요한 합을 누적해서 판단하는 방식으로 조건을 체크합니다.
