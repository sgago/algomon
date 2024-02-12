package dynamic

/*
	What is Dynamic Programming?
	https://algo.monster/problems/dynamic_programming_intro

	The first 5 Fibonacci numbers are:

	Fibonacci(1) = 0
	Fibonacci(2) = 1
	Fibonacci(3) = 1
	Fibonacci(4) = 2
	Fibonacci(5) = 3

	So, the first five Fibonacci numbers are 0, 1, 1, 2, and 3.

	To calculate the numbers, we can
*/

import (
	"fmt"
	"testing"
)

func TestFib(t *testing.T) {
	// dp := FibDp(10)
	// fmt.Println(dp)

	lis := LisDpAgain([]int{0, 1, 3, 2})
	fmt.Println(lis)
}

func FibDp(n int) int {
	// This ain't brute, we actually just did DP by mistake lol.
	// We're storing f(n-2) and f(n-1) for the next f(n)'s use.
	// Here, though, we only need to store two values, though.
	// I literally thought, welp, I just need to store the last two states,
	// and sum those for the next go.
	lastOne := 1
	lastTwo := 0
	result := 0

	for i := 0; i < n; i++ {
		if i <= 1 {
			result++
			continue
		}

		result = lastOne + lastTwo
		lastTwo, lastOne = lastOne, result
	}

	return result
}

func LisDpAgain(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	lis := 0

	/*
		0|0 -> 1 = 0+1 = f(init)+1 = 0+1
		1|1 -> 2 = 1+1 = f(0)+1 = 1+1
		2|3 -> 3 = 1+1+1 = f(1)+f(0)+1 = 1+1+1 = 3
		3|2 -> We don't add 1. It's still 3. It's 1+1+1 = f(0)+f(1)+1 again = 1+1+1 = 3
		     We take the lis at 2 instead. Why?
			 We just walk backwards until we find a number that 3 is greater than.
			 In this case, we skip 4, get 2.
			 So, just like in fib, we store f(init)=0
			 Next, we see idx 1, val 1. 1 > 0, so incr 1. Stash idx 1, lis 1 at 1.
			 Next we see 3 at idx 2. 3 > 1, so incr 1. We can grab stash[2-1] + 1.
			 Next we see 2. 2 is not > 3. We need to walk back until 2 is > than something. In this case 1.
	*/
	stash := make([]int, len(nums)+1)

	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			stash[i] = stash[i-1] + 1
			lis = max(lis, stash[i])
		} else {
			for j := i - 1; j > 0; j-- {
				// What am I doing here, truly? I'm trying to find the first index of
				// a number that's > nums[i]. I'm slowly walking j back until I find it.
				// We could do a binary search between i and j instead of a loop.
				if nums[i] > nums[j] {
					stash[i] = stash[j] + 1
					lis = max(lis, stash[i])
				}
			}
		}
	}

	return lis
}
