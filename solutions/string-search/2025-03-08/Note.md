# 파일 확장자 정리 프로그램
👉🏻[문제 링크](https://www.acmicpc.net/problem/20291)

## 문제 설명
친구로부터 노트북을 중고로 산 스브러스는 바탕화면에 온갖 파일들이 정리되지 않은 채 가득한 것을 보고 경악했습니다.
그러나 파일을 지우면 안 된다는 친구의 메시지를 확인하고, 파일들을 확장자별로 정리하여 개수를 세는 프로그램을 만들기로 했습니다.

스브러스의 요청 사항은 다음과 같습니다:

1. 파일을 확장자별로 정리하여 개수를 출력한다.
2. 확장자를 사전순으로 정렬하여 출력한다.

## 입력
- 첫째 줄: 바탕화면에 있는 파일의 개수 **N** (1 ≤ N ≤ 50,000)
- 둘째 줄부터 N개의 줄에 파일의 이름이 주어진다.
  - 파일 이름은 알파벳 소문자와 점(`.`)으로만 구성됨.
  - 점은 정확히 한 번 등장하며, 파일 이름의 첫 글자 또는 마지막 글자로 오지 않음.
  - 각 파일 이름의 길이는 최소 3, 최대 100.

## 출력
- 확장자별 개수를 한 줄에 하나씩 출력한다.
- 확장자는 사전순으로 정렬하여 출력한다.

## 예제 입력
```
8
sbrus.txt
spc.spc
acm.icpc
korea.icpc
sample.txt
hello.world
sogang.spc
example.txt
```

## 예제 출력
```
icpc 2
spc 2
txt 3
world 1
```

---

## 해결 방법

### 1. 입력값 처리
- `N`을 입력받은 후, `N`개의 파일 이름을 읽어들입니다.
- 파일 이름을 `.` 기준으로 나누어 확장자를 추출합니다.

### 2. 확장자 개수 세기
- `map[string]int` 타입의 **해시맵**을 사용하여 확장자별 등장 횟수를 기록합니다.
- 처음 등장하는 확장자는 별도의 **슬라이스**에 저장하여 중복 없이 정리합니다.

### 3. 확장자 정렬 후 출력
- 확장자를 **사전순으로 정렬**합니다.
- 정렬된 확장자 목록을 순회하며 개수를 출력합니다.

---

## 코드 구현
```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

	var n int
	fileExtension := make(map[string]int)
	files := make([]string, 0)

	fmt.Fscan(reader, &n)

	for i := 0; i < n; i++ {
		input := ""
		fmt.Fscan(reader, &input)
		fileName := strings.Split(input, ".")
		if fileExtension[fileName[1]] == 0 {
			files = append(files, fileName[1])
		}
		fileExtension[fileName[1]]++
	}

	sort.Slice(files, func(i, j int) bool {
		return files[i] < files[j]
	})

	for i := 0; i < len(files); i++ {
		fmt.Fprintln(writer, files[i], fileExtension[files[i]])
	}
}
```

---

## 코드 설명
1. **입력 처리**:
   - `fmt.Fscan(reader, &n)`을 사용하여 파일 개수를 입력받습니다.
   - 반복문을 통해 파일 이름을 입력받고 `strings.Split(input, ".")`을 이용해 확장자를 추출합니다.

2. **확장자 개수 저장**:
   - 해시맵 `fileExtension`을 사용하여 확장자별 개수를 저장합니다.
   - 새로운 확장자가 등장하면 `files` 슬라이스에 추가합니다.

3. **확장자 정렬**:
   - `sort.Slice(files, func(i, j int) bool { return files[i] < files[j] })`를 사용하여 확장자 목록을 사전순으로 정렬합니다.

4. **출력**:
   - 정렬된 확장자를 순회하며 `fmt.Fprintln(writer, files[i], fileExtension[files[i]])`을 이용해 출력합니다.

---

## 시간 복잡도 분석
- 파일 이름을 **한 번씩 순회**하며 확장자를 추출하고 저장하므로 `O(N)`
- 확장자 목록을 **정렬**하는 과정이 `O(M log M)` (M: 서로 다른 확장자의 개수)
- 최종적으로 `O(N + M log M)`의 시간 복잡도를 가짐

---

## 핵심 정리
✅ 해시맵(`map`)을 이용해 확장자별 개수를 효율적으로 카운트
✅ 확장자 리스트를 사전순 정렬 후 출력
✅ 시간 복잡도: `O(N + M log M)`, 대략 `O(N log N)` 수준으로 최적화됨
