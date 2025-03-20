# 꽃길
👉🏻[문제 링크](https://www.acmicpc.net/problem/14620)

## 문제 설명
N x N 크기의 정원이 있으며, 각 칸에는 꽃을 심을 수 있는 비용이 존재한다. 꽃을 심을 때, 꽃잎이 퍼지는 모양은 십자형(상하좌우 한 칸씩)이다. 정원에서 세 개의 꽃을 심을 때, 꽃잎이 서로 겹치지 않도록 하면서 최소 비용이 되도록 하는 문제이다.

## 해결 방법
1. **모든 가능한 위치를 탐색**
   - 정원의 가장자리를 제외한 모든 칸에서 꽃을 심을 수 있다.
   - 세 개의 꽃을 선택하는 모든 조합을 확인한다.
2. **꽃을 심는 비용 계산**
   - 각 꽃은 해당 위치의 비용과 십자형으로 퍼지는 네 개의 칸의 비용을 포함한다.
3. **꽃잎이 겹치지 않는지 확인**
   - 꽃을 심었을 때, 이미 다른 꽃잎이 존재하는지 검사한다.
4. **최소 비용 갱신**
   - 가능한 모든 경우 중 가장 적은 비용을 선택한다.

## 코드 구현 (Golang)
```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func CheapestCase(n int, garden [][]int) int {
	res := 0
	var x1, y1, x2, y2, x3, y3 int

	for i := 0; i < n*n; i++ {
		check1 := make(map[[2]int]struct{})
		cost1 := 0
		if x1 == n {
			x1 = 0
			y1++
		}
		if x1 == 0 || y1 == 0 || x1 == n-1 || y1 == n-1 {
			x1++
			continue
		}
		cost1 += (garden[y1][x1] + garden[y1-1][x1] + garden[y1][x1-1] + garden[y1][x1+1] + garden[y1+1][x1])
		check1[[2]int{y1, x1}] = struct{}{}
		check1[[2]int{y1 - 1, x1}] = struct{}{}
		check1[[2]int{y1, x1 - 1}] = struct{}{}
		check1[[2]int{y1, x1 + 1}] = struct{}{}
		check1[[2]int{y1 + 1, x1}] = struct{}{}
		x2, y2 = x1, y1
		for j := i; j < n*n; j++ {
			check2 := make(map[[2]int]struct{})
			cost2 := 0

			if x2 == n {
				x2 = 0
				y2++
			}
			if x2 == 0 || y2 == 0 || x2 == n-1 || y2 == n-1 {
				x2++
				continue
			}
			if _, ok := check1[[2]int{y2, x2}]; ok {
				x2++
				continue
			}
			if _, ok := check1[[2]int{y2 - 1, x2}]; ok {
				x2++
				continue
			}
			if _, ok := check1[[2]int{y2, x2 - 1}]; ok {
				x2++
				continue
			}
			if _, ok := check1[[2]int{y2, x2 + 1}]; ok {
				x2++
				continue
			}
			if _, ok := check1[[2]int{y2 + 1, x2}]; ok {
				x2++
				continue
			}
			cost2 += (garden[y2][x2] + garden[y2-1][x2] + garden[y2][x2-1] + garden[y2][x2+1] + garden[y2+1][x2])
			check2[[2]int{y2, x2}] = struct{}{}
			check2[[2]int{y2 - 1, x2}] = struct{}{}
			check2[[2]int{y2, x2 - 1}] = struct{}{}
			check2[[2]int{y2, x2 + 1}] = struct{}{}
			check2[[2]int{y2 + 1, x2}] = struct{}{}
			x3, y3 = x2, y2

			for k := j; k < n*n; k++ {
				cost3 := 0
				if x3 == n {
					x3 = 0
					y3++
				}
				if x3 == 0 || y3 == 0 || x3 == n-1 || y3 == n-1 {
					x3++
					continue
				}
				if _, ok := check1[[2]int{y3, x3}]; ok {
					x3++
					continue
				}
				if _, ok := check1[[2]int{y3 - 1, x3}]; ok {
					x3++
					continue
				}
				if _, ok := check1[[2]int{y3, x3 - 1}]; ok {
					x3++
					continue
				}
				if _, ok := check1[[2]int{y3, x3 + 1}]; ok {
					x3++
					continue
				}
				if _, ok := check1[[2]int{y3 + 1, x3}]; ok {
					x3++
					continue
				}
				if _, ok := check2[[2]int{y3, x3}]; ok {
					x3++
					continue
				}
				if _, ok := check2[[2]int{y3 - 1, x3}]; ok {
					x3++
					continue
				}
				if _, ok := check2[[2]int{y3, x3 - 1}]; ok {
					x3++
					continue
				}
				if _, ok := check2[[2]int{y3, x3 + 1}]; ok {
					x3++
					continue
				}
				if _, ok := check2[[2]int{y3 + 1, x3}]; ok {
					x3++
					continue
				}
				cost3 += (garden[y3][x3] + garden[y3-1][x3] + garden[y3][x3-1] + garden[y3][x3+1] + garden[y3+1][x3])
				if res == 0 {
					res = cost1 + cost2 + cost3
				} else {
					if res > (cost1 + cost2 + cost3) {
						res = cost1 + cost2 + cost3
					}
				}
				x3++
			}
			x2++
		}
		x1++
	}

	return res
}

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	garden := make([][]int, n)
	for i := 0; i < n; i++ {
		garden[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Fscan(reader, &garden[i][j])
		}
	}

	fmt.Fprintln(writer, CheapestCase(n, garden))

}
```

## 시간 복잡도
- O(N^6): 세 개의 꽃을 심을 수 있는 모든 조합을 확인해야 하므로 탐색의 최악의 경우 O(N^6)까지 증가할 수 있다. 하지만 N이 작아(6 <= N <= 10) 가능하다.

## 개선점
- 현재 풀이는 완전 탐색 기반이므로, 백트래킹을 활용하면 탐색 범위를 줄일 수 있다.
