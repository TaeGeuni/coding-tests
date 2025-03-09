# 재귀함수가 뭔가요?
👉🏻[문제 링크](https://www.acmicpc.net/problem/17478)

## 문제 개요

이 문제는 **재귀 함수**의 개념을 이해하고, 이를 활용하여 특정 패턴의 출력을 생성하는 것이 목표입니다.
재귀적으로 호출되는 함수가 있으며, 각 호출마다 들여쓰기를 추가하여 계층적인 구조를 유지해야 합니다.

## 입출력 설명

### 입력
- 하나의 정수 `N`이 주어지며, 이는 재귀 호출의 최대 깊이를 의미합니다.

### 출력
- 각 깊이마다 특정한 문장을 출력해야 하며, 깊이가 증가할 때마다 `____` (밑줄 4개)로 들여쓰기가 추가됩니다.
- `N`번째 깊이에서는 고정된 문장을 출력하고, 이후에는 역순으로 "라고 답변하였지."를 출력하며 마무리됩니다.

## 코드 분석

### 1. 입력 및 출력 설정
```go
var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)
```
- `bufio.Reader`와 `bufio.Writer`를 사용하여 **입출력 속도를 최적화**합니다.
- `os.Stdin`과 `os.Stdout`을 활용하여 입력을 읽고 출력을 빠르게 수행합니다.

### 2. 재귀 함수 정의
```go
func Recursion(num, depth int) {
    indent := ""
    for i := 0; i < depth; i++ {
        indent += "____"
    }
```
- `depth`에 따라 들여쓰기(`indent`)를 추가하여 출력 포맷을 맞춥니다.

#### 3. 재귀 호출 및 종료 조건
```go
    fmt.Fprintln(writer, indent+`"`+"재귀함수가 뭔가요?"+`"`)

    if depth == num {
        fmt.Fprintln(writer, indent+`"`+"재귀함수는 자기 자신을 호출하는 함수라네"+`"`)
    } else {
        fmt.Fprintln(writer, indent+`"`+"잘 들어보게. 옛날옛날 한 산 꼭대기에 이세상 모든 지식을 통달한 선인이 있었어.")
        fmt.Fprintln(writer, indent+"마을 사람들은 모두 그 선인에게 수많은 질문을 했고, 모두 지혜롭게 대답해 주었지.")
        fmt.Fprintln(writer, indent+"그의 답은 대부분 옳았다고 하네. 그런데 어느 날, 그 선인에게 한 선비가 찾아와서 물었어."+`"`)
        Recursion(num, depth+1)
    }
```
- `depth == num`일 때는 고정된 답변("재귀함수는 자기 자신을 호출하는 함수라네")을 출력하고 종료합니다.
- 그렇지 않으면 재귀적으로 `Recursion(num, depth+1)`을 호출하여 다음 깊이의 출력을 생성합니다.

#### 4. 종료 시 출력
```go
    fmt.Fprintln(writer, indent+"라고 답변하였지.")
}
```
- 재귀 호출이 끝나면 역순으로 "라고 답변하였지."를 출력하며 마무리됩니다.

### 5. 메인 함수
```go
func main() {
    defer writer.Flush()
    var num int
    fmt.Fscan(reader, &num)

    fmt.Fprintln(writer, "어느 한 컴퓨터공학과 학생이 유명한 교수님을 찾아가 물었다.")
    Recursion(num, 0)
}
```
- `defer writer.Flush()`를 사용하여 프로그램 종료 시 모든 출력이 한 번에 처리되도록 합니다.
- `fmt.Fscan(reader, &num)`을 사용하여 정수를 입력받습니다.
- `Recursion(num, 0)`을 호출하여 문제의 요구 사항을 만족하는 출력을 생성합니다.

## 시간 복잡도 분석

이 코드는 `N`번의 재귀 호출을 수행하며, 각 호출마다 `fmt.Fprintln`을 여러 번 실행합니다. 따라서 시간 복잡도는 **O(N)** 입니다. (단, 문자열 덧셈(`+`)으로 인해 추가적인 비용이 있을 수 있음)

## 결론

이 문제는 **재귀 호출의 개념**과 **출력 패턴을 유지하는 방법**을 이해하는 데 도움을 줍니다. 핵심은 재귀적으로 호출하면서 적절한 들여쓰기를 추가하고, 역순으로 문장을 출력하는 것입니다.

