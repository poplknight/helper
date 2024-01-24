package stream

type unique[T any, K comparable] struct {
	fc func(T) K
	m  map[K]struct{}
}

func NewHash[T any, K comparable](uniqueFunc func(T) K) IUnique[T] {
	return &unique[T, K]{fc: uniqueFunc, m: make(map[K]struct{})}
}

func (h *unique[T, K]) Exist(item T) bool {
	k := h.fc(item)
	if _, ok := h.m[k]; ok {
		return true
	}
	h.m[k] = struct{}{}
	return false
}
