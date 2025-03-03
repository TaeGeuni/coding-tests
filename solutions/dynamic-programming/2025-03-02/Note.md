# 파스칼 삼각형을 이용한 부분 삼각형 합 구하기
👉🏻[문제 링크](https://www.acmicpc.net/problem/15489)

## 문제 설명
파스칼 삼각형에서 특정 위치 (R, C)를 꼭짓점으로 하고 한 변이 W인 정삼각형의 내부 숫자 합을 구하는 문제입니다. 파스칼 삼각형은 다음 규칙을 따릅니다:

- 삼각형의 양 끝 값은 항상 1입니다.
- 내부 값은 바로 위의 두 수의 합으로 구해집니다.

## 접근 방법
1. **파스칼 삼각형을 생성**합니다.
2. **주어진 위치에서 W 크기의 삼각형 내부 숫자들을 합산**합니다.

---

## 코드 설명
### 1. 파스칼 삼각형 생성 (`generatePascalsTriangle`)
```go
func generatePascalsTriangle(R, W int) map[[2]int]int {
    pascMap := make(map[[2]int]int)
    for i := 1; i < R+W; i++ {
        for j := 1; j <= i; j++ {
            if j == 1 || j == i {
                pascMap[[2]int{i, j}] = 1
            } else {
                pascMap[[2]int{i, j}] = pascMap[[2]int{i - 1, j - 1}] + pascMap[[2]int{i - 1, j}]
            }
        }
    }
    return pascMap
}
```
- (i, j) 좌표를 키로 하는 맵을 활용하여 파스칼 삼각형을 생성합니다.
- 위의 두 수의 합으로 내부 값을 계산합니다.

### 2. 부분 삼각형의 합 계산 (`sumTriangle`)
```go
func sumTriangle(pascMap map[[2]int]int, R, C, W int) int {
    res := 0
    count := C
    for i := R; i < R+W; i++ {
        for j := C; j <= count; j++ {
            res += pascMap[[2]int{i, j}]
        }
        count++
    }
    return res
}
```
- 꼭짓점 (R, C)부터 W 크기의 정삼각형 내 값을 합산합니다.
- 각 행에서 선택할 수의 개수를 점점 증가시킵니다.

### 3. 메인 함수 (`main`)
```go
func main() {
    defer writer.Flush()
    var R, C, W int
    fmt.Fscan(reader, &R, &C, &W)
    pascMap := generatePascalsTriangle(R, W)
    fmt.Fprintln(writer, sumTriangle(pascMap, R, C, W))
}
```
- 입력을 받고 파스칼 삼각형을 생성한 후, 결과를 출력합니다.

---

## 시간 복잡도 분석
- **파스칼 삼각형 생성:** O(N^2) (최대 29x29 연산)
- **부분 삼각형 합산:** O(W^2) (최대 29x29 연산)

## 예제 실행
### 입력
```
3 1 4
```
### 파스칼 삼각형의 해당 부분
![이미지 URL](https://onlinejudgeimages.s3-ap-northeast-1.amazonaws.com/problem/15489/1.png)

### 출력
```
42
```