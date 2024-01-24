package stream

import "sort"

type slice[T any] []T

func NewSlice[T any](items ...T) ISliceStream[T] {
	if len(items) == 0 {
		items = make([]T, 0, 8)
	}
	return (*slice[T])(&items)
}

func (s *slice[T]) Filter(f func(T) bool) ISliceStream[T] {
	res := make([]T, 0, len(*s))
	for _, item := range *s {
		if f(item) {
			res = append(res, item)
		}
	}
	return NewSlice[T](res...)
}

func (s *slice[T]) Map(f func(T) T) ISliceStream[T] {
	res := make([]T, 0, len(*s))
	for _, item := range *s {
		res = append(res, f(item))
	}
	return NewSlice[T](res...)
}

func (s *slice[T]) Reduce(f func(T, T) T) ISliceStream[T] {
	if len(*s) == 0 {
		return NewSlice[T]()
	}

	res := (*s)[0]
	for _, item := range (*s)[1:] {
		res = f(res, item)
	}
	return NewSlice[T](res)
}

func (s *slice[T]) Sort(f func(T, T) bool) ISliceStream[T] {
	res := make([]T, len(*s))
	copy(res, *s)
	sort.Slice(res, func(i, j int) bool { return f(res[i], res[j]) })
	return NewSlice[T](res...)
}

func (s *slice[T]) Distinct(m IUnique[T]) ISliceStream[T] {
	res := make([]T, 0, len(*s))
	for _, item := range *s {
		if !m.Exist(item) {
			res = append(res, item)
		}
	}
	return NewSlice[T](res...)
}

func (s *slice[T]) Reverse() ISliceStream[T] {
	res := make([]T, len(*s))
	copy(res, *s)
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-i-1] = res[len(res)-i-1], res[i]
	}
	return NewSlice[T](res...)
}

func (s *slice[T]) ForEach(f func(T)) {
	for _, item := range *s {
		f(item)
	}
}

func (s *slice[T]) ToSlice() []T {
	return *s
}

func (s *slice[T]) MergeStream(streams ...ISliceStream[T]) ISliceStream[T] {
	var items []T
	for _, stream := range streams {
		items = append(items, stream.ToSlice()...)
	}
	return s.Merge(items...)
}

func (s *slice[T]) Merge(items ...T) ISliceStream[T] {
	res := make([]T, len(*s)+len(items))
	copy(res, *s)
	copy(res[len(*s):], items)
	return NewSlice[T](res...)
}

func (s *slice[T]) Head(n int) ISliceStream[T] {
	if len(*s) <= n {
		return s
	}
	return NewSlice[T]((*s)[:n]...)
}

func (s *slice[T]) Tail(n int) ISliceStream[T] {
	if len(*s) <= n {
		return s
	}
	return NewSlice[T]((*s)[len(*s)-n:]...)
}

func (s *slice[T]) Len() int {
	return len(*s)
}

func (s *slice[T]) Peek(f func(T)) ISliceStream[T] {
	for _, item := range *s {
		f(item)
	}
	return s
}
