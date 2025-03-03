// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// )

// var reader = bufio.NewReader(os.Stdin)
// var writer = bufio.NewWriter(os.Stdout)

// func main() {
// 	defer writer.Flush()

// 	N := 0
// 	fmt.Fscan(reader, &N)
// 	tile := make([]int, N)

// 	for i := 0; i < N; i++ {
// 		if i == 0 {
// 			tile[i] = 1
// 		} else if i == 1 {
// 			tile[i] = 2
// 		} else {
// 			tile[i] = (tile[i-2] + tile[i-1]) % 15746
// 		}
// 	}

// 	fmt.Fprintln(writer, tile[N-1])

// }

package main

import (
	"bufio"
	"fmt"
	"os"
)

// 전역 입력 및 출력 버퍼 설정
var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush() // 프로그램 종료 시 버퍼에 저장된 데이터를 출력

	N := 0
	fmt.Fscan(reader, &N) // 사용자 입력 받기

	// N이 1 또는 2일 경우 미리 결과 출력 후 종료
	if N == 1 {
		fmt.Fprintln(writer, 1)
		return
	} else if N == 2 {
		fmt.Fprintln(writer, 2)
		return
	}

	// 피보나치 점화식을 활용하여 결과 계산
	// dp[n] = dp[n-1] + dp[n-2] (모듈러 연산 적용)
	a, b := 1, 2 // 초기값 설정 (N=1일 때 1, N=2일 때 2)
	c := 0       // 결과 저장 변수

	for i := 3; i < N+1; i++ {
		c = (a + b) % 15746 // 점화식 적용 (큰 수 방지를 위해 모듈러 연산)
		a, b = b, c         // 이전 두 값을 갱신하여 메모리 절약
	}

	// 최종 결과 출력
	fmt.Fprintln(writer, c)
}
