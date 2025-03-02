# 시리얼 번호 정렬 문제 해설
👉🏻[문제 링크](https://www.acmicpc.net/problem/1431)


## 문제 설명
다솜이는 기타의 시리얼 번호를 정렬하려고 합니다. 시리얼 번호는 알파벳 대문자(A-Z)와 숫자(0-9)로 이루어져 있으며, 정렬 기준은 다음과 같습니다:

1. **길이가 짧은 순서**로 정렬합니다.
2. **길이가 같다면, 숫자의 합이 작은 순서**로 정렬합니다.
   - 숫자인 문자들만 더합니다.
3. **위 두 조건으로도 정렬되지 않으면, 사전순으로 비교**하여 정렬합니다.
   - 숫자가 알파벳보다 사전순으로 작습니다.

## 입력 형식
- 첫 번째 줄에 기타의 개수 `N`이 주어집니다. (`N <= 50`)
- 이후 `N`개의 줄에 시리얼 번호가 주어집니다.
- 시리얼 번호는 중복되지 않으며 길이는 최대 `50`입니다.

## 출력 형식
- 정렬된 시리얼 번호를 한 줄에 하나씩 출력합니다.

---

## 코드 설명

### 1. 입력 처리
```go
var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

defer writer.Flush()

var guitarNum int
fmt.Fscan(reader, &guitarNum)

serialNumbers := make([]string, guitarNum)

for i := 0; i < guitarNum; i++ {
    fmt.Fscan(reader, &serialNumbers[i])
}
```
- `bufio.Reader`와 `bufio.Writer`를 사용하여 입출력 성능을 최적화합니다.
- 기타 개수 `guitarNum`을 입력받고, `serialNumbers` 배열을 생성하여 시리얼 번호들을 저장합니다.

### 2. 정렬 조건 적용
```go
sort.Slice(serialNumbers, func(i, j int) bool {
    if len(serialNumbers[i]) != len(serialNumbers[j]) {
        return len(serialNumbers[i]) < len(serialNumbers[j])
    }

    a := 0
    b := 0

    for m := 0; m < len(serialNumbers[i]); m++ {
        data := string(serialNumbers[i][m])
        num, _ := strconv.Atoi(data)
        a += num
    }
    for m := 0; m < len(serialNumbers[j]); m++ {
        data := string(serialNumbers[j][m])
        num, _ := strconv.Atoi(data)
        b += num
    }

    if a != b {
        return a < b
    }

    return serialNumbers[i] < serialNumbers[j]
})
```
#### 정렬 기준
1. **길이 비교**: `len(serialNumbers[i])`와 `len(serialNumbers[j])`를 비교하여 짧은 것이 먼저 오도록 정렬합니다.
2. **숫자의 합 비교**:
   - 각 문자열을 순회하며 숫자만 더하는 방식으로 합을 구합니다.
   - `strconv.Atoi`를 사용하여 문자를 숫자로 변환하는데, 알파벳인 경우 변환이 실패하여 0이 더해집니다.
   - 숫자의 합이 작은 것이 먼저 오도록 정렬합니다.
3. **사전순 비교**: 위 조건으로도 정렬되지 않으면, 문자열을 그대로 비교하여 사전순으로 정렬합니다.

### 3. 정렬된 결과 출력
```go
for i := 0; i < guitarNum; i++ {
    fmt.Fprintln(writer, serialNumbers[i])
}
```
- 정렬된 `serialNumbers`를 한 줄씩 출력합니다.

---

## 예제
### 입력 예제 1
```
5
ABCD
145C
A
A910
Z321
```

### 출력 예제 1
```
A
ABCD
Z321
145C
A910
```

### 입력 예제 2
```
2
Z19
Z20
```

### 출력 예제 2
```
Z20
Z19
```

---

## 시간 복잡도 분석
- `sort.Slice`를 사용하여 **O(N log N)** 의 시간 복잡도로 정렬을 수행합니다.

---
