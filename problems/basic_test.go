package problems

import (
	"container/heap"
	"fmt"
	"slices"
	"sort"
	"testing"
)

func TestArrays(t *testing.T) {
	// An array of 5 empty elements
	var a [5]int
	fmt.Println("emp:", a)

	// Declare and init an array in one line
	b := [5]int{0, 1, 2, 3, 4}
	fmt.Println("dcl:", b)

	// A 2D array declare and init'ed in one line
	c := [3][3]int{
		{0, 1, 2},
		{3, 4, 5},
		{6, 7, 8},
	}

	fmt.Println("2d:", c)
}

func TestSlices(t *testing.T) {
	// Make a slice
	s := make([]int, 5)
	fmt.Println("emp", s)

	// Copy a slice
	copy(s, []int{0, 1, 2, 3, 4})
	fmt.Println("cpy", s)

	// Has a slice operator slice[low:high]
	// where low and high are inclusive and exclusive, respectively.
	// So s[2:5] would return s[2], s[3], and s[4]. Not s[5]!
	fmt.Println("s[2:5]", s[2:5])
	fmt.Println("s[2:]", s[2:])
	fmt.Println("s[:3]", s[:3])
}

func TestSorting(t *testing.T) {
	nums := []int{5, 3, 1, 4, 2}
	strs := []string{"def", "xyz", "abc"}

	// Method 1: The slices.Sort (newer), a generic func
	slices.Sort(nums)
	slices.Sort(strs)

	fmt.Println("nums:", nums)
	fmt.Println("strs:", strs)

	// Method 2: For sorting other things with a compare func
	cmp := func(a, b int) int {
		if a == b {
			return 0
		} else if a > b {
			return 1
		}

		return -1
	}

	slices.SortFunc(nums, cmp)

	// Method 1: The sort.Ints (old)
	sort.Ints(nums)
}

func TestChannels(t *testing.T) {
	messages := make(chan string)

	go func() {
		// Sends "ping" into messages channel
		messages <- "ping" // sender blocks until the receiver is ready
	}()

	// Receive a channel value into messages. Or like Send any channel values into this var.
	msg := <-messages // receiver blocks until sender is ready

	// Again, sends and recieves block, allowing us to wait at the end of the func for the value.
	fmt.Println("rcvd:", msg)

	// This'll panic with "all goroutines are asleep - deadlock!"
	//bad := <-messages
	//fmt.Println("bad", bad)

	// Unbuffered channels are strict synchronization point where both sender
	// and receiver must be ready. Use unbuffered channels when you need strong
	// synchronization between sender and receiver, the ensure data are shared.
}

func TestChannelBuffer(t *testing.T) {
	// Channels are unbuffered, meaning they will only accept sends (chan <-)
	// if there is a corresponding receive (<- chan).\
	// Buffered channels will receive a limited number of channels without a receiver.
	messages := make(chan string, 2) // This is a buffered channel, it has a capacity > 0

	go func() {
		// The sender chan send data to the channel without blocking, until it is full
		messages <- "buffered"
		messages <- "channel"
		//messages <- "this'll panic cause the buffer is full"
	}()

	// Similarily, we can receive from the channel without blocking until it's empty
	fmt.Println(<-messages)
	fmt.Println(<-messages)
	//fmt.Println(<-messages) // This'll panic too

	// We do get a degree of async freedom with buffered channels.
	// Useful where small delays or async is acceptable. Useful where small delays
	// or bursts of data are expected.
}

func TestChannelDirections(t *testing.T) {
	// When using channels as params, you can increase the type safety by specifying directions.

}

// Alias a slice type
type MinHeap []int

// And implement the heap.Interface type.
var _ heap.Interface = (*MinHeap)(nil)

func (h *MinHeap) Len() int {
	return len(*h)
}

func (h *MinHeap) Less(i int, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *MinHeap) Pop() any {
	result := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return result
}

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Swap(i int, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func TestHeap(t *testing.T) {
	// Create an empty MinHeap
	h := &MinHeap{}

	// Push elements onto the heap h, using heap.Push(h, whatever)
	heap.Push(h, 3)
	heap.Push(h, 1)
	heap.Push(h, 4)
	heap.Push(h, 2)

	// Pop elements from the heap h (retrieve in sorted order for a min-heap)
	// using the heap.Pop(h)
	for h.Len() > 0 {
		fmt.Printf("%d\v", heap.Pop(h))
	}
}
