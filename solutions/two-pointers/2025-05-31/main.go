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
		mostSimilarNumber = mostSimilarNumber * (-1)
	}

	for left < right {
		num := nums[right] + nums[left]
		check[k-num]++
		if k-num < 0 && mostSimilarNumber > -(k-num) {
			mostSimilarNumber = -(k - num)
		} else if k-num >= 0 && mostSimilarNumber > k-num {
			mostSimilarNumber = k - num
		}

		a, b := nums[right]+nums[left+1], nums[right-1]+nums[left]
		numA, numB := 0, 0
		if k-a < 0 {
			numA = -(k - a)
		} else {
			numA = k - a
		}
		if k-b < 0 {
			numB = -(k - b)
		} else {
			numB = k - b
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
