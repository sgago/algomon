# The Study Guide
Welcome to my personal study guide for leetcode problems and systems design.

- [The Study Guide](#the-study-guide)
- [Data structures and algorithms](#data-structures-and-algorithms)
  - [Math](#math)
  - [Runtimes](#runtimes)
    - [Summary](#summary)
    - [O(1)](#o1)
    - [O(logN)](#ologn)
    - [O(KlogN)](#oklogn)
    - [O(N)](#on)
    - [O(KN)](#okn)
    - [O(N + M)](#on--m)
    - [O(V + E)](#ov--e)
    - [O(NlogN)](#onlogn)
    - [O(N^2)](#on2)
    - [O(2^N)](#o2n)
    - [O(N!)](#on-1)
    - [Amortized](#amortized)
    - [Time complexity math](#time-complexity-math)
    - [Tricks](#tricks)
  - [Hash functions and maps](#hash-functions-and-maps)
    - [Common hash functions](#common-hash-functions)
    - [Pigeonhole principle](#pigeonhole-principle)
    - [References](#references)
  - [Sorting](#sorting)
    - [Summary](#summary-1)
    - [Go sort](#go-sort)
    - [Cycle sort](#cycle-sort)
      - [Semi sort](#semi-sort)
  - [Binary search](#binary-search)
    - [Binary search keywords](#binary-search-keywords)
  - [Trees](#trees)
    - [Terminology](#terminology)
    - [Binary trees](#binary-trees)
    - [Binary search trees](#binary-search-trees)
    - [Balanced and unbalanced binary trees](#balanced-and-unbalanced-binary-trees)
    - [Tree traversal](#tree-traversal)
    - [Tree keywords](#tree-keywords)
  - [Heaps](#heaps)
    - [Go heap](#go-heap)
  - [Depth first search](#depth-first-search)
  - [Backtracking](#backtracking)
  - [Graphs](#graphs)
  - [Dynamic programming](#dynamic-programming)
    - [Strategy for solving DP](#strategy-for-solving-dp)
    - [Top down or bottom up DP?](#top-down-or-bottom-up-dp)
    - [Tricks](#tricks-1)
  - [Disjoint union set (DSU)](#disjoint-union-set-dsu)
  - [Intervals](#intervals)
  - [Keywords](#keywords)
- [Systems design](#systems-design)
  - [Communication](#communication)
    - [The Internet protocol suite](#the-internet-protocol-suite)
    - [Border gateway protocol (BGP)](#border-gateway-protocol-bgp)
    - [User datagram protocol (UDP)](#user-datagram-protocol-udp)
    - [Reliable communications](#reliable-communications)
      - [Opening and the TCP handshake](#opening-and-the-tcp-handshake)
      - [Closing and the TIME\_WAIT](#closing-and-the-time_wait)
      - [Established connections and congestion control](#established-connections-and-congestion-control)
    - [Secure communications](#secure-communications)
      - [Encryption](#encryption)
      - [Authentication and certificates](#authentication-and-certificates)
      - [Data integrity](#data-integrity)
    - [Discovery and DNS](#discovery-and-dns)
    - [Application programming interfaces](#application-programming-interfaces)
      - [RESTful HTTP requests and responses](#restful-http-requests-and-responses)
      - [Synchronous HTTP requests](#synchronous-http-requests)
      - [Asynchronous HTTP requests](#asynchronous-http-requests)
      - [Synchronous HTTP responses](#synchronous-http-responses)
      - [Asynchronous HTTP responses](#asynchronous-http-responses)
      - [HTTP versions](#http-versions)
    - [Messaging](#messaging)
  - [Scalability](#scalability)
    - [Replication](#replication)
      - [Consistency models](#consistency-models)
        - [Linearizability](#linearizability)
      - [Strong consistency](#strong-consistency)
        - [Sequential consistency](#sequential-consistency)
        - [Causal consistency](#causal-consistency)
          - [The CALM theorem](#the-calm-theorem)
        - [Back to causal consistency](#back-to-causal-consistency)
        - [Eventual consistency](#eventual-consistency)
    - [CAP and PACELC theorems](#cap-and-pacelc-theorems)
    - [HTTP caching](#http-caching)
    - [Partitioning](#partitioning)
      - [Ranged partitioning](#ranged-partitioning)
      - [Hash partitioning](#hash-partitioning)
    - [Load balancing](#load-balancing)
      - [Load balancing strategies](#load-balancing-strategies)
      - [Layer 4 and 7 load balancing](#layer-4-and-7-load-balancing)
      - [Load balancing failover](#load-balancing-failover)
      - [Load balancer uptime](#load-balancer-uptime)
      - [Service discovery](#service-discovery)
    - [Forward and reverse proxies](#forward-and-reverse-proxies)
    - [Content delivery networks (CDN)](#content-delivery-networks-cdn)
      - [CDN networks](#cdn-networks)
      - [CDN caching](#cdn-caching)
      - [Push and pull CDNs](#push-and-pull-cdns)
    - [Blob storage](#blob-storage)
      - [Blob vs File](#blob-vs-file)
        - [Traditional hierarchical file-based systems](#traditional-hierarchical-file-based-systems)
        - [Blob storage](#blob-storage-1)
      - [Blob storage](#blob-storage-2)
      - [Blob access](#blob-access)
    - [Microservices](#microservices)
      - [API gateway](#api-gateway)
        - [GraphQL](#graphql)
  - [Reliability](#reliability)
    - [Common failures](#common-failures)
    - [Risk](#risk)
    - [Redundancy](#redundancy)
    - [Fault isolation](#fault-isolation)
      - [Shuffle sharding](#shuffle-sharding)
      - [Cellular architecture](#cellular-architecture)
      - [Pool architecture](#pool-architecture)
  - [Monitoring](#monitoring)
    - [Pre-aggregation](#pre-aggregation)
    - [Service level terminology](#service-level-terminology)
    - [Service level indicators](#service-level-indicators)
- [Napkin math](#napkin-math)
  - [Costs](#costs)
  - [Uptime in nines](#uptime-in-nines)
  - [Little's Law](#littles-law)
  - [The rule of 72](#the-rule-of-72)
  - [Sorting](#sorting-1)
  - [Stacks and queues](#stacks-and-queues)
  - [Data storage](#data-storage)
  - [Networking](#networking)
  - [Computations](#computations)
  - [DBs with SQL](#dbs-with-sql)
  - [DBs with NoSQL](#dbs-with-nosql)
  - [RAM disk providers](#ram-disk-providers)
  - [Storage devices](#storage-devices)
  - [Serialization](#serialization)
  - [Hashing](#hashing)
  - [Common HTTP status codes](#common-http-status-codes)
- [General references](#general-references)

# Data structures and algorithms
[Top](#the-study-guide)

## Math
There's some math to brush up on first. Most of it's high-school level, fortunately!
- **Logarithms:** To what power must we raise a certain base to get a number? We frequently deal with base Log base 2 here, which tells us how many times we need to multiply a number to reach a certain value. For example, `Log2(8) = 3` (`8 = 2 * 2 * 2 = three 2's`) reads like, "How many times do we need to multiple 2 to get 8?". Be aware of context! In mathematics classes, we usually assume Log base 10 when the log base is omitted, but for computers we are *probably* assuming Log base 2!
- **Permutations and factorials:** These come up in combinatorial problems and similar. Say we want to count the number of unique ways we can arrange the letters a, b, and c. There's 6 possible ways (sequences): `3 * 2 * 1` or `3!`. Rarely, does something more complicated come up, but it does happen like with [shuffle sharding](#shuffle-sharding). For example, with a total of 6 items, how many ways can we arrange them in groups of two? This is given by `6!/(2!4!) = 15` or `N!/R!(N - R)!` where `N` is the total elements and `R` is the size of the subsets.
- **Arithmetic sequences** An arithmetic sequence is a sequence of numbers where the difference between consecutive terms is constant. Like `1, 2, 3, 4, 5` or `2, 4, 5, 8, 10`. The sum of an arithmetic sequence can be quickly calculated with `(first_element + last_element) * number_of_elements / 2`.

## Runtimes
How long an algorithm takes to run for a given input. Also called "time complexity" or "TC" for short.

### Summary
Runtime | Name | Example
--- | --- | --- 
`O(1)` | Constant | Math, assignments
`O(alpha(N))` | Inverse Ackermann | Rare. Close to constant time; think O(4) at most. Appears is Disjointed Set Union.
`O(logN)` | Log | Binary search, binary search tree search
`O(KlogN)` | K linear | K binary searches
`O(N)` | Linear | Traverse array, tree traversal
`O(KN)` | K Linear | Performing a linear operation K times
`O(N + M)` | Linear NM | Traverse array, tree traversal on two separate collections
`O(V + E)` | Graph | Traverse graph with V vertices and E edges
`O(NlogN)` | Sort | Quick and merge sort, divide n' conquer
`O(N^2)` | Quadratic | Nested loops
`O(2^N)` | Exponential | Combinatorial, backtracking, permutations
`O(N!)` | Factorial | Combinatorial, backtracking, permutations
Amortized | Amortized | High order terms that are rarely done, smaller order done more frequently. Like doing O(N^2) once at startup and O(logN) every other time.

### O(1)
Constant time. A constant set of of operations.
- Hashmap lookups due to pointer arithmetic magic
- Array access
- Math and assignments
- Pushing and popping from a stack

```go
// Some non-looping, non-recursive code
i = 1
i = i + 100000
fmt.Println("more constant time code")
```

### O(logN)
Log time. Grows slowly. log(1,000,000) is only about 20.
- Binary searches, due to constantly splitting solution space in half
- Balanced binary search tree lookups, again cause it's halved.
- Processing number digits
Unless specified, in programmer world, we mean log base 2.

```go
for i := N; i > 0; i /= 2 {
    // Constant time code
}
```

### O(KlogN)
Typically, when you need to do a log(N) process K times. Examples:
- Heap push/pop K times, like merging N sorted lists
- Binary search K times

### O(N)
Linear time. Typically, looping through a data struct a constant number of times. Examples:
- Traversing an entire array or linked list.
- Two pointer solutions
- Tree or graph traversal due to visiting all the nodes or vertices
- Stack and queue

```go
for i := 0; i < N; i++ {
    // Constant time code
}
```

### O(KN)
Typically, when you need to process N K times. Very exciting.

### O(N + M)
Typically, when you have two inputs of size N and M. Say you loop once N times and then loop M times.
Again, very exciting.

### O(V + E)
For both DFS and BFS on a graph, the time complexity is O(|V| + |E|), where V is the number of vertices and E is the number of edges. The number of edges in a graph could be 1 to |V|^2, we really don't know. So we include both terms here.

### O(NlogN)
When we need to do a logN time process N times.
- Divide and conquer, where divide is logN and merge is N
- Sorting can get down to this.

### O(N^2)
Quadratic time. Not terrible for N < 1000, but does grow quickly.
Usually, interviewers want better than this. If you've come up
with a N^2 runtime solution, there's probably something better.
- Nested loops, where outer and inner loops run N times

```go
for i := 0; i < N; i++ {
    for i := 0; i < N; i++ {
        // Constant time code
    }
}

// OR

for i := 0; i < N; i++ {
    for j := i; j < N; j++ {
        // Constant time code
    }
}
```

The bottom loop is tricky because it's a factor of N.
So it's N^2 too.

### O(2^N)
Grows very rapidly and often requires memoization to reduce runtime.
- Combinatorial problems, backtracking, and subsets
- Often involves recursion

Note, this one is harder to analyze at first.

### O(N!)
Grows insanely rapidly. Only solvable for small N and typically requires memoization.
- Combinatorial problems, backtracking, permutations.

Note, this one is hard to analyze/spot at first.

### Amortized
Amortized time, meaning to gradually write off the initial time costs, if an operation is rarely done. For example, if we had N^2*N tasks, we could consider the solution O(N) instead of O(N^2) = N*O(1) * O(N) if we only do the N^2 task rarely. For example, if we dynamically size an array one time at startup.

### Time complexity math
N is *sort of* like infinity. It swallows smaller N terms and constants. Unlike infinity however, it does not swallow up other values if multiplication is involved, like with NlogN.
- 2N -> N
- N + logN -> N
- NlogN -> NlogN
- 3N^3 + 2N^2 + N -> N^3
- N^2 + 2^N + N! -> N!

Note that, like the other runtimes, we ignore constant factors and lower order terms: 5N^2 + N = 5N^2 -> N^2.

### Tricks
The maximum number of elements, usually noted in the constraints, will give you *clues* about the runtime and *maybe* the corresponding solution. Why? The leetcode runners are docker containers with a very short shelf life by design. There's a maximum number of operations they'll let your run. If your code takes too long, then leetcode will kill your container and return a time limit expired (TLE) error. So, they actually have to limit the number of input elements.

N Constraint | Runtimes | Description
--- | --- | ---
N < 20 | 2^N or N! | Brute force or backtracking.
20 < N < 3000 | N^2 | Dynamic programming.
3000 < N < 10^6 | O(N) or O(NlogN) | 2 pointers, greedy, heap, and sorting.
N > 10^6 | O(logN) or O(1) | Binary search, math, cycle sort adjacent problems.

Again, guessing the optimal solution from the size of input elements constraint is certainly error prone. This could *suggest a maybe solution*.

## Hash functions and maps
In short, a hash function converts arbitrary sized data into a fixed value, typically an Int32. For example, summing all integers in an array and mod'ing them by 100. We convert a bunch of data or text into a smaller, ergonomic number.

Typically, you don't have to write hash functions from scratch outside of, say, for `GetHashcode` for C# objects or similar.

A hash collision occurs when a hash function generates the same hash value for different data.

A good hash function is typically:
- Fast to compute (low time complexity or runtime)
- Very low chance of collision
- All possible values have a somewhat equal chance of occurring.

### Common hash functions
Name | Description
--- | ---
SHA | Cryptographic hash. SHA-3 is the latest. SHA-2 and -1 have vulnerabilities now.
Blake2 | An improvement over SHA-3, high speed and used in crypto mining.
Argon2 | Password hashing function designed to be resistant to brute force or dictionary attacks. Uses large amount of memory (memory-hard) to make attacks more difficult, for hackers using specialized hardware to crack passwords.
MurmurHash | Fast an efficient non cryptographic has. Useful for hash tables.
CRC | Cyclic redundancy check. Non cryptographic. Fast, but not used for security. Typically, the CRC is appended to messages, like HTTP, to check for corruption.
MD5 | Fast 128 bit hash, but no longer recommended for security.

### Pigeonhole principle
Collisions are unavoidable, so we need to design around it. For example, if "anne" and "john" created the same hash, we'd overwrite the same hash table entry. To avoid this, we can use [separate chaining](https://en.wikipedia.org/wiki/Hash_table#Separate_chaining) or other strategies.

### References
[List of hash functions](https://www.geeksforgeeks.org/hash-functions-and-list-types-of-hash-functions/)

## Sorting
- Time complexity = The amount of time it takes to sort the collection as a function of the size of the input data, represented in big O notation. Basic sorting is usually N^2, advanced are usually NlogN.
- Stability = If two elements have equal keys, then the order of these elements remains unchanged. This can be valuable for historical data, user expectations, or multi-criteria sorts where not sorting equal elements is important.
- In-place = The sorting algorithm sorts the input data structure without need to allocate additional memory to store the sorted results. This is valuable for large data sets.
- Simple = Simple algorithms are those that are relatively straightforward to implement with a one or two loops. The more complicated algorithms like quick and merge sort use divide and conquer strategies. This does not mean they are super easy to just bang out, however.
- Adaptable = Sometimes, input data are already somewhat sorted and we can minimize the number of comparisons that we need to make. For example, gnome sort is O(N) for an already sorted collection!
- Parallelizable = The sorting algorithm can divide the the sorting into subtasks that can be executed in parallel. Merge, quick, radix, and bucket sorts are all parallelizable due to divide and conquer.

### Summary
A summary of common algorithms, courtesy of ChatGPT.
| Algorithm | Time Complexity (Worst Case) | Stable | In-Place | Adaptable | Parallelizable | Description |
| --- | --- | --- | --- | --- | --- | --- |
| [Bubble](./sort/bubble/bubble.go) | O(n^2) | Yes | Yes | Yes | Yes (limited) | Repeatedly compares and swaps adjacent elements.          |
| [Selection](./sort/selection/selection.go) | O(n^2) | No | Yes | No | Yes (limited) | Repeatedly selects the minimum element and swaps with the current position. |
| [Insertion](./sort/insertion/insertion.go) | O(n^2) | Yes | Yes | Yes | Yes (limited) | Builds the sorted list one element at a time by inserting into the correct position. |
| [Merge](./sort/merge/merge.go) | O(n log n) | Yes | No | Yes | Yes | Divides the input into halves, recursively sorts them, and merges them. |
| [Quick](./sort/quick/quick.go) | O(n^2) (rare), O(n log n) | No | Yes | Yes | Yes (limited) | Chooses a pivot, partitions the data, and recursively sorts the partitions. |
| Heap | O(n log n) | No | Yes | No | Yes | Builds a binary heap and repeatedly extracts the maximum element. |
| Shell | O(n log^2 n) (worst known) | No | Yes | No | No | A variation of insertion sort with multiple passes and varying gap sizes. |
| Radix | O(nk) (k is the number of digits) | Yes | Yes | No | Yes | Processes digits or elements in multiple passes, each pass sorted independently. |
| Bucket | O(n^2) (worst case) | Yes | No | Yes | Yes | Distributes elements into buckets and sorts each bucket independently. |
| [Cycle](./sort/cycle/cycle.go) | O(N^2) | No | Yes | No | No | Based on the idea that elements can be sorted by setting them to their correct index position. Theoretically optimal in terms of writes.

### Go sort
We don't typically author sort algorithms from scratch in production. Some computer scientist has implemented pattern defeating quick sort (pdqsort) for us so that we don't have to. In go, use `slices.Sort` or similar:
```go
arr := []int{5, 3, 1, 4, 2}
slices.Sort(arr)

// OR

cmp := func(a, b int) int {
  if a == b {
    return 0
  } else if a > b {
    return 1
  }

  return -1
}

slices.SortFunc(arr, cmp)
```

### Cycle sort
We know a collection is sorted if all elements are in increasing or decreasing order, but how would we know if an individual element is sorted? Naively, we might think `left element < the element being considered < right element` but this won't work for a collection like `5, 1, 2, 3, 4`. We know an element is sorted if it's index is equal to a count of all elements smaller than it. For example, in `1, 2, 3, 4, 5`, we know 2 is sorted because its index in the slice is equal a count of all numbers less than it, in this case, just 1.

[Cycle sort](./sort/cycle/cycle.go) is based on this idea of elements being at their proper index position.

#### Semi sort
There's a class of problems that are solved via semi-sorting elements. For example, the [find the minimum missing positive number](https://leetcode.com/problems/first-missing-positive/description/) like problems can be solved this way. (Warning: This semi sort uses O(N) space and it isn't an optimal solution for this problem!) This problem can be solved faster than O(NlogN) time if we only partially, kind-of sort the array. In short, we simply move elements to the same index as their value ([here](./sort/semi/semi.go)). In a new slice, we copy 0 to index position 0, 1 to index position 1, etc. "But what about index out of range errors!?" If element values is not in-range, we can just ignore them, copy the value as is, or use a bad value.

## Binary search

```go
func binarySearch(nums []int) {
  ans := 0
  left, right := 0, len(nums)-1

  for left <= right {
    mid := left + (right-left)/2

    if nums[mid] > nums[left] {
      // Pull in left side
      left = mid + 1
    } else if nums[mid] <= nums[right] {
      // Pull in right side
      right = mid - 1
      ans = mid // Maybe get answer from here
    }
  }

  return ans
}
```

### Binary search keywords
These keywords may indicate a binary search if they appear in the story problem text.
- Sorted integers, sorted strings
- Rotated array, find peak
- Any array/slice that could be considered implicitly sorted

## Trees
Trees are a type of graph composed of nodes and edges.
- Trees are acyclic, nodes don't loop back to themselves and create cycles.
- There's a path from the root node to any other node.
- Trees have N-1 edges, when N is the number of nodes.
- Nodes have exactly one parent node.
- Trees are directed.
- Trees are rooted.

### Terminology
Name | Description
--- | ---
Root node | The top most ancestor node, one with out any parents.
Internal node | Every node that has at least one child.
Leaf node | Every node that does not have any children.
Ancestor | All nodes between the path from the child to the parent.
Descendent | All nodes between the path from a parent to the child.
Level | Number of ancestors from the current node to the root.
Arity | Number of operators or terms. For trees, where each node has no more than N-ary children.

I dislike ancestor/descendent. Call 'em parent and child like we're all 5 years old.

Example of a tree:
```
    A
   / \
  B   C
 / \   \
D   E   G
```
- A through G are all nodes.
- / and \ are edges. There's N-1 = 6-1 = 5 edges and 6 nodes.
- A is the root node.
- D, E, and G are all leaf nodes.
- A, B, and C are internal nodes.
- A is at level 0, B and C at level 1, and D, E, and G at level 2.
- A is a parent of B and C. B is a parent of D and E. C is a parent of G. D, E, and G are leaf nodes and have no children themselves.
- D and E are a child of B. G is a child of C. B and C are children of A. A is the root node and has no parents.

### Binary trees
Binary trees are a tree where each node has 0 to 2 children.

A full binary tree is one in which every node has 0 or 2 children. 1 is not allowed.
```
      x
     / \
    x   x
   / \
  x   x
 / \
x   x
```

A complete binary tree is where all the levels, except the last, are completely filled out. In the last level, the nodes are as far left as possible. This shows up in heaps.
```
      x
     / \
    x   x
   / \
  x   x
```

A perfect binary tree is one in which all the internal nodes have exactly 2 children. 1 or 0 are children are not allowed. All leaf nodes have the same number of children.
```
       x
     /   \
    x     x
   / \   / \
  x   x x   x
```
Perfect trees are used to estimate time complexity for combinatorial problems where the search space is a perfect binary tree. They have some unique properties.
- The number of nodes is 2^L-1 where L is the number of levels.
- The number of internal nodes is # of leaf nodes - 1.
- The total number of nodes is = 2 * leaf nodes - 1.

### Binary search trees
Binary search trees (BSTs) are a special type of binary tree where all left descendants < node < all right descendants.
```
      8
     / \
    3   10
   / \    \
  1   5    14
       \
        7
```
Notice that 3 is to the left of 8 because 3 < 8. Similarly, 14 is to the right of 10.
Note that the in-order traversal of the tree visits the nodes in monotonically increasing order.

### Balanced and unbalanced binary trees
Unbalanced trees are have a search time of N. The start to look more like a list than a tree.
```
1
 \
  2
   \
    3
     \
      4
```
Balanced binary trees are those where the difference in height (levels) between the left and right subtrees of all nodes is not more than 1. Balanced trees allow for a search time of logN.
```
      8
     / \
    3   10
   / \    \
  1   5    14
       \
        7
```

Balanced binary trees include [red-black](https://en.wikipedia.org/wiki/Red%E2%80%93black_tree) and [AVL](https://en.wikipedia.org/wiki/AVL_tree) trees.

### Tree traversal
Tree traversal are types of traveling through the nodes of a tree.
- In-order visits the left branch, current node, then right branch.
- Pre-order visits the current node, left subtree, and right subtree.
- Post-order visits the left subtree, right subtree, and current node.

For example, given the following binary search tree,
```
     8
    / \
   3   10
  / \    \
 1   5    14
      \
       7
```
the visits to each node would be:
- In-order: 1 3 5 7 8 10 14
- Pre-order: 8 3 1 5 7 10 14
- Post-order: 1 7 5 3 14 10 8

### Tree keywords
These keywords may indicate a tree if they appear in the story problem text.
- Shortest, level-order, zig-zag order, etc.
- Max depth

## Heaps
A min heap is a special tree data structure where
1. Almost complete - every level in the tree is almost filled, except the last level. The last level is left justified.
1. Each node has a greater key (priority) than it's parent.

Here's a heap structure:
```
       1
     /   \
    2     3
   / \   / \
  7   8 9   11
  |
  12
```

This is not a heap:
```
       1
     /   \
    2     3
   / \   / \
  7   8 9   11
  |
  5 // Need to heapify this, 5 < 7, so our heap property is broken
```

Also, not a heap:
```
       1
     /   \
    2     3
   / \   / \
  7   8 9   11
      |
      12 // Need to left justify this
```

Also, not a heap:
```
       1
     /   \
    2     3
   / \   /
  7   8 9 // Need to fill intermediate levels
  |
  12
```

Couple notes:
- Priority queue is an abstraction over a heap, minheap and maxheap are the concrete implementations.
- Max heaps are the same, but we just change the each-child-key-is-greater property to each child is less.
- Usually, heaps are binary tree, but you can also get k-ary heaps or k-heaps.
- A priority queue is an abstraction on the heap, a min/max heap is the concrete implementation.
- They are typically implemented with an array.
- It's a sorted tree but it is not a binary search tree.

Heaps support three main operations:
1. Heapify - Rearrange the nodes such that the heap such that the nodes are in min keys are always the at the root. This is an O(logN) operation.
2. Insert - Inserts a new element into the heap and calls heapify because to maintain the heap properties. This is an O(logN) operation.
3. Pop/Delete - Removes and returns the min root element. Another O(logN) operation.

In it's use, it's sort of like a stack. Push nodes in, pop nodes out except you always get the min keyed node.

### Go heap
Like sorting, we don't have to author heaps from scratch in go.
For example, we can alias a slice type and implement heap.Interface on it:
```go
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
```

## Depth first search
A depth first search (DFS) looks for solutions by going deep first. That is, it searches for solutions in a pre-order traversed way. Some more terminology:
- Backtracking - returning after visiting a non-solution node
- Divide and conqueror - When we have two or more recursive calls, that splits our issue into subproblems, like O(logN).

When do we use a DFS? We can use them in
1. Trees
  - Traverse through a tree to find, create, delete nodes
  - Traverse through a tree to find the max subtree, detect a balanced tree, etc.
2. Combinatorial problems
  - Find all the different ways of arranging something
  - Find all possible combinations of something 
  - Find all solutions to a puzzle
3. Graphs
  - Find a path from A to B in a graph
  - Find connected components
  - Detect cycles

There are two main ways of coding a DFS. We can use recursion or we can use a stack + loop.
1. With recursion, we call the recursive func on the next state.
2. With a loop, we pop off a stack, add any possible solutions, and then keep going.

One thing to keep in mind with DFS, is kicking state into and out of the recursive functions and/or stack.

The template for a DFS via recursion is
```
function dfs(node, state):
    if node is null:
        ...
        return

    left = dfs(node.left, state)
    right = dfs(node.right, state)

        ...

    return ...
```

The main hard parts of DFS are 1. deciding which state to pass in, 2. deciding which state to pass out and 3. sorting out the recursive calls.

Now, sometimes, when solving DFS problems, we need to
- Pass information back up through the return value, like the max depth.
- Pass information about state into the recursive calls, like max value
Alternatively, we can store state, say a max number, in a global variable.

## Backtracking
Backtracking tacks on some new concepts on top of trees and DFS. In short, the problems don't always give us tree or graph to work with. Sometimes, we have to make the generate the tree or graph as we go.
1. Again, we can make the tree as we go, creating and deleting child nodes as we traverse through.
2. We can drag state around via parameters/returns or with global/struct variables.
3. If we get some crazy 2^N to N! memory usage with backtracking combinatorial problems, we may need to memoize intermediate solutions to cut down on memory usage. For a small to mid N, N! will kill our poor computer. Memoize typically means using a map or similar to store intermediate and final solutions to the combinatorial problems.

## Graphs
Trees are rooted, connected, acyclic, undirected graphs. Trees contain N nodes and N-1 edges and there are only one path between 2 nodes.

This is a tree (and a graph):
```
    1
   / \
  2   3
 /
4
```

This is only a graph due to the cycle and disconnected vertex.
```
    1
   / \
  2 - 3   // Cycle among nodes 1, 2, and 3

4         // Node 4 is disconnected from others
```

Trees and graphs have different terminology
- Vertices are nodes in trees.
- Vertices are connected by edges.
- Vertices connected by an edge are neighbors (children and parents in trees).
- Edges can be directed or undirected. Usually the edges are undirected.
- Paths are sequences of vertices. Cycles start and end at the same vertex.
- A connected graph means every vertex is joined by a path to a vertex; otherwise, the graph is disconnected.

Typically, graphs are stored via adjacency lists or maps. For example, this graph
```
    1
   / \
  2   3
 /
4
```

can be represented in go via a map like

```go
graph := map[int][]int{
  1: {2, 3},
  2: {1, 2, 3},
  3: {1, 2},
  4: {2},
}
```

Note that we don't need to have an adjacency list upfront to solve problems.

So, for BFS and DFS on the graph, we can do the stack/queue dance through it. However, we need a way to dodge any cycles or we'll get stuck in an inf loop or stack overflow. To do this, we can store a map of visited nodes when searching. Like

```go
visited := make(map[int]bool)
visited[1] = true
```

Another clever trick is to wipe out the value in the adjacency graph somehow. Make it negative or something that indicates we visited it to avoid storing a whole other data structure for visited/not-visited stuff.

When deciding between BFS or DFS to explore graphs, choose BFS for shortest distance or graphs of unknown or infinite sizes due to exploring all adjacent neighbors first.

DFS is better at using less memory for wide graphs (graphs with large breadth of factors). Put another way, BFS stores the breadth of the graph as it searches. DFS is also better at finding nodes that are far away such as a maze exit.

## Dynamic programming
A problem can be solved via dynamic programming (DP) if
1. The problem can be divided into sub-problems
2. The sub-problems from bullet 1 overlap

Really, DP == DFS + memoization + pruning. Pruning is important, to save space and reduced wasted calculations.

In DP, the formula used to tabulate Fibonacci numbers is `dp[i] = dp[i - 1] + dp[i - 2]`. These formulas are called the recurrence relation and is critical. Without this, you'll end up flailing around. They also, aren't super obvious. Take finding the longest increasing subsequence (LIS). The recurrence relation from top-down is `lis(i) = max(lis(i-1), lis(i-2) ... lis(0))` but only for `nums at i-1, i-2, ..., 0 < num at i`.

DP problems can be solved in top-down or bottom-up.

### Strategy for solving DP
1. **Identify overlapping subproblems:** Determine if the problem exhibits overlapping subproblems. If not, DP will not work and you'll need to try a different approach.
2. **Handle or preprocess the input (optional):** In some cases, the problem might not provide a straightforward array for DP. For example, see the [perfect squares problems](./problems/dynamic/perfectsquares_test.go) where we need to generate perfect squares first or the [largest divisible subset problem](./problems/dynamic_programming/topdown/largestdivsubset_test.go) where the input needs to be sorted first. Futhermore, you might not realize this straight away which makes the understanding and psuedocoding the problem important.
3. **Define the memo:** Clarify what needs to memoized. In *most* cases, the memo is the same as, or aligns closely with, the problems output.
4. **Initialize the memo with base cases:** We can typically initialize the memo with the base cases. For example, in the [perfect squares problem](./problems/dynamic/perfectsquares_test.go), we can initialize the memo with perfect squares like 1, 4, 9, 25, etc. immediately.
5. **Define the recurrence relation:** Develop the recurrence relation, a formula for transition from one state to the next. For example, clearly state how to transition from `dp[i]` to `dp[i+1]`. Start with transition from the base cases to the next case. Write the recurrence relation in something *you* can understand for coding like:
```
The recurrence relation is:
memo[i] = max(memo[i-1], memo[i-2]... memo[0])+1
  - applies ONLY for each memo[i-1] when nums[i]%nums[i-1]==0.
  - memo[0] is a base case and equals 0
```

### Top down or bottom up DP?
Honestly, it depends on the problem. Generally, from my personal experience from practicing, I prefer DFS + backtracking + memoizing. It lends itself to an iterative problem solving, unlike bottom-up which *feels like* all or nothing. Here's why I like top-down *generally* speaking:
1. Read the problem, do the design work, figure out wtf is going on.
2. Typically, I create a `state` struct to track things we need to expand/solve the problem.
3. Add a queue and load up it's initial state.
4. Write the loop.
5. Pop an item off the queue, check for solutions, and then push on the next states.
6. Once that's solved, add memoizing as an optimization.

### Tricks
- If a sequence length is relatively small, say under 3000 elements or so, it may suggest DP due to N^2 or worse time complexity. So, small inputs, may mean DP because the tests will never finish otherwise.

## Disjoint union set (DSU)

## Intervals
```
start1-----end1   // Interval 1

  start2-----------end2    // Interval 2 that overlaps with 1

                           start3-----end3    // Interval 3 that doesn't overlap with either 1 or 2
```
We can determine if intervals 1 and 2 overlap if `end1 >= start2 && end2 >= start1`. Notice that the formula returns false for intervals 1 and 3.

## Keywords


# Systems design
[Top](#the-study-guide)

1. The network is reliable;
2. Latency is zero;
3. Bandwidth is infinite;
4. The network is secure;
5. Topology doesn't change;
6. There is one administrator;
7. Transport cost is zero;
8. The network is homogeneous.

From [Fallacies of distributed computing](https://en.wikipedia.org/wiki/Fallacies_of_distributed_computing)

## Communication
[Top](#the-study-guide)

### The Internet protocol suite
Here's a model of the abstraction layers of the internet. To be blunt, these layer models are totally bogus; however, they are helpful for conceptualizing the layers of abstraction in Internet communications.

![](./diagrams/out/ip-protocol-suite/ip-protocol-suite.svg)

- The **link layer** operates on local network links like Ethernet or WiFi and provides interfaces to the underlying network hardware. Switches operate at this layer and forward Ethernet packets based on their MAC addresses.
- The **internet layer** routes packets based on their IP address. The IP protocol is core at this layer. Packets are delivered on a best-error can can be dropped, duplicated, corrupted, or arrive out of order. Routers work at this layer, forwarding packets along based on their IP. Note that MAC addresses allow packets to be forwarded from one machine to the next. IP addresses provide the start and end machines.
- The **transport layer** transmits data between two processes. There are many processes on a machine that want to communicate and they do so through port numbers. TCP protocol is used at this layer and attempts make a reliable channel over an unreliable one (lol). Segments are given a number which lets the receiver know if packets are dropped, duplicated, or arriving out of order.
- The **application layer** defines the high level communications like HTTP or DNS. Often, this is the target abstraction for our work.

There are other layers: physical and data link. We'll briefly not them here:
- Physical concerns itself with voltages, pins, cabling, wireless frequencies, etc.
- Data link concerns itself with providing frames for the raw bits and provides some error correction/detection.

Also, TCP is on it's way out and being replaced by [QUIC](https://en.wikipedia.org/wiki/QUIC), aka TCP/2, which multiplexes connections between two endpoints using UDP.

### Border gateway protocol (BGP)
Building and communicating routing tables lies with BGP. BGP maps out routes for forwarding packets along. Note that BGP is concerned about the minimum number of "hops"; it doesn't concern itself with congestion or latency.

ChatGPT says,
> "Absolutely! Let's imagine the internet is like a big neighborhood with houses (which are computer networks) that want to talk to each other. Now, each house has its own unique address, just like how each computer network on the internet has its own address.

Now, Border Gateway Protocol, or BGP, is like the system that helps these houses (networks) know how to talk to each other. Imagine if you want to send a letter to your friend's house in another part of the neighborhood. You need to know the best way to get there, right?

BGP is like the map of the neighborhood for the internet. It helps all the houses (networks) figure out the best paths to reach each other. Just like how you might have different routes to get to different places in your neighborhood, BGP helps networks find the best routes to send information (like emails, pictures, or web pages) to each other.

But here's the tricky part: BGP doesn't know about all the small streets and houses. It mainly knows about the big roads and major intersections (these are like the main connections between different internet providers). So, when a house wants to send something to another house, BGP helps them find the best path using these big roads.

However, just like in a big neighborhood, sometimes things can change. Maybe a road is under construction, or a new road is built. BGP needs to quickly figure out these changes and update the map so that everyone can still send their information along the best paths.

So, in simple terms, BGP is like the map that helps computer networks on the internet find the best paths to talk to each other, just like you finding the best way to get to your friend's house in your neighborhood."

### User datagram protocol (UDP)
An alternative to TCP is UDP, a connection-less protocol that only sends discrete packets of a limited size. It's bare-bones and offers no reliability guarantees. UDP bootstraps other protocols, that want some but not all of TCP's guarantees.

Online multiplayer video games or video streaming may leverage UDP. There's no value in retrying, it would only degrade the user experience.

### Reliable communications
TCP uses segments (not packets) that let receivers detect missing, duplicate, corrupted, and out of order data. Each segment is associated with a timer; if a receiver does not acknowledge the segment, it is resent.

Operating systems manage the **sockets** the store connection states: opening, established, closing. There are more than 3 states, but this keeps it simple.

#### Opening and the TCP handshake
The TCP handshake introduces a full round-trip before any app data is sent. Until a connection is opened the bandwidth is effectively zero. The faster a connection is established, the sooner communication can begin. Ergo, reducing round-trip time by moving servers next to each other reduces the cold start penalty.

![](./diagrams/out/tcp-handshake/tcp-handshake.svg)

Closing the connection, on the other hand, involves multiple round-trips. Additionally, if another connection might occur soon, it doesn't make sense to close the connection so it might stay open.

#### Closing and the TIME_WAIT
Sockets, and the resources they consume, do not immediately close. The enter a waiting state, where late arriving segments are dropped, so that they aren't considered part of a new connection. If you try to open and close many sockets (ports) in a tight loop, you can hit resource exhaustion on ports.

#### Established connections and congestion control
Once communication is started, the sender tries to avoid bombing the receiver with a ton of data. The receiver will shoot back it's buffer size to the sender, so that it doesn't get overwhelmed. TCP is rate-limited just like rate limiting on API key or IP address.

TCP will also try to avoid crushing the underlying network with a ton of traffic. The sender will hold onto a congestion window that'll track the number of segments without acknowledgement. When a segment is acknowledged, the sender can increase the traffic; when not acknowledge, the window is decreased. In fact, bandwidth can be represented by `bandwidth = window_size/round_trip_time`.

### Secure communications
TCP/IP does nothing to secure communications. We need to secure against:
- Spying on data (encryption)
- Unknown or wrong sender/receiver of data (certificates)
- Accidental or malicious changes to data (message auth code via SHA or similar)

Transport layer security (TLS) swoops in, runs on top of TCP, and provides encryption, authentication, and data integrity.

#### Encryption
Encryption means that the data are obfuscated and can only be read by the receivers. When TLS starts, the server and client swap public keys for asymmetric encryption. There's a really great blog on the subject [here](https://blog.cloudflare.com/a-relatively-easy-to-understand-primer-on-elliptic-curve-cryptography/). Once the keys are sent, both sender and receiver use symmetric encryption which is faster and cheaper to minimize overhead. Note that
- The shared keys are regenerated periodically to maintain safety.
- Basically, all traffic should use encryption due to modern CPUs having cryptographic instructions.

Note that the TCP handshake runs first, then followed by the TLS handshake. QUIC aims to speed this along some.
![](./diagrams/out/Tcp-vs-quic-handshake.svg)
(Diagram by By Sedrubal - Own work, CC BY-SA 4.0, https://commons.wikimedia.org/w/index.php?curid=114587250)

#### Authentication and certificates
Even though we can secure communications, we still need to verify that the server is who it claims to be. This is done via certificates which include data about the owner, expiration, public key, and digital signature. The folks that grant certificates are called certificate authorities or CAs.

Certificates need to be present in the client in the client's trusted store. A trust store, also known as the certificate store, is a repository of certificates that the client trusts.
1. The server presents the certificate during TLS/SSL handshake
2. The client verifies the certificates, working up the certificate chain. It'll verify the certificates are valid and not expired.
3. The client will decide if the server is trustworthy. If not, it'll terminate the connection and/or display a warning message.
4. If all is good, the connection may proceed.

Here's an example of a certificate:
```
-----BEGIN CERTIFICATE-----
MIIDdzCCAl+gAw (many more random characters follow)...
-----END CERTIFICATE-----
```

Key values in the certificate are:
- **Version** of the X509 standard being used. Like version 3.
- The unique **serial number** of the certificate.
- The certificate authority that **issued** the certificate, aka, the issuer.
- **Validity period** states when the certificate is valid.
  - **Not before** some date, indicating when the certificate starts being valid
  - **Not after** some date, indicating when the certificate stops being valid
- **Subject** indicates which entity the certificate is issued for.
- The **public key** associated with the certificate and can be exchanged during TLS handshake.
- The **signature algorithm** used by the certificate authority to sign the certificate.
- A digital **signature** used to verify the authenticity of the certificate.

Now, each certificate is chained to an issuing identity, another CA, that granted the certificate. This creates a chain of certificates. The top-level, final certificate is the root CA, like Let's Encrypt.

Here's an example of a website (service), intermediate, and root CA chain.
![](./diagrams/out/certificate-chain/certificate-chain.svg)

We use a chain of certificates because:
1. It creates a hierarchy of trust. Root CAs are typically installed into a client's OS.
2. Allows client to verify the entire chain. If any one certificate is unreliable, the whole chain is considered untrustworthy.
3. We don't need to have 10 billion certificates stored in every client. The keys are distributed around.
4. New intermediate CAs can be added easily, making things scale nicely.

Note that a common mistake is to let the certificates expire, a single point of failure for your whole web stack's security. This'll cause all clients to not trust you. Automating certificate replacement is worth the effort at scale.

#### Data integrity
With encryption we can prevent others from reading the data. With authentication we can prove who we're talking with. However, even with all this, bits could get flipped accidentally or maliciously. A hash-based message authentication code (HMAC) is sent along during TLS. Note that TCP's checksum can fail to detect errors for 1 in 16 million to 1 in 10 billion packets. So, with packets of 1KB in size, it happens once in 16 GB to 10 TB of data transmitted.

### Discovery and DNS
[This video](https://www.youtube.com/watch?v=drWd9HIhJdU) is a solid deep dive into DNS and is worth a watch if you have the time. Even the first 15 minutes is valuable.

IP addresses are cool and all, but we need a way to lookup the IP addresses of servers. There's 2^128 IPv6 and 2^32 IPv4 addresses. Good luck remembering them. Worse yet, IP addresses of servers change all the time. Sometimes the admins need to move requests to a different cluster or spread load among many clusters. DNS helps with this.

Funnily enough, DNS is its own layer 7 protocol.

Anyway, the get-ip-address-from-domain-name is done via Domain Name System (DNS) and DNS resolution.

We're going from `www.amazon.com` to `192.0.2.1` in IPv4 or `2001:d68::1` in IPv6.

You can, of course, slap a port and IP address into the URL bar, but that's a lot of work. Instead we can take the domain name like something.com and resolve it into its IP address 1.2.3.4.

1. If it's a new domain name, the packets are routed to your ISP's DNS resolver. Your browser will cache IP addresses for domain names to save time.
2. The ISP's resolver will iteratively resolve the hostname for clients. It will also cache results.
3. If the ISP's resolver doesn't have it, it'll send a query to the root name server (NS). Root NS map the top-level domains (TLD) of the request, like `.com`. It'll give the address of the TLD server.
4. Once the ISP's resolver has the `.com` part, it'll query the TLD name server with `something.com`. This is the iterative part where the ISP's name resolver will iterate a lot.
5. The TLD name server maps `something.com` to an authoritative name server responsible for the domain.
6. We can return the IP address of `something.com` finally.
7. If there's an optional subdomain, like `definitely.something.com`, then the name server would return the authoritative name server for the subdomain and the process keeps going.

Again, the results are typically cached along the way. Similar to HTTP caching, DNS results are cached with a time to live (TTL) for how long the result is valid for. And, like HTTP caching, we can get some eventual consistency problems that surface. If the cache TTL is too long, then clients will try to connect to the wrong IP address. If the cache TTL is too short, you'll increase DNS load and increase avg response times for clients due to the several round trips DNS takes.

In terms of resiliency, DNS is single point of failure since normal humans simply won't find the IP address of your server and type it in to the browser URL. If, there's a failure and the TTL is stale, we could still try returning the stale result. Continuing to operate or partially operate while a dependency is down is called static stability.

Note that DNS used to be in plaintext, but now it uses TLS. Yay.

### Application programming interfaces
Once we can create semi-reliable and secure connections, we can finally discuss having the client invoke operations on the server. Typically, this is done through application programming interfaces (APIs). APIs can be **direct** or **indirect**.
Term | Definition | Example
--- | --- | ---
Direct | Client communicates directly with a server. | Request-response over HTTP or gRPC. Client sends a `GET /hello`, server responds with `200 OK Hi!`.
Indirect | Client communicates indirectly with a server via a message broker. They don't communicate directly. | A message bus via RabbitMQ or Google's Pub/Sub. Client adds a message, including any optional data, and the server does some processing when it receives the message.

We're going to focus on HTTP for request-response as it's quite popular. gRPC is also popular for inter-service communications via request-response communication.

Anyway, data are serialized via JSON, Protobuf, or similar. XML too, but I suspect it's not as widely supported now. JSON is human readable but slower to serialize and deserialize. Protobuf is not human readable and faster to serialize and deserialize. It's all about tradeoffs.

Now, for clients and servers both, requests and responses can be handled in synchronously, asynchronously, or with a mix of both.

#### RESTful HTTP requests and responses
- Requests are stateless adn contain all necessary information to process. Some requests can be re-executed and create the same response (idempotent). Some, are not.
- Responses are implicitly or explicitly labeled as cacheable or non-cacheable. If cacheable, the client can reuse the response.

#### Synchronous HTTP requests
After the client sends a request, it can simply block until the server responds. While "simple" this typically degrades the user experience quickly and wastes CPU time. We could be doing other stuff while waiting on this request. Note, performing synchronous HTTP requests is typically seen as un-cool.

```go
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func makeSyncHttpRequest(url string) {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	fmt.Println("Sync Response:", string(body))
}

func main() {
	url := "https://jsonplaceholder.typicode.com/posts/1"

	// Make a synchronous HTTP request
	makeSyncHttpRequest(url)

	// Continue with other work
}
```

#### Asynchronous HTTP requests
Alternatively, many languages like TypeScript, C#, and Go hide callbacks via async and await abstractions. We can avoid having the client freeze while waiting for the server to respond.

```go
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

func makeAsyncHttpRequest(url string, wg *sync.WaitGroup) {
	defer wg.Done()

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	fmt.Println("Async Response:", string(body))
}

func main() {
	var wg sync.WaitGroup

	url := "https://jsonplaceholder.typicode.com/posts/1"

	// Make an asynchronous HTTP request using a goroutine
	wg.Add(1)
	go makeAsyncHttpRequest(url, &wg)

	// Do other work concurrently if needed

	// Wait for all goroutines to finish
	wg.Wait()
}
```

#### Synchronous HTTP responses
The server receives a request, does some processing, and returns the entire response immediately. This is typically appropriate when the response is under ~10 MB and can be processed by the server relatively quickly, say in under ~60 seconds. For example, getting a single product's details would typically be straightforward. Listing all products would also be suitable, but the client and server may have to perform paging to get the entire list.

#### Asynchronous HTTP responses
Similar to synchronous, the server receives a request, initiates processing, but only returns partial or intermediate results. Clients will need to check back later to determine when processing is completed and where to get the results. This is appropriate for long running processes greater than ~60s or those with greater than ~10MB of data may benefit from this. Clients may be able to poll the processing status or they could provide a postback for the server to call later.

#### HTTP versions

### Messaging
Messaging is an indirect form of communication where data are passed through a message broker. This nicely decouples the the services. Messaging is typically suitable for processes that take a long time or require large amounts of data.

Pub/Sub, ServiceBus, and RabbitMQ, are all examples of message queue providers. Like HTTP requests, messages typically have header and body data as well as a unique message ID. Also, messages can be stored in JSON, protobuf, and so on. Some brokers even provide support for native objects like Azure. Messages are typically only removed from the broker once the consumer - the service that actually receives and uses the message -  successfully processes the message.

However, unlike request-response via HTTP transactions, message queues are inherently asynchronous. Furthermore, messages can be queued within the broker even if the message consumers are down.

Consumers can choose to process multiple messages at a time, a.k.a. batch processing or message batches. This increases processing latency of individual messages but does improve the applications throughput. This extra latency is typically an acceptable tradeoff and should be considered.

Producers and consumers can be arranged in a few different ways:
- Point-to-point (p2p): A message is delivered to exactly one consumer.
- Publish-subscribe (pubsub): A message is delivered to all consumers.

Message queues are not a silver bullet and there are many issues to consider:
- The broker adds additional latency through additional the queue itself and more communication.
- There is a dependency on the broker itself. It it goes down, so does your system.
- Not all message queues are durable. Non-durable queues will lose all messages if the broker crashes.
- This is a form of sequential consistency. Consumers lag producers.
- There are similar problems to the IP protocol. Messages can be missed, dropped, duplicated, and arrive out of order. Note that some brokers provide message de-duplication.
- We have similar time and timestamp considerations to make, if there are many message producers. Multiple producers will not share the same clock.
- Broker storage is not infinite; there might be a maximum limit to the number of messages a broker can store.
- Consumers might be unable to keep up with producers and the queue which will cause a backlog of messages to accrue. This can be due to not enough consumer replicas, outages, and errors. Poisonous messages, those that repeatedly fail, can waste consumer time and degrade the entire system.
- Completely bad or poisoned messages can be dumped to a dead letter channel for manual inspection and correction by a human. This, of course, takes time. Alternatively, we can decide to drop (delete/ignore) bad messages, but we this means we might miss something important.
- Message queues are partitioned and bring along the same issues that partitioning brings. Messages may not arrive in order. Also, partitions can become "hot" and consumers might not be able to keep up. Message shuffling among the various partitions can reduce throughput.
- *Exactly once processing* does not exist. *At least once and maybe more* can lead to issues. Therefore, we can require messages to be idempotent, so that we can simulate exactly once processing.
- In fact, you may need to choose between *at most once but maybe zero* or *at least once but maybe more* processing.
- Still idempotent messages will not cure everything. Say we create a consumer that sends an email to a customer when it receives a message. However, due to a bug, the consumer crashes *immediately* after successfully sending the email and is unable to acknowledge (ACK) the message. The consumer and broker will simply try again and send another email up to some max retry limit, if there even is a retry limit.

Producer and consumer messaging can be arranged in different styles or configurations, depending on the problem they're solving.
1. One-way
  - Producers can "fire n' forget" messages in a p2p configuration.
2. Request-response
  - Producers can send consumers requests through one message channel and receiver responses from consumers via another response channel. Consumers can tag responses with the request they're responding to and send feedback upstream.
3. Broadcast
  - Messages can be sent to all consumers. After all, there could be many consumers that care about the same message but for different reasons. For example, one consumer could send a confirmation email to the customer and another consumer could do payment processing.

## Scalability
[Top](#the-study-guide)

Basically, we have three ways to scale out applications:
1. Decomposition: Breaking the processing into separate services known as decomposition or functional decomposition.
2. Partitioning: Dividing data into partitions and distributing data among different nodes.
3. Replication: Replicating functionality or data among nodes.

### Replication
As an application scales up, eventually, a single server simply won't be able to keep up. It'll simply fall over.

So, just clone your server's code and you'll be good to go, right? Wrong. Welcome to replicas, consistency models, distributed locking, CRDTs, and so on.

#### Consistency models
Distributed systems are often modeled around consistency models which define how the updates to a distributed system are observed. With different models, you may see data visibility, ordering of operations, performance differences, fault tolerance, ease of implementation, and ease of maintenance. In short, as a system becomes more available, it becomes less consistent.

##### Linearizability
Also known as atomic consistency. This is a specific form of strong consistency which adds real-time ordering constraints to operations.
- At some point, the operation appears to occur instantly, called the linearization point.
- This is a stronger form of consistency than strong (lol, great, not confusing at all) due to the stringent ordering.
- It's also probably the slowest due to the ordering, synchronizing, and coordination involved. Getting all the computers to agree (coordination) takes time.
- The emphasis is all on real-time ordering. When clients check back later, the linearized system updates the data between reads.

ChatGPT says,
> "Imagine you and your friend Alice have magical walkie-talkies. With linearizability, when you send a message to Alice, it's like she receives it instantly, and her response also comes back to you immediately. It's as if the messages are magically happening at the same time."

#### Strong consistency
Strong consistency ensures all nodes have the same view of data at any given time.
- All operations appear to be synchronous and read operations appear to return the most recent write.
- Typically uses distributed locks to agree on the order of operations.

ChatGPT says,
> "Strong consistency is like having a rule that everyone, including friends in different rooms, always has the same information at the same time. If you say something to Alice, all your friends instantly know about it. It's like the information spreads quickly to everyone."

##### Sequential consistency
Sequential consistency ensures operations occur in the same order for observers but does not make any real-time guarantees about when an operation's side effect becomes visible. There's a boost in performance but we drop some consistency on the ground. A read from Server A may appear different than Server B, but the operations are sequential so Server A and B will eventually converge and agree. In other words, the replicas are diverging on their view of the world.

A producer/consumer system, synchronized with a message queue, is an example of a sequential consistency. The consumer lags behind the producer.

ChatGPT says,
> "Now, let's say you and Alice have regular walkie-talkies, but you both agree to take turns talking. So, the messages go back and forth in the order you send them. It's like having a clear order for your conversation."

##### Causal consistency
Causal consistency relaxes some guarantees of strong in favor of speed. Causal guarantees that causally related operations are in a consistent order and preserves causality. Before we continue, we need to discuss the CALM theorem quickly. So, stay CALM (Get it? Stay calm!!? About the theorem? Funny, right?)

###### The CALM theorem
The Consistency As Logical Monotonicity (CALM) theorem uses logic to reason about distributed systems and introduces the idea of monotonicity in the context of logic. Furthermore, it helps give us a framework to determine if we can move coordination off the critical path. CALM tells us we can get to coordination-free distributed implementations only if the system is monotonic.

Logically monotonic means that the output only further refines the input and there's no taking back any prior input.

Very important note: The consistency in CALM is not the same as the consistency in CAP.

Variable assignment is non-monotonic. When you assign something to a variable, the previous value is gone forever. Take a counter value that does

write(1), write(2), write(3) => 3
but then if they show up out of order:
write(1), write(3), write(2) => 2

We end up with the wrong value.

In contrast, incrementing allows us to reorder in any way and still get the correct output:
increment(1), increment(1), increment(1) => 3

Check out [Keeping Calm PDF](https://arxiv.org/pdf/1901.01930.pdf) for a much better and detailed description.

##### Back to causal consistency
Back to causal consistency. Causal maintains happened-before order (the causal order) among operations. This makes causal attractive for many applications because:
- It's consistent "enough" and easier to work on the eventual consistency.
- Allows building a system that's available and partition tolerant.

This requires that nodes *agree* on the causally related operations but may *disagree* on the order of unrelated ones. Put another way, the nodes preserve the logical order.

Causal systems are typically backed by conflict-free replicated data types (CRDTs), such as
- Last writer wins (LWW) - Values are associated with a logical timestamp/version. When the value is broadcasted, nodes only keep the greatest timestamp. Conflicts due to concurrent updates are usually resolved by taking the greater timestamp, but this might not always make sense.
- Multi-value (MV) - Store the operations in a log of operations that all nodes share. New values are inserted into the MV register. Systems will need to share out their MVs.

ChatGPT says,
> "With causal consistency, you and your friends agree on some logical order for the messages. If you tell something important to Alice and then mention it to Bob, everyone knows that Alice got the message first. It's about maintaining the cause-and-effect relationship."

##### Eventual consistency
Eventual consistency relaxes guarantees of strong and sequential. Given enough time, all nodes will converge to the same result. Note that, during updates or partitions, nodes may have different values. For example, reading from Server A and B may yield a stale, earlier result which is very confusing.

Imagine uploading an image to a social network and add it to gallery. Except, when you try to view the gallery, you get a 404; strangely, you already received a message upload success message! Very odd to an observer, but very real in distributed world.

Regardless, eventual can be appropriate tradeoff for certain systems. Slightly different results are perfectly acceptable to achieve higher speeds such as returning the number of users on a website. The number of active users can be stale, and it's typically not a big deal. Note that eventual consistency can be a maintenance burden due to subtle, unexpected bugs.

ChatGPT says,
> "Now, suppose your friends have regular walkie-talkies, and sometimes messages take a while to reach everyone due to delays. Eventual consistency says that, given enough time without any more messages, eventually, everyone will have the same information. It's about waiting until things settle down."

### CAP and PACELC theorems
["CAP Theorem: You don't need CP, you don't want AP, and you can't have CA"](https://www.youtube.com/watch?v=hUd_9FENShA)

Engineering is, in part, about tradeoffs. Distributed systems are no different. CAP theorem, can be summarized as "strong consistency, availability, and partition tolerance: pick two of three." Except...
- Network partitions are simply unavoidable, so you really just get to pick between availability and consistency
- But, also, CAP defines available as *eventually* getting a response but we know that perfect availability is impossible
- And also a slow response is as bad as not receiving one at all.
- Network partitions are rare in a data center. It can certainly happen though.

So, while helpful, CAP is limited in its practical application. This is, again, about tradeoffs. The PACELC theorem, an extension of CAP, expresses this as, when a network is partitioned (P), choose between availability (A) and consistency (C). Else, when operating normally (E), choose between latency (L) and consistency (C). We see that this is not some binary choice between AP and CP, but rather a spectrum of tradeoffs between the various consistency models and latency. Indeed, some systems like [Azure's Cosmos DB](https://learn.microsoft.com/en-us/azure/cosmos-db/consistency-levels) allow you to choose the consistency model you want which is neat.

### HTTP caching
Static resources like CSS, JS, JSON, etc. don't typically change super often. The client's browser can cache the results which boosts response time, lowers server load, etc. In terms of APIs, only reads (GET and HEAD) are cacheable. The other verbs modify data.

The first time a client requests static resource, we can reply with the resource plus add a Cache-Control HTTP header with how long to keep the resource (time to live or TTL) and the ETag (resource version).

![](./diagrams/out/http-cache-miss/http-cache-miss.svg)

If the cache resource is stale (age > max-age), we'll request it from the server again. Otherwise, the next time the same resource is requested, we'll get it from the cache.

![](./diagrams/out/http-cache-stale/http-cache-stale.svg)

This is great because we simply added a HTTP header and reduced server load and boosted response time for the client. Unfortunately, we added eventual consistency, especially if the caching is poorly managed.

For example, in the worst cases, say the dev team changed an endpoint with a breaking change and updated the JavaScript that hits this endpoint. If the client is still using an old JavaScript file due to improper headers/caching, they'll get an error when trying to use the endpoint with the wrong JavaScript code. This will require clients to perform a hard reload, wait for the resource to expire, or simply try again and hope that the consistency is there.

Regardless, this is typically a good price to pay by being careful when using the cache, using cache-busting techniques, use CSS/JS/whatever hash codes or versions like `<link rel="stylesheet" type="text/css" href="styles.css?hash=abcd1234">` or `<link rel="stylesheet" type="text/css" href="styles.css?v=1">`. Alternatively, we can force clients to retrieve static resources from different URLs.

Note that HTTP caching treats reads separate from writes like CQRS.

### Partitioning
As our data storage and access grows, one machine will simply be unable to handle everything by itself. To continue to scale, we can partition (divide up) or federate data among other nodes. With more nodes, we will be able to spread the load through an increased system capacity. Typically, this partitioning, or mapping data to an appropriate *shard*, is done through a coordination service like Apache Zookeeper which will be able to map data to nodes.

Partitioning has costs, however.
- We now have a new dependency on a coordination service with all the fun that brings.
- Adding or removing partitions could require shifting a lot of data. We want to avoid shifting data to redistribute load where possible because it's time consuming and disruptive.
- Performing aggregate calculations (sum all order totals, count all customers, etc.) requires querying data among many shards.
- "Hot" or frequently accessed partitions may limit the systems ability to scale.

There's basically two broad types of partitioning: ranged and hash.

#### Ranged partitioning
Ranged partitioning means dividing the data by time, names, or some other value. Ranged partitioning can be tricky, however, as we need to distribute our keys well to keep up with demand and load.

For example, say we created shards to hold customers, one for each letter of the alphabet. We might expect that the Z shard won't get nearly as much traffic as S. In fact, we might get so many S customers that we need to split (rebalance) the S shard into two just to keep up! Similar issues can arise from using time to divide data.

There are some ways to fight against this. We can allocate bonus shards ahead of time, more than we need, to keep up with demand. However, the shards may still have hotspots or be under or over utilized still. Another option is *dynamic sharding* where the system will automagically split and merge partitions based on size and access. If a shard is too large or accessed very frequently, we can split the shard to spread load and vice versa.

Note that we can attempt to prefix IDs to get a better distribution, but this can be tricky and complex.

#### Hash partitioning
Another option is to use a hashed value for distributing requests or data. Hash values themselves can end up being more random than customer names and would potentially distribute data or requests more evenly.

Note that we prefer to use [consistent hashing](https://en.wikipedia.org/wiki/Consistent_hashing) to distribute load instead of naively using modulo or similar. Adding or removing partitions using modulo would be quite disruptive, causing most keys and associated data to be re-organized into different nodes. Consistent hashing minimizes this disruption.

Note that we lose convenient sorting that ranged hashing provides when using hash partitioning. This can be somewhat relieved by using secondary keys for sorting on an individual partition.

### Load balancing
[Top](#the-study-guide)

If we have multiple application servers, we can route requests to all them. Now, this is actually somewhat tricky to do well.

Servers, behind the load balancers (LBs) will scale up and down transparently to consumers. If LBs detect faulty servers, it can remove them from the server pool.

#### Load balancing strategies
There are tons of strategies for shelling out load to server replicas, but it's tricky. For example, we could:
- Split requests evenly -> Except some servers might have more hardware resources than others. Not all servers are exactly the same!
- We could expose a health endpoint and read the CPU load -> Except a CPU that just comes online could get bombed by requests. Then, when it's overloaded, it'll go back down to zero only to be hammered again!

Name | Description
--- | ---
Round robin | Sling requests to servers in order, one at a time. LB1 -> LB2 -> LB3 repeat.
Weighted round robin | Some servers are more powerful than others. We can give more requests to the stronger servers. Say LB2 is stronger, we could do LB1 -> LB2 -> LB2 -> LB3.
Least connections | Sometimes, a server might draw the short straw and get overwhelmed with a bunch of long running requests. This is a dynamic LB strategy where we give requests to servers with the least active of connections.
Weighted least connections | Similar to least connections. Servers are dynamically scored based on the servers resources.
Resource based (adaptive) | The LB can get health checks from the servers and their overall health. Scores and requests can be dynamically assigned based on the servers returned status check.
Resource based, (Software defined network adaptive) | Using OSI layers 2, 3, 4, and 7, the LB can make optimized traffic decisions based on network congestion or similar.
Fixed weighting | Servers are assigned a scored (weighted). The highest score will get the requests and then failover to the next highest score. This one is useful where a single server is expected to handle all the traffic but be assigned a failover server.
Weighted response time | The LB can use a server's response time to decide where to send traffic. The fastest server wins the next request.
Source IP hash | The source and destination IPs are hashed and used to direct requests. This one is useful if all requests need to go to the same server.
URL hash | Similar to source IP hash except the requested URL is used instead and helps ensure requests are sent to the same servers.

So, lots of strategies we can try. Supposedly, randomly picking two servers and assigning the request to the one with lower load works well ([here](https://brooker.co.za/blog/2012/01/17/two-random.html)).

#### Layer 4 and 7 load balancing
Load balancers can make routing decisions based on information extracted from bytes in the TCP stream (transport layer 4) or from the HTTP header, cookies, or resource type (application layer 7). Note that, layer 4 vs 7 is talk here is intended to be more useful for learning than 100% accurate. It's more complicated than pick layer 4 vs 7 - done.

- Layer 4 requires less computing resources as there's less protocols to process, you maybe don't have to wait for the entire packet, etc.
- Layer 7 load balancing is often desired for sheer flexibility since, at layer 7, we have a full understanding of the request being made.

Typically, besides higher flexibility, hardware nowadays is powerful enough and to make layer 7 LBs outweigh layer 4 LBs.

#### Load balancing failover
A single LB presents a single point of failure. If the single load balancer goes down, then the whole system behind it goes down as well. To maintain availability, we can have multiple load balancers in active-active or active-passive configurations.
- Active-active = There are multiple, ready LBs in use. If one active goes down, then another is ready to take the additional load.
- Active-passive = One LB is active, another LB is in standby. The standby LB is ready to hop in if the active one goes down.

#### Load balancer uptime
The more nines you need, the more LBs you'll need to supply. Now, this isn't perfect, but the availability is roughly `1-(Nines^LB_Count)`. So, two servers with 99% availability becomes `1-(0.01*0.01)=0.9999`. Now, this doesn't really account for chained failures 

#### Service discovery
Again, servers can come up and down frequently. A coordination service like Apache Zookeeper or Kubernetes can help with this.

References:
- [Load balancing algorithms](https://kemptechnologies.com/load-balancer/load-balancing-algorithms-techniques)

### Forward and reverse proxies
ChatGPT says,
> "Let's imagine you want to get a toy from a toy store, but instead of going there yourself, you send your friend to bring it back for you. In this scenario, your friend is acting like a proxy.

Now, there are two types of proxies: forward proxies and reverse proxies.

Forward Proxy (like your friend going to the store): Imagine you want to visit different toy stores, but you don't want the stores to know it's you every time. So, you send your friend to the stores, and your friend brings the toys back to you. The stores only see your friend, not you. This is like a forward proxy; it helps you access different things on the internet without directly interacting with them.

Reverse Proxy (like a helper at the toy store): Now, let's say you go to a huge toy store, and there's a helper at the entrance. Instead of going inside to get each toy yourself, you tell the helper what you want, and the helper gets the toys for you. The helper is like a reverse proxy because it helps you get what you need from the store without you having to go all the way in. In the internet world, a reverse proxy helps websites give you the things you want (like web pages or pictures) without you directly talking to the main server.

In summary, proxies are like helpers or friends that help you get things from different places, and reverse proxies specifically help websites give you what you want without you having to go to the main server every time."

### Content delivery networks (CDN)
CDNs are collections of caching servers. When clients request certain resources, like images or video, they can be served from physically close CDNs that cache the content. If the CDN doesn't have the resource, it'll request the content from the "origin server" (read: your application server) from the original URL. The resource is stashed locally and finally served to the client.

#### CDN networks
The main power of CDNs isn't the caching, it's actually the CDN network itself. BGP doesn't concern itself with latencies or congestion. CDNs can exploit and optimize various techniques to increase bandwidth. They have persistent connections and optimal TCP window sizes to maximize bandwidth.

Furthermore, CDN are optimized in terms of location and placement. There's a global DNS load balancing, an extension of DNS, that considers client location via IP address and returns a list of the closest clusters, network congestion, and cluster health. Next, CDNs are placed at internet exchange points, points where ISPs connect together.

#### CDN caching
The top level servers in CDN are the edge servers/clusters. If content is unavailable at the edge, the edge will get the content from the origin server while leveraging the overlay network.

Now, imagine there's hundreds of edge CDN servers. They could end up overwhelm origin servers (read: your servers) due to a low cache hit ratio. More edge servers, closer physically, but more misses. Less servers, further physically, but more content hits. To alleviate this pressure on the origin servers, CDNs have intermediate cache servers with a larger amount of content.

Note that CDN content is partitioned among many servers because no one computer can handle all the content. Reminder that partitioning is a core scalability pattern.

#### Push and pull CDNs
Before clients can get resources from a CDN, the content needs to be delivered to the CDN somehow. And there are tradeoffs to consider.
- Resources can be **pushed** to the CDN. That is, software engineers push assets up and then those assets are propagated through the other CDN nodes. Push is flexible and can be accurate, but it requires engineers to put in maintenance effort.
- **Pull** CDNs will fetch assets based on request. If the CDN doesn't have the asset, it'll be retrieved from the origin server. Pull CDNs relax the maintenance burden and save space as assets are only uploaded on request. Unfortunately, pull CDN disadvantage comes in the form of duplicate requests to the origin server. If an asset isn't cached and CDNs receive many requests, they can send duplicate requests to the origin server for content. Also, first time visitors will have a slow experience. One could offset this by manually requesting pages as soon as they are available, however.

### Blob storage
Now, even with using a CDN, our file storage and access speed is limited. We can circumvent this using blob storage which store files in a scalable, available, and durable way and can be configured to be private or public.

Modern blob storage has tons of great abstractions like files, queueus, and tables. Amazon's S3 even allows you to query file data with SQL via S3 Select which is super cool. We'll only focus on file storage here.

#### Blob vs File
Blobs themselves are considered to be more generic than a file. Blobs can include really any binary data include images, audio files, video files, DB records, executables, cryptographic data, etc.

##### Traditional hierarchical file-based systems
These traditional file storage systems include File Allocation Tables (FAT) and New Technology File System (NTFS). Broadly speaking, they store files like:
```
C:\
├── Program Files
│   ├── Microsoft Office
│   │   ├── Word.exe
│   │   ├── Excel.exe
│   │   └── PowerPoint.exe
│   └── Adobe
│       ├── Photoshop.exe
│       └── Illustrator.exe
├── Users
│   ├── User1
│   │   ├── Documents
│   │   │   └── Report.docx
│   │   └── Pictures
│   │       └── Photo1.jpg
│   └── User2
│       ├── Documents
│       └── Pictures
└── Windows
```
- **Organization:** They organize data in a hierarchical or tree-like structure with directories (folders) and files.
- **Metadata:** Each file and folder typically has associated metadata, such as file attributes, permissions, and timestamps.
- **Access Methods:** Files are accessed using a file path within the hierarchical structure, and traditional file operations like reading, writing, and deleting are common.
- **Structured Storage:** These systems provide a structured way to store various types of data, including text files, executables, and multimedia files, within a file hierarchy.
- **File Naming:** Files are often identified by names within their respective directories.

##### Blob storage
Binary Large OBject (BLOB) are used to efficiently store and retrieve data. Blob's are considered more general than files and do not have the strict adherence to hierarchy. It also introduces different access patterns and using DNS to find files.

- **Binary Large Objects (Blobs):** Blob storage is designed specifically for handling binary large objects (BLOBs), which can include files but also any other binary data.
- **Flat Namespace:** Blob storage often provides a flat namespace, meaning it doesn't enforce a hierarchical structure on the stored data. Instead, it allows flexibility in organizing data.
- **Unstructured Data:** Blob storage is well-suited for unstructured data and is not limited to the traditional file and folder structure.
- **Scalability:** Blob storage is highly scalable and can handle massive amounts of data, making it suitable for cloud-based applications.
- **Different Access Patterns:** While traditional file systems are designed for file-based access patterns, blob storage is often used in scenarios where there are diverse access patterns and a need for versatile data storage.

#### Blob storage
In Azure Storage (AS), there are many *storage clusters* in different locations. There's lots of redundancy and backup power to achieve a high degree of nines.

A centralized *location service* creates new accounts, allocates accounts to clusters, and moves accounts to redistribute load. Customers can pick which cluster to allocate and a new DNS record will be added that maps an account name to a cluster.

Each storage cluster is composed of front-end, partition, and streaming layer abstractions, each with there own manager or control plane.
- **Front-end** A reverse-proxy that authenticates and routes requests to the appropriate partition server.
- **Partition** Our high level file/blob operations are translated here. It holds indices with metadata that point to an account's stream of data. Per it's name, it handles managing partitions and dynamically splitting them when they're too large. The partition layer also replicates accounts around for load balancing and disaster recovery.
- **Stream** This is a distributed append-only file system. There's no changing an existing blob, only replacing it. *Extents* store chunks of binary data that make up the blob. *Streams* then consist of many extents. This layer abstracts away the underlying storage concerns, uses streaming data, and manages the extents. It also allocates new extents or tries to fix unavailable or faulty extents. Underneath the hood, this layer is using chain replication for the extents which will be passed to the low-level storage servers. This chain replication provides strong consistency guarantees.

#### Blob access
In terms of accessing blobs on AS, the account and file name in AS can be used to create a unique URLs and can be looked up via DNS. For example, `https://my_account_name.blob.core.windows.net/my_file_name`. Furthermore, we can adjust the access privileges on blob data. We could make a blob completely public facing or require certain IAM settings.

### Microservices
Monolithic architectures, a single large code base with all the business concerns in one code base, are often great for certain companies, teams, and projects. It's easy to coordinate changes, run the application, handle inter-process communication well, and lower complexity. At least initially. It's often recommended to start with a monolith and design with clear domains/boundaries.

Monoliths struggle to scale in terms of service demands and programmer efficiency. Monoliths can become unwieldy if they grow overly large, and this is especially true if they evolve poorly. New features, understanding code, adding features, and bug fixing can grow to be hard. Suddenly, reverting a change becomes very hard and time consuming.

Microservices, on the other hand, can alleviate these pain points while introducing some new ones. Note that there's really nothing micro about microservices and the name is a misnomer. Microservices are independently managed and deployed services that communicate via APIs or message queues. Typically, one team owns the service and can dictate its code, releases, testing, availability, etc. It's also easier on new hires, too.

Now, microservices add a great deal of complexity overall.
- Without conventions and standardization, each team has as choose-your-own-adventure with their microservices. Nothing is stopping one team from using Java while others use Python; one team could test with xUnit and another could test with NUnit. There's also many cross-cutting concerns such as logging.
- Half the problems with distributed systems come from communications. This introduces more new and interesting communication problems.
- Poorly handled API evolution, shared databases, overly communicative, or badly laid out domains or responsibilities can all introduce a horrible *distributed monolith* anti-pattern. A distributed monolith has all the cons of a monolith but bonus horrors of a distributed system. Agility is reduced as releases must be coordinated, scalability is limited due to the coupling, finding errors is hard, and correlated errors become more frequent. To avoid it, establish clear service responsibilities, avoid shared databases, use async communications such as message queues, implement API versioning, and monitor monitor monitor.
- Resource is more provisioning with microservices. Again, this needs to be standardized or each team will end up with something different.
- Debugging and testing can be more challenging. Certain issues will only surface through service interacting together.
- Microservices often have separate datastores and this creates eventual consistency. There's nothing stopping services from changing data and disagreeing about what is true.

Again, if a monolith is well designed with proper boundaries, we can pull microservices out when needed.

#### API gateway
We want to minimize the number of endpoints that clients need to hit. We don't want consumers to have to worry about addressing our 50 different microservices. We can introduce an API gateway, an abstraction and reverse proxy, that hides the microservice architecture from clients. We can route requests via the reverse proxy, change internal endpoints without breaking clients, stitch together data from multiple services, and handle data inconsistencies. An API gateway is especially beneficial for public APIs backed by microservices. The gateway can also provide rate-limiting, caching, and authentication/authorization. Again, we need to be mindful of tradeoffs with a gateway: scaling, coupling, and microservice API changes. In terms of implementing an API gateway, you can
- Create your own
- Start with a reverse proxy like Nginx (pronounced engine-X)
- Use a managed service like Google's [Apigee](https://cloud.google.com/apigee?hl=en)

##### GraphQL
The gateway can also perform translation. For example, HTTP requests can be converted to a gRPC call or it could return less data for the mobile app. Graph APIs can help with this. [GraphQL](https://graphql.org/) (graph-ickle), for example, is a popular server-side library in this space that allows clients to simply request what they need. The idea is creating a backend that serves frontend needs (backend-for-frontends) where the frontend can query only for what it needs. Theoretically, if designed well, GraphQL should increase frontend flexibility for changes, decreases load through getting multiple resources in a single request, and reduce querying superfluous data. Of course, GraphQL has tradeoffs. GraphQL is complicated to get started and adds yet another dependency. We need to be cognizant of GraphQL security, query depths and limits, and so on else we run the risk of overloading our servers or exposing data to malicious actors. On the frontend, [Relay](https://relay.dev/) is a React client library and can make getting data from the GraphQL backend easier.

## Reliability
[Top](#the-study-guide)

At scale, anything that can go wrong usually does. Computer parts break, network errors galore, programmer errors, and so on.

### Common failures
There's several, common failures that can occur:
1. **Hardware problems** Like CPUs, memory, or HDDs dying.
2. **Bad error handling** Programmers mishandle certain exceptions or don't handle them at all. 
3. **Configuration changes** Updating certain configuration values can be just as deadly as a bad coding mistake. Configurations should be managed. For example, changing feature flags or YAML configuration files should be done carefully.
4. **Single points of failure (SPOF)** Any SPOF isn't super great. A single DNS server, LB, replica, availability zone, etc. can all bring down your application. Humans, those with certain positions or privileges, can bring down entire systems. For example, a coder can push a single bug and bring down systems.
5. **Network issues** There are just tons of reasons for slow or missing responses. Even semi-slow responses can lead to *gray failures* of sorts.
6. **Memory or resource leaks** Memory leaks can cause a system to slow down over time. As RAM starts to get choked out, the OS will move RAM to HDD disk aggressively and/or garbage collectors will get more aggressive. This all eats up CPU time and degrades performance. Ports and threads can also leak.
7. **Load pressure** Sudden spikes in load or load pressure can cause system outages for pieces that doesn't scale. For example, in a prior gig, they put up a product for free. The sudden influx of customers caused the system to go down. Alteratively, Christmas season shoppers can suddenly overwhelm a system or similar.
8. **Cascading failures** Say a LB is hands requests to two downstream servers A and B. Everything is going smoothly but then server B suddenly dies. LB says, "No problem, I'll send all the requests to A now." Suddenly, server A can't keep up and dies. A cascading failure.

### Risk
So, risks and failures are not the same thing. Risk is a failure that has a percent chance of occurring. You have two choices with risk: mitigation and ignoring. For example, we can simply ignore the risk of a issuing a single HTTP request or we can try to mitigate issues by retrying.

Note that we can try to rank certain risks, issues, gaps, etc. and determine mitigation or ignoring steps. For example, low risk and impact could maybe be ignored; alternatively, high risk and impact should would be worth mitigating. One could even factor the cost of mitigation into the chart.

### Redundancy
Replicas are one defense against failures, but again this has downsides in terms of complexity.

Unfortunately, redundancy doesn't solve all our risks. *Correlated errors*, failures that can hit many servers at once, will still hurt. For example, redundancy won't help if a software bug deployed to all replicas. Likewise, a total data center outage will bring down all the replicas too. To help, cloud providers replicate stacks to multiple data centers. Multiple data centers in a region make an *availability zone* (AZ). Data centers are close enough to have low latency but far enough away to reduce disaster chance.

### Fault isolation
We can also isolate problems much like the bulkheads on a ship. Without a bulkhead, a single leak can sink the entire ship. However, by incorporating bulkheads, only a specific part of the ship is affected during a leak, as we've compartmentalized it with walls. The ship is degraded, but it doesn't succumb entirely. Distributed systems encounter similar vulnerabilities, akin to "leaks" called *poison pills*. A single user might send malicious requests capable of bringing down the entire system, or normal bots and crawlers, mimicking human behavior with rapid request capabilities, can overwhelm servers.

Similar to ship leaks, challenges arising from poison pills can be mitigated by using fault isolation. For instance, consider having 6 replicas and an antagonistic computer bot attempting to disrupt your service. Each request from the bot takes down every replica it can reach. With 6 requests, the entire system could be brought down.

To address this, let's establish 3 pairs of servers to introduce redundancy and consistently direct users to the same server pair. By confining the bot to the same server pair, even if 2 out of 3 server pairs remain operational, only approximately 33% of users will be affected. While not ideal, this approach is superior to the bot impacting all servers and causing a complete system outage.

#### Shuffle sharding
Now, we can enhance our system further by employing shuffle sharding. The goal is to reduce the number of users that share replicas. Instead of consistently directing the same users to specific replicas, we can dynamically reassign users to minimize shared users between pairs. For instance, users A and B may initially share server 1, but with redundancy measures, we can reposition user A to server 2 and user B to server 3. Consequently, if user A causes a server to go down, user B can still receive service requests, ensuring uninterrupted functionality.

There will still be some user overlap, but it will be less likely. With this pattern, we can increase our bulkheads to 15! The combination formula is `n!/(r!(n-r)!) = 6!/(2!4!) = 15`. This tells the number of unique subgroups we can have.

#### Cellular architecture
We can combine all dependencies together into a cell: LBs, storage, servers, etc. and scale out cells on requests. Requests can be shared with cells via a gateway service. A benefit is that we can learn the maximum limits of each cell and scale out accordingly. We can totally understand the performance of a cell and how it handles load.

#### Pool architecture
An alternative to cellular architecture is pool architecture. Similar to cellular pattern, we can spin up and scale out services as needed but we can do this on a per customer basis. Small clients could share resources while jumbo customers could have their own separately scaled out resources. This way, a high traffic customer won't resource the smaller customers. See [this article](https://medium.com/@raphael.moutard/forget-your-microservices-the-unparalleled-benefits-of-pool-architecture-63b462989856) for more details.

## Monitoring
Monitoring is mainly used to alert teams to failures and monitor system health via dashboards: pub/sub queue sizes, dead letter queue sizes, CPU utilization, RAM utilization, SLO/SLI values, replica counts, request counts, throughput, replica restarts, and so on.

A metric is a time series of raw measurements or samples of resource usage or behavior represented by a timestamp and floating point number. Metrics are typically tagged with many labels, a set of key-value pairs, that make filtering easier. We can indicate the service, region, node, cluster, app name, etc. Typically metrics are stored in a time series database (TSDB) like [Prometheus](https://prometheus.io/) that often use different [data models](https://prometheus.io/docs/concepts/data_model/).

Minimally, we want services to output their load like request throughput and their internal state such as memory usage.

### Pre-aggregation
Typically, the telemetry services, like Azure Application Insights telemetry, will batch and send out telemetry data. Application Insights can store data locally in the event a computer shuts down unexpectedly and can be used to correlate data from the frontend and backend of systems. Client libraries will typically perform some pre-aggregation themselves before transmitting the data.

Storing and querying the time series data can get expensive. Typically, the telemetry services pre-aggregate data into different buckets: 1 min, 5 minutes, 1 hour, 3 hours, 1 day, etc. and include summary statistics like sums, averages, or percentiles. AWS's CloudWatch pre-aggregates data on ingestion.

### Service level terminology
| Abbreviation | Term | Definition | Example
| ---- | ---------- | ------- | ---- |
| SLI | Service Level Indicator | A carefully defined quantitative aspect of the level of service that is being provided. | Request latency in milliseconds, error rates, system throughput, etc. Queries per second (QPS) is not a good SLI since we don't control how many queries from users we get. This doesn't mean that QPS isn't worth knowing, however.
| SLO | Service Level Objective | A target or range of acceptable SLI values. | Availability >= 99.99%, avg. response times <= 400ms, P99 response times <= 400ms, etc. Again, queries per second (QPS), on the other hand, isn't under our control and doesn't make for a good SLI/SLO.
| SLA | Service Level Agreement | XXXXXXXXXXX

### Service level indicators
Metric values can vary, sometimes drastically. Just because one request in a million took 7 seconds doesn't mean that an engineer should be woken up at 3am. Service level indicators (SLIs) measure one level of service; typically an SLI is an aggregated error rate, throughput, response time, or similar. Typically, an SLI is calculated via good count / total count. Some common SLIs include:
- Availability = successful requests / all requests.
- Response time = fraction of requests faster than some threshold

We want to pick SLIs that *best measure the user experience*. So, for response times, measuring the response time client-side would be best provided this is not cost and effort prohibitive.

Next, we typically analyze response times via summary statistics like average or percentiles. Again, 1 request in 1000 that's 10 minutes isn't worth waking an engineer. However, outliers can greatly skew our average; for example, a single 10 minute response out of 100. We can filter out *long-tails*, values far from the central distribution, that skew averages by using percentiles. For example, 99th or 99.9th percentiles, typically represented like P99. **This does not mean we should ignore long-tail values. Your most important customers could be having a poor experience.** We want to keep our long-tails in check. They can consume tons of resources and improving them may improve resiliency and average-case scenarios.

# Napkin math
[Top](#the-study-guide)

"Estimation is a vital skill for an engineer. It’s something you can get better at by practicing." - [Roberto Vitillo](https://robertovitillo.com/back-of-the-envelope-estimation-hacks/)

Also known as back of the envelope calculations. These numbers are **heavily** rounded for memorizing. We want to be in the ballpark for creating useful mental models and to develop a gut feel for numbers being discussed to streamline conversations and make sure we've got a top-notch mental model of what's being talked about. If you need accurate numbers, then do your Ti-83 math instead.

## Costs
Cloud costs are important. It's good to have a sense of how much cloud computing costs will run. If we're using 400, four core CPUs per month, we want to know what that'll cost.

Name | $/time/quantity | Description
--- | --- | ---
Core | $10/mo/core | Cost of cloud computing cores per month (typically have many)
SSD | ¢10/mo/GB | Cost of SSD storage per month GB
HDD | ¢1/mo/GB | Cost of a HDD storage per month per GB
CDN | ¢1/mo/GB | Cost in AWS or GCP for CDN storage per month per GB
Network | ¢1/mo/GB | Networking utilization costs in AWS or GCP per month per GB

## Uptime in nines
Obviously, zero downtime in systems is ideal, but this ain't realistic. We want to be as close as financially reasonable to 100% up time. We talk about this using the "number of nines". The nines themselves don't really tell you the actual amount of downtime allowed. For example, what's 4 nines (99.99%) of 365 days? Why, it's 0.0365 days obviously! Yeah, no, still not helpful. Memorize the times.

| No. of Nines | %        | Daily downtime | Annual downtime |
|--------------|----------|----------------|-----------------|
| 1 nine       | 90%      | 150min         | 36days          |
| 1.5 nines    | 95%      | 75min          | 18days          |
| 2 nines      | 99%      | 15min          | 4days           |
| 3 nines      | 99.9%    | 2min           | 9hrs            |
| 4 nines      | 99.99%   | 10s            | 1hr             |
| 5 nines      | 99.999%  | 1s             | 5min            |
| 6 nines      | 99.9999% | 100ms          | 30sec           |

Also note that downtime among serial systems is typically additive. Say we've get two services named A and B where A depends on B. If they both have 5 nines of uptime, then they could both be down for 1s per day but there's no guarantee these times overlap. So, the worst case is going to be 2 services * 1s of downtime = 2s of downtime.

The more nines, the more expensive it is. For example, in terms of load balancers (LB), you may need to consider LBs+1, LBs+2, or 2*LBs to get your desire uptime. Now, this formula isn't perfect, but the availability is roughly `1 - ((1 - server_nines_as_decimal)^server_Count)`. So, two servers with 99% availability becomes `1-(0.01*0.01)=0.9999`. Now, again, this doesn't account for chained failures among many replicas.

## Little's Law
Lot's of things can be queues: message brokers like Google's Pub/Sub or even a web service. Little's law can help us relate the the average of new items arriving into a queue, the average time an item spends in a queue, and the number of items in the queue.
`avg_items_in_queue = avg_rate_of_new_items_coming_into_the_queue * time_items_stay_in_queue` or `q = r * t`. Say we have a service that takes an average of a 100ms to process a request and we are getting 2 million requests per second (rps). The equation via Little's Law is `r * t = 2_000_000 request/S * 0.1 S/request = 200_000 requests total = q`. We can take this a step further and calculate the number of computers we need. Assume that the requests are data heavy, and that we'll need an entire CPU thread for each request. If we have 200K requests per second and we have 8 core machines available, we'll need `200_000 requests / 8 core machines = 25K machines`.

## The rule of 72
The rule of 72 estimates how long it'll take for a value to double if it grows some percent rate. The rule is `time = 72/rate` or `t = 72/r`. For example, if requests are increasing at 10% per day, then `t = 72/r = 72/10 = 7.2 days`.

## Sorting
We need to memorize a basic sorting algorithm, like insertion sort:
```go
func InsertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		// This is the magic. j starts at the ith element and walks smaller elements to the front of the array.
		// So, while j-1 > j, swap it so that j-1 is < j and decrement j. Note j > 0 cause we're doing j-- stuff.
		for j := i; j > 0 && arr[j-1] > arr[j]; j-- {
      arr[j-1], arr[j] = arr[j], arr[j-1] // Swap
		}
	}
}
```

## Stacks and queues
If you need to implement your own go stack or queue, you don't need to go crazy with locking, generics, or efficiency. Here's how to code rudimentary stacks and queues.

```go
// A straightforward stack without generics, locking, etc.
type Stack []int

func (s *Stack) Push(x int) {
	*s = append(*s, x)
}

func (s *Stack) Pop() int {
	result := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return result
}

func (s *Stack) Peek() int {
	return (*s)[len(*s)-1]
}
```

```go
// A straightforward queue without generics, locking, etc.
type Queue []int

func (s *Queue) Enq(x int) {
	*s = append(*s, x)
}

func (s *Queue) Deq() int {
	result := (*s)[0]
	*s = (*s)[1:len(*s)]
	return result
}

func (s *Queue) Peek() int {
	return (*s)[0]
}
```

## Data storage

## Networking

## Computations

## DBs with SQL

## DBs with NoSQL

## RAM disk providers
Like Redis or Memcached.

## Storage devices

## Serialization

## Hashing

## Common HTTP status codes
Code | Name | Description
--- | --- | ---
200 | OK | Indicates the request succeeded.
201 | Created | XXXXXXX
202 | Accepted | XXXXXX
204 | No content | XXXXXXX
307 | Temporary redirect | XXXXXX
308 | Permanent redirect | XXXXXXX
400 | Bad request | XXXXXX
401 | Unauthorized | XXXXXX
403 | Forbidden | XXXXXXX
404 | Not found | XXXXXXX
408 | Timeout | XXXXXXX
409 | Conflict | XXXXXX
429 | Too many requests | XXXXXXX
500 | Internal server error | XXXXXX
501 | Not implemented | XXXXXXX
502 | Bad gateway | XXXXXXXX
503 | Service unavailable | XXXXXXX
504 | Gateway timeout | XXXXXXX

# General references
- [Understanding Distributed Systems](https://understandingdistributed.systems/) is a really great book on the topic and worth a deep dive. The references are all free or easy to access and worth deep-diving on too.
- [System design primer on GitHub](https://github.com/donnemartin/system-design-primer) is a solid guide for systems interviews.
- [Site Reliability Engineering](https://sre.google/sre-book/table-of-contents/) from Google is a great resource on reliability, on-call, etc.