package cycle

import "cmp"

func Sort[T cmp.Ordered](arr []T) {

}

func SortFunc[T any](arr []T, less func(i, j T) bool) {

}

// Function to perform Cycle Sort on an integer slice
func sortInt(nums []int) {
	n := len(nums)

	// Traverse the array to start building cycles
	for start := 0; start < n-1; start++ {
		// Pick an element as the starting point of the cycle
		item := nums[start]
		pos := start

		// Find the correct position for the current element in the cycle
		for i := start + 1; i < n; i++ {
			if nums[i] < item {
				pos++
			}
		}

		// If the element is already in the correct position, move to the next element
		if pos == start {
			continue
		}

		// Swap the elements to place the current element in its correct position
		nums[pos], item = item, nums[pos]

		// Rotate the rest of the cycle until the current element is back in its starting position
		for pos != start {
			pos = start

			// Find the correct position for the current element in the remaining cycle
			for i := start + 1; i < n; i++ {
				if nums[i] < item {
					pos++
				}
			}

			// Handle cases where there are duplicate elements
			for item == nums[pos] {
				pos++
			}

			// Swap the elements to continue rotating the cycle
			nums[pos], item = item, nums[pos]
		}
	}
}
