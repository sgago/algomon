package dynamic

import (
	"fmt"
	"testing"

	mstack "github.com/sgago/algomon/col/monotonic/stack"
)

/*
Longest Increasing Subsequence (LIS)
https://algo.monster/problems/longest_increasing_subsequence

Input
	nums: the integer sequence
Output
	the length of longest increasing subsequence

Examples
Example 1:
Input:
	nums = [50, 3, 10, 7, 40, 80]
Output: 4

Explanation:
The longest increasing subsequence is [3, 7, 40, 80] which has length 4.

Example 2:
Input:
	nums = [1, 2, 4, 3]
Output: 3

Explanation:
Both [1, 2, 4] and [1, 2, 3] are longest increasing subsequences which have length 3.

Mmk, so why do we think this is DP, outside of being in the DP problems section (lol)?
Do we have overlapping subproblems? Yes, yes, we definitely do.
The LIS of num[i] is going to be f(i)=max(num[i-1], num[i-2], ... num[0])+1.
Furthermore, it has keywords "longest" and "sequence" are DP indicators.
Most nums[i] in a sequence depend on the nums before it.


i|v
0|1 -> 1. Why? 0+1. Our init state is 0.
1|2 -> 2. Why? 1+1. We grab the previous state and add 1. f(0)+1=1+1
2|4 -> 3. Why? 2+1. We grab the previous state and add 1. f(1)+1=f(0)+1+1=1+1+1
3|3 -> 3. Why? 2+1. We don't grab the prevous cause it's greater! We grab the previous max.
          max(f(2), f(1))


Let's look at the array [1, 2, 4, 3].
f(-1) = 0 -> Initial conditions, always 0
1|f(0) = 1 -> f(-1)+1
2|f(1) = 2 -> f(0)+1
4|f(2) = 3 -> f(1)+1
3|f(3) = 3 -> f(1)+1! We need to pick whichever state has the max seq.

*/

func TestBrute(t *testing.T) {
	brute := LisBrute([]int{1, 2, 4, 3})
	fmt.Println(brute)
}

func LisBrute(nums []int) int {
	m := 0

	for i := 0; i < len(nums); i++ {
		lis := 0

		// How far from i, can we extend out?
		for j := i + 1; j < len(nums) && nums[j] > nums[i]; j++ {
			lis++
		}

		m = max(lis, m)
	}

	return m
}

func TestLisDp(t *testing.T) {
	dp := LisDp([]int{1, 2, 4, 3})
	fmt.Println(dp)
}

func LisDp(nums []int) int {
	dp := make([]int, len(nums)+1) // Create memo, with +1 for init cond
	// dp[0] = 0 // It's zero by default, but our memo's init cond is 0
	lis := 0

	for i := 1; i < len(nums); i++ {
		ni := nums[i-1]
		dp[i] = 1 // Any number is at least 1 or dp[0]+1==1

		for j := 1; j < i; j++ {
			nj := nums[j-1]

			if nj < ni {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}

		lis = max(lis, dp[i])
	}

	return lis
}

func TestLisDfs(t *testing.T) {
	dfs := LisDfs([]int{0, 1, 3, 2, 4, 5, -1, 0, 3})
	fmt.Println(dfs)
}

func LisDfs(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	memo := make(map[int]int)
	memo[0] = 1

	dfs := lisDfs(nums, len(nums)-1, memo)

	return dfs
}

func lisDfs(nums []int, idx int, memo map[int]int) int {
	if num, ok := memo[idx]; ok {
		return num // Already computed, return result
	}

	/*
		This is the struggle, right here. Truly, the recurrence relation makes or breaks you.
		We'll need to list out preciesly how to calculate a subproblem from other subproblems.
		f(4) = max(nums[i]...), but only for nums < nums[4], where i < 4
	*/

	lis := 1
	memo[idx] = 1 // Begin a new LIS

	for j := 1; j < idx; j++ {
		memo[j] = lisDfs(nums, j, memo)

		if nums[j-1] < nums[idx] {
			lis = max(lis, memo[j]+1)
		}
	}

	return lis
}

func TestLisMonotonic(t *testing.T) {
	s := mstack.New[int](10)

	lis := 0

	for _, num := range []int{0, 1, 3, 2, 4, 5, -1, 0, 3} {
		lis = max(lis, len(s.Push(num))+s.Len()-1)
	}

	fmt.Println(lis)
}
