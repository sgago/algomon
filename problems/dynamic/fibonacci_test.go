package dynamic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibonacciNumber3(t *testing.T) {
	actual := Fib(3)
	assert.Equal(t, 3, actual)
}

func TestFibonacciNumber4(t *testing.T) {
	actual := Fib(4)
	assert.Equal(t, 5, actual)
}

func TestFibonacciNumber5(t *testing.T) {
	actual := Fib(5)
	assert.Equal(t, 8, actual)
}

func Fib(n int) int {
	return fib(n, make(map[int]int))
}

func fib(n int, memo map[int]int) int {
	if n == 0 || n == 1 {
		return 1
	}

	if val, ok := memo[n]; ok {
		return val
	}

	memo[n] = fib(n-1, memo) + fib(n-2, memo)

	return memo[n]
}

/*
Fibonacci numbers are not the whole tree nonsense.
It's fib(n) = fib(n-1) + fib(n-2) with f(0) = f(1) = 1.
*/
