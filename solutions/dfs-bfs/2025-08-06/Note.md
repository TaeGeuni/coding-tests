# 보물섬 문제 풀이
👉🏻[문제 링크](https://www.acmicpc.net/problem/2589)

## 문제 설명

보물섬 지도는 `L`(육지)과 `W`(바다)로 구성된 직사각형 격자입니다. 인접한 육지를 상하좌우로 이동할 수 있으며, 한 칸 이동에 한 시간이 소요됩니다. 보물은 최단 거리로 이동했을 때 가장 긴 시간이 걸리는 두 육지 칸 사이에 숨겨져 있습니다. 이때, 최장 최단거리(가장 먼 육지 두 곳의 최단 거리)를 구하는 문제입니다.

## 입력

- 첫 줄에 세로 크기 `v`, 가로 크기 `h`가 주어집니다. (1 ≤ v, h ≤ 50)
- `v`줄에 걸쳐 `L`(육지), `W`(바다)로 이루어진 보물 지도가 주어집니다.

## 출력

- 보물이 묻힌 두 곳 사이를 최단 거리로 이동하는 시간을 출력합니다.

## 예제 입력

```
5 7
WLLWWWL
LLLWLLL
LWLWLWW
LWLWLLL
WLLWLWW
```

## 예제 출력
```
8
```


## 풀이 방식

- 모든 육지 칸(`L`)에서 BFS를 수행하여, 해당 칸에서 도달 가능한 가장 먼 육지 칸까지의 최단 거리를 구합니다.
- 그 중에서 최댓값을 선택합니다.

## Go 코드

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

type Point struct {
    X int
    Y int
}

type PointAndCount struct {
    p     Point
    count int
}

func treasureHunt(v, h int, treasureMap []string) int {
    res := 0
    dx := []int{0, 0, -1, 1}
    dy := []int{-1, 1, 0, 0}

    for i := 0; i < v; i++ {
        for j := 0; j < h; j++ {
            if treasureMap[i][j] == 'L' {
                visited := make([][]bool, v)
                for k := range visited {
                    visited[k] = make([]bool, h)
                }
                queue := make([]PointAndCount, 0, v*h)
                queue = append(queue, PointAndCount{Point{X: j, Y: i}, 0})
                visited[i][j] = true
                head := 0
                maxDistInBfs := 0

                for head < len(queue) {
                    node := queue[head]
                    head++

                    if node.count > maxDistInBfs {
                        maxDistInBfs = node.count
                    }

                    for k := 0; k < 4; k++ {
                        nx, ny := node.p.X+dx[k], node.p.Y+dy[k]
                        if nx < 0 || nx >= h || ny < 0 || ny >= v {
                            continue
                        }
                        if !visited[ny][nx] && treasureMap[ny][nx] == 'L' {
                            visited[ny][nx] = true
                            queue = append(queue, PointAndCount{Point{X: nx, Y: ny}, node.count + 1})
                        }
                    }
                }
                if maxDistInBfs > res {
                    res = maxDistInBfs
                }
            }
        }
    }
    return res
}

func main() {
    defer writer.Flush()
    v, h := 0, 0
    fmt.Fscan(reader, &v, &h)
    treasureMap := make([]string, v)
    for i := 0; i < v; i++ {
        fmt.Fscan(reader, &treasureMap[i])
    }
    fmt.Fprintln(writer, treasureHunt(v, h, treasureMap))
}
```
---
## 시간복잡도

- 최대 2500개의 육지 칸에서 BFS를 수행하므로, 전체 복잡도는 O(N² × N²)이며 제한 범위 내에서 실행 가능합니다.