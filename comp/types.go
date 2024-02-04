package comp

type Types interface {
	int |
		int8 |
		int16 |
		int32 |
		int64 |
		uint |
		uint8 |
		uint16 |
		uint32 |
		uint64 |
		float32 |
		float64 |
		string
}

type Equatable[T any] interface {
	Equals(other T) bool
}

type Lesser[T any] interface {
	Less(other T) bool
}

type Comparable[T any] interface {
	Compare(other T) int
}
