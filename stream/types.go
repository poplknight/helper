package stream

type ISliceStream[T any] interface {
	// Filter The function provides the function of filtering items
	Filter(func(T) bool) ISliceStream[T]
	// Map The function provides the function of converting items
	Map(func(T) T) ISliceStream[T]
	// Reduce The function provides the function of reducing items
	Reduce(func(T, T) T) ISliceStream[T]
	// Sort The function provides the function of sorting items
	Sort(func(T, T) bool) ISliceStream[T]
	// Distinct The function provides the function of deduplication items
	Distinct(IUnique[T]) ISliceStream[T]
	// Reverse The function provides the function of reversing items
	Reverse() ISliceStream[T]
	// Peek The function provides the function of peeking items
	Peek(func(T)) ISliceStream[T]
	// MergeStream The function provides the function of merging ISliceStream items
	MergeStream(...ISliceStream[T]) ISliceStream[T]
	// Merge The function provides the function of merging slice
	Merge(...T) ISliceStream[T]
	// Head The function provides the function of getting the first item
	Head(int) ISliceStream[T]
	// Tail The function provides the function of getting the last item
	Tail(int) ISliceStream[T]
	// Len The function provides the function of getting the number of items
	Len() int
	// ToSlice The function provides the function of converting to slice
	ToSlice() []T
	// ForEach The function provides the function of traversing items
	ForEach(func(T))
}

type IUnique[T any] interface {
	// Exist used to check whether the key exists
	Exist(T) bool
}
