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

func BinarySearch(n int, maxGem int, gems []int) int {
	left, right := 1, maxGem
	result := maxGem

	for left <= right {
		mid := (left + right) / 2
		count := 0

		for _, gem := range gems {
			// 각 색깔의 구슬을 mid개 이하로 나눠주기 위해 필요한 아이 수
			count += (gem + mid - 1) / mid // == ceil(gem / mid)
		}

		if count > n {
			left = mid + 1 // 아이가 부족하므로 envy level을 늘려야 함
		} else {
			result = mid // 조건 만족, envy level을 줄여볼 수 있음
			right = mid - 1
		}
	}

	return result
}

func main() {
	defer writer.Flush()

	var n, m int
	fmt.Fscan(reader, &n, &m)

	gems := make([]int, m)
	maxGem := 0

	for i := range gems {
		fmt.Fscan(reader, &gems[i])
		if gems[i] > maxGem {
			maxGem = gems[i]
		}
	}

	fmt.Fprintln(writer, BinarySearch(n, maxGem, gems))
}
