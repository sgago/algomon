package steve

func nsort(min int, arr []int) []int {
	result := make([]int, len(arr))

	for _, num := range arr {
		result[num-min] = num
	}

	return result
}

func quick(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	pivot := arr[mid]

	left := make([]int, 0)
	right := make([]int, 0)

	for i, num := range arr {
		if i == mid {
			continue
		}

		if num < pivot {
			left = append(left, num)
			// Stats about left here
		} else {
			right = append(right, num)
			// Stats about right here
		}
	}

	right = append([]int{pivot}, right...)

	rl := quick(left)
	rr := quick(right)

	return append(rl, rr...)
}
