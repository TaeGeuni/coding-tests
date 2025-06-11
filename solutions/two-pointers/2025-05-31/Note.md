# 서로 다른 두 정수의 합이 K에 가장 가까운 조합의 수
👉🏻[문제 링크](https://www.acmicpc.net/problem/9024)

## 문제 설명

여러 개의 서로 다른 정수 S = \{a₁, a₂, …, aₙ\}와 또 다른 정수 K가 주어졌을 때, S에 속하는 서로 다른 두 개의 정수의 합이 K에 가장 가까운 조합의 수를 구하는 문제입니다.

예를 들어,  
S = \{-7, 9, 2, -4, 12, 1, 5, -3, -2, 0\}, K = 8  
→ 가장 가까운 조합은 {12, -4}이고 조합의 수는 1입니다.  
K = 4  
→ 가장 가까운 조합은 {-7, 12}, {9, -4}, {5, -2}, {5, 0}, {1, 2}이며, 총 5개입니다.

---

## 입력

- 첫째 줄에 테스트 케이스의 개수 t가 주어진다.
- 각 테스트 케이스는 두 줄로 구성된다.
  - 첫 줄: 정수의 개수 n (2 ≤ n ≤ 1,000,000), 정수 K (-10⁸ ≤ K ≤ 10⁸)
  - 둘째 줄: n개의 정수 (-10⁸ ≤ aᵢ ≤ 10⁸)

---

## 출력

- 각 테스트 케이스마다 K에 가장 가까운 두 정수의 조합 수를 출력한다.

---

## 예제 입력

```
4
10 8
-7 9 2 -4 12 1 5 -3 -2 0
10 4
-7 9 2 -4 12 1 5 -3 -2 0
4 20
1 7 3 5
5 10
3 9 7 1 5
```

## 예제 출력

```
1
5
1
2
```

---

## 풀이 설명

1. 먼저 입력 받은 수열을 오름차순 정렬한다.
2. 투 포인터 방법을 사용하여 양 끝에서부터 접근한다.
3. 현재 두 수의 합과 K의 차이의 절댓값을 기준으로 가장 작은 값을 추적한다.
4. 그 차이에 해당하는 경우의 수를 카운팅한다.
5. 포인터를 이동시키면서 최솟값이 갱신되는지 비교하며 탐색을 진행한다.

---

## Go 코드 구현

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

func SumOfTheClosestNumbers(n, k int, nums []int) int {
    sort.Ints(nums)
    left, right := 0, n-1
    check := make(map[int]int)
    mostSimilarNumber := k - (nums[right] + nums[left])
    if mostSimilarNumber < 0 {
        mostSimilarNumber = -mostSimilarNumber
    }

    for left < right {
        num := nums[right] + nums[left]
        check[k-num]++
        diff := k - num
        if diff < 0 {
            diff = -diff
        }
        if diff < mostSimilarNumber {
            mostSimilarNumber = diff
        }

        a, b := nums[right]+nums[left+1], nums[right-1]+nums[left]
        numA, numB := k-a, k-b
        if numA < 0 {
            numA = -numA
        }
        if numB < 0 {
            numB = -numB
        }

        if numA > numB {
            right--
        } else {
            left++
        }
    }

    if mostSimilarNumber == 0 {
        return check[mostSimilarNumber]
    }
    return check[-mostSimilarNumber] + check[mostSimilarNumber]
}

func main() {
    defer writer.Flush()

    var t int
    fmt.Fscan(reader, &t)

    res := make([]int, t)

    for i := 0; i < t; i++ {
        var n, k int
        fmt.Fscan(reader, &n, &k)

        nums := make([]int, n)
        for j := 0; j < n; j++ {
            fmt.Fscan(reader, &nums[j])
        }
        res[i] = SumOfTheClosestNumbers(n, k, nums)
    }

    for i := 0; i < t; i++ {
        fmt.Fprintln(writer, res[i])
    }
}
```
