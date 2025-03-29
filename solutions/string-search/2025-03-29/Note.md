# 문제: 터널 내 추월 확인
👉🏻[문제 링크](https://www.acmicpc.net/problem/2002)

대한민국을 비롯한 대부분의 나라에서는 터널 내에서의 차선 변경을 법률로 금하고 있다. 조금만 관찰력이 있는 학생이라면 터널 내부에서는 차선이 파선이 아닌 실선으로 되어 있다는 것을 알고 있을 것이다. 이는 차선을 변경할 수 없음을 말하는 것이고, 따라서 터널 내부에서의 추월은 불가능하다.

소문난 명콤비 경찰 대근이와 영식이가 추월하는 차량을 잡기 위해 한 터널에 투입되었다. 대근이는 터널의 입구에, 영식이는 터널의 출구에 각각 잠복하고, 대근이는 차가 터널에 들어가는 순서대로, 영식이는 차가 터널에서 나오는 순서대로 각각 차량 번호를 적어 두었다.

N개의 차량이 지나간 후, 대근이와 영식이는 자신들이 적어 둔 차량 번호의 목록을 보고, 터널 내부에서 반드시 추월을 했을 것으로 여겨지는 차들이 몇 대 있다는 것을 알게 되었다. 대근이와 영식이를 도와 이를 구하는 프로그램을 작성해 보자.

---

## 입력

입력은 총 2N+1개의 줄로 이루어져 있다. 

1. 첫 줄에는 차의 대수 `N`(1 ≤ N ≤ 1,000)이 주어진다. 
2. 둘째 줄부터 `N`개의 줄에는 대근이가 적은 차량 번호 목록이 주어진다.
3. N+2째 줄부터 `N`개의 줄에는 영식이가 적은 차량 번호 목록이 주어진다. 

각 차량 번호는 6글자 이상 8글자 이하의 문자열로, 영어 대문자('A'-'Z')와 숫자('0'-'9')로만 이루어져 있다.

같은 차량 번호가 두 번 이상 주어지는 경우는 없다.

---

## 출력

첫째 줄에 터널 내부에서 반드시 추월을 했을 것으로 여겨지는 차가 몇 대인지 출력한다.

---

## 예제 입력 및 출력

### 입력 1
```
4
ZG431SN
ZG5080K
ST123D
ZG206A
ZG206A
ZG431SN
ZG5080K
ST123D
```

### 출력 1
```
1
```

### 입력 2
```
5
ZG508OK
PU305A
RI604B
ZG206A
ZG232ZF
PU305A
ZG232ZF
ZG206A
ZG508OK
RI604B
```

### 출력 2
```
3
```

### 입력 3
```
5
ZG206A
PU234Q
OS945CK
ZG431SN
ZG5962J
ZG5962J
OS945CK
ZG206A
PU234Q
ZG431SN
```

### 출력 3
```
2
```

---

## 코드 구현

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

func Overtake(input, output []string) int {
	res := 0

	check := make(map[string]bool)

	for i := 0; i < len(output); i++ {
		if check[input[0]] {
			input = input[1:]
			i--
			continue
		}
		if input[0] == output[i] {
			input = input[1:]
		} else {
			check[output[i]] = true
			res++
		}
	}

	return res
}

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)
	input := make([]string, n)
	output := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &input[i])
	}
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &output[i])
	}

	fmt.Fprintln(writer, Overtake(input, output))
}
```

---

## 풀이 설명

1. 차량 번호 목록 `input`(입구 차량 번호)와 `output`(출구 차량 번호)를 비교하여 추월한 차량을 찾는다.
2. `check` 맵을 사용하여 이미 추월한 차량을 기록한다.
3. 출구 차량 번호를 순회하며 다음을 확인한다:
   - `input` 목록에서 현재 출구 차량 번호와 일치하면, 해당 차량을 `input`에서 제거.
   - 일치하지 않는 경우, 해당 차량이 추월한 것으로 간주하고 `check`에 기록.
4. 추월한 차량의 수를 반환한다.

---

## 시간 복잡도

- 입력 크기 `N`에 대해 차량 번호 비교를 수행하므로 시간 복잡도는 `O(N)`이다.
