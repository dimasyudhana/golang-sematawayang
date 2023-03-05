package main

import "fmt"

func runningSum(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	ans[0] = nums[0]
	for i := 1; i < n; i++ {
		ans[i] = ans[i-1] + nums[i]
	}
	return ans
}

func main() {
	var nums = []int{1, 2, 3, 4}
	fmt.Print(runningSum(nums))
}
